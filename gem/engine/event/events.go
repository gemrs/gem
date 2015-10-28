package event

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/event"
	"github.com/sinusoids/gem/gem/python/modules"
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
	event := event.NewEvent(key)
	events = append(events, event)

	return event
}

func init() {
	lock := py.NewLock()
	defer lock.Unlock()

	/* Create package */
	var err error
	var module *py.Module
	if module, err = modules.Init("gem.engine.event", []py.Method{}); err != nil {
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
