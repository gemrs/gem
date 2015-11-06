package log

import (
	"io"
	"text/template"
)

var standardFormat = "[{{.Level}}] {{.Tag}}: {{.Message}}"

// TextTarget formats Records using a template, and writes to an io.Writer
type TextTarget struct {
	w   io.Writer
	fmt *template.Template
}

func NewTextTarget(w io.Writer) *TextTarget {
	target := &TextTarget{
		w: w,
	}
	target.SetFormat(standardFormat)
	return target
}

// SetFormat customizes the log template. The Record is passed as a parameter to the template.
func (t *TextTarget) SetFormat(fmt string) {
	t.fmt = template.Must(template.New("record").Parse(fmt))
}

// Handle formats a record and writes to the io.Writer
func (t *TextTarget) Handle(r Record) {
	err := t.fmt.ExecuteTemplate(t.w, "record", r)
	if err != nil {
		panic(err)
	}
}

// BufferingTarget sits between the dispatcher and another Handler, and can temporarily
// buffer all records to memory. When flushed, a buffering target forwards all buffered
// records to the target.
type BufferingTarget struct {
	buffer   []Record
	redirect bool
	target   Handler
}

func NewBufferingTarget(target Handler) *BufferingTarget {
	return &BufferingTarget{
		target:   target,
		redirect: false,
	}
}

// Buffered returns the slice of buffered records
func (b *BufferingTarget) Buffered() []Record {
	return b.buffer
}

// Redirect turns on buffering and stops forwarding Records to the handler.
func (b *BufferingTarget) Redirect() {
	b.buffer = make([]Record, 0)
	b.redirect = true
}

// Flush flushes all buffered records, clears the buffer, and turns off buffering
func (b *BufferingTarget) Flush() {
	for _, r := range b.buffer {
		b.target.Handle(r)
	}
	b.redirect = false
	b.buffer = nil
}

func (b *BufferingTarget) Handle(r Record) {
	if b.redirect {
		b.buffer = append(b.buffer, r)
	} else {
		b.target.Handle(r)
	}
}
