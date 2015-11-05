package log

import (
	"io"
	"text/template"
)

var standardFormat = "[{{.Level}}] {{.Tag}}: {{.Message}}"

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

func (t *TextTarget) SetFormat(fmt string) {
	t.fmt = template.Must(template.New("record").Parse(fmt))
}

func (t *TextTarget) Handle(r Record) {
	err := t.fmt.ExecuteTemplate(t.w, "record", r)
	if err != nil {
		panic(err)
	}
}

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

func (b *BufferingTarget) Buffered() []Record {
	return b.buffer
}

func (b *BufferingTarget) Redirect() {
	b.buffer = make([]Record, 0)
	b.redirect = true
}

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
