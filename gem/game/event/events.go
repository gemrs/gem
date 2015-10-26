package event

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/event"
	"github.com/sinusoids/gem/gem/python/modules"
)

var (
	PlayerLoadProfile      = createEvent("PlayerLoadProfile")
	PlayerLogin            = createEvent("PlayerLogin")
	PlayerLogout           = createEvent("PlayerLogout")
	PlayerFinishLogin      = createEvent("PlayerFinishLogin")
	EntitySectorChange     = createEvent("EntitySectorChange")
	EntityRegionChange     = createEvent("EntityRegionChange")
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
	lock := py.NewLock()
	defer lock.Unlock()

	/* Create package */
	var err error
	var module *py.Module
	if module, err = modules.Init("gem.game.event", []py.Method{}); err != nil {
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
