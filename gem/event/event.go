package event

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/log"
)

type Event struct {
	py.BaseObject

	key       string
	observers map[int]Observer
	logger    log.Logger
}

func (e *Event) Init(key string) {
	e.key = key
	e.observers = make(map[int]Observer)
	e.logger = log.New(key)
}

func (e *Event) Key() string {
	return e.key
}

func (e *Event) Register(o Observer) {
	o.setLogger(e.logger)
	e.observers[o.Id()] = o
}

func (e *Event) Unregister(o Observer) {
	delete(e.observers, o.Id())
}

func (e *Event) NotifyObservers(args ...interface{}) {
	for _, observer := range e.observers {
		observer.Notify(e, args...)
	}
}
