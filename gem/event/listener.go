package event

import (
	"github.com/sinusoids/gem/gem/log"
	"github.com/sinusoids/gem/gem/util/id"
	"github.com/sinusoids/gem/gem/util/safe"
)

var nextId <-chan int

func init() {
	nextId = id.Generator()
}

type Callback func(*Event, ...interface{})

type Listener struct {
	id     int
	fn     Callback
	logger log.Logger
}

func NewListener(fn Callback) *Listener {
	return &Listener{
		id: <-nextId,
		fn: fn,
	}
}

func (l *Listener) Id() int {
	return l.id
}

func (l *Listener) setLogger(logger log.Logger) {
	l.logger = logger
}

func (l *Listener) Notify(e *Event, args ...interface{}) {
	defer safe.Recover(l.logger)

	l.fn(e, args...)
}
