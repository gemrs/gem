package event

import (
	"github.com/gemrs/willow/log"
)

type Event struct {
	key       string
	observers map[int]ObserverIface
	log       log.Log
}

func NewEvent(key string) *Event {
	return &Event{
		key:       key,
		observers: make(map[int]ObserverIface),
		log:       log.New("event", log.MapContext{"event": key}),
	}
}

func (e *Event) Key() string {
	return e.key
}

func (e *Event) Register(o ObserverIface) {
	e.observers[o.Id()] = o
}

func (e *Event) Unregister(o ObserverIface) {
	delete(e.observers, o.Id())
}

func (e *Event) NotifyObservers(args ...interface{}) {
	for _, observer := range e.observers {
		observer.Notify(e, args...)
	}
}
