package event

import (
	"github.com/gemrs/gem/gem/event"
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
	event := event.NewEvent(key)
	events = append(events, event)

	return event
}
