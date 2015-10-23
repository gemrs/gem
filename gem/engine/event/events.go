package event

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/event"
	"github.com/sinusoids/gem/gem/python"
)

var (
	Startup  = createEvent("Startup")
	Shutdown = createEvent("Shutdown")
	PreTick  = createEvent("PreTick")
	Tick     = createEvent("Tick")
	PostTick = createEvent("PostTick")
)

var events = []*event.Event{}

func createEvent(key string) *event.Event {
	event, err := event.NewEvent(key)
	if err != nil {
		panic(err)
	}

	events = append(events, event)

	return event
}

func init() {
	/* Create package */
	var err error
	var module *py.Module
	if module, err = python.InitModule("gem.engine.event", []py.Method{}); err != nil {
		panic(err)
	}

	createEventObjects(module)
}

func createEventObjects(module *py.Module) {
	for _, event := range events {
		if err := module.AddObject(event.Key(), event); err != nil {
			panic(err)
		}
	}
}
