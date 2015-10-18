package gem

import (
	"github.com/qur/gopy/lib"

	"gem/event"
)

var (
	StartupEvent  = createEvent("Startup")
	ShutdownEvent = createEvent("Shutdown")
	PreTickEvent  = createEvent("PreTick")
	TickEvent     = createEvent("Tick")
	PostTickEvent = createEvent("PostTick")
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

func createEventObjects(module *py.Module) {
	for _, event := range events {
		if err := module.AddObject(event.Key(), event); err != nil {
			panic(err)
		}
	}
}
