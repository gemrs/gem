package event

import (
	"github.com/qur/gopy/lib"

	"gem/event"
	"gem/python"
)

var (
	PlayerLoadProfile      = createEvent("PlayerLoadProfile")
	PlayerLogin            = createEvent("PlayerLogin")
	PlayerLogout           = createEvent("PlayerLogout")
	PlayerFinishLogin      = createEvent("PlayerFinishLogin")
	PlayerSectorChange     = createEvent("PlayerSectorChange")
	PlayerRegionChange     = createEvent("PlayerRegionChange")
	PlayerAppearanceUpdate = createEvent("PlayerAppearanceUpdate")
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
	if module, err = python.InitModule("gem.game.event", []py.Method{}); err != nil {
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
