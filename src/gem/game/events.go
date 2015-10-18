package game

import (
	"github.com/qur/gopy/lib"

	"gem/event"
)

var (
	PlayerLoadProfileEvent  = createEvent("PlayerLoadProfile")
	PlayerLoginEvent        = createEvent("PlayerLogin")
	PlayerLogoutEvent       = createEvent("PlayerLogout")
	PlayerSectorChangeEvent = createEvent("PlayerSectorChange")
	PlayerRegionChangeEvent = createEvent("PlayerRegionChange")
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
