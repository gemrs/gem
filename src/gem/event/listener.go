package event

import (
	"gem/util/id"
)

var nextId <-chan int

func init() {
	nextId = id.Generator()
}

type Callback func(*Event, ...interface{})

type Listener struct {
	id int
	fn Callback
}

func (l *Listener) Init(fn Callback) error {
	l.id = <-nextId
	l.fn = fn
	return nil
}

func (l *Listener) Id() int {
	return l.id
}

func (l *Listener) Notify(e *Event, args ...interface{}) {
	l.fn(e, args...)
}
