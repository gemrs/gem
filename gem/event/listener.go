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
	id    int
	fn    Callback
	owner expire.Expirable
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

func (l *Listener) Notify(e *Event, args ...interface{}) {
	log := e.log.Child("listener", log.MapContext{"id": l.id})
	defer safe.Recover(log)

	select {
	case <-l.owner.Expired():
		// If the owner expired, unregister this listener
		e.Unregister(l)
	default:
		l.fn(e, args...)
	}
}
