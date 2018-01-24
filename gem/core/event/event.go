//glua:bind module gem.event
package event

import (
	"github.com/gemrs/willow/log"
)

//go:generate glua .

//glua:bind
type Event struct {
	key       string
	observers map[int]ObserverIface
	log       log.Log
}

//glua:bind constructor Event
func NewEvent(key string) *Event {
	return &Event{
		key:       key,
		observers: make(map[int]ObserverIface),
		log:       log.New("event", log.MapContext{"event": key}),
	}
}

//glua:bind
func (e *Event) Key() string {
	return e.key
}

//glua:bind
func (e *Event) Register(o ObserverIface) {
	e.observers[o.Id()] = o
}

//glua:bind
func (e *Event) Unregister(o ObserverIface) {
	delete(e.observers, o.Id())
}

func (e *Event) NotifyObservers(args ...interface{}) {
	for _, observer := range e.observers {
		observer.Notify(e, args...)
	}
}
