package event

import (
	"github.com/gemrs/gem/gem/event"
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
