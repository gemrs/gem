package event

import (
	"github.com/sinusoids/gem/gem/log"
	"github.com/sinusoids/gem/gem/util/expire"
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
	owner  expire.Expirable
}

func NewListener(owner expire.Expirable, fn Callback) *Listener {
	return &Listener{
		id:    <-nextId,
		fn:    fn,
		owner: owner,
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

	select {
	case <-l.owner.Expired():
		// If the owner expired, unregister this listener
		e.Unregister(l)
	default:
		l.fn(e, args...)
	}
}
