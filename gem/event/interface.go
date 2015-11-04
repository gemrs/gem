package event

import (
	"github.com/sinusoids/gem/gem/log"
)

type Observable interface {
	Key() string
	Register(Observer)
	Unregister(Observer)
	NotifyObservers(...interface{})
}

type Observer interface {
	Id() int
	Notify(*Event, ...interface{})
	setLogger(log.Logger)
}
