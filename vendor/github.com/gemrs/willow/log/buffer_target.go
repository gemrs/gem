package log

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
