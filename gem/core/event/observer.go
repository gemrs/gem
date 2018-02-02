package event

import (
	"github.com/gemrs/gem/gem/util/expire"
	"github.com/gemrs/gem/gem/util/id"
	"github.com/gemrs/gem/gem/util/safe"
	"github.com/gemrs/willow/log"
)

var nextId <-chan int

func init() {
	nextId = id.Generator(0)
}

type Callback func(*Event, ...interface{})

type Observer struct {
	id    int
	fn    Callback
	owner expire.Expirable
}

func NewObserver(owner expire.Expirable, fn Callback) *Observer {
	return &Observer{
		id:    <-nextId,
		fn:    fn,
		owner: owner,
	}
}

func (l *Observer) Id() int {
	return l.id
}

func (l *Observer) Notify(e *Event, args ...interface{}) {
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
