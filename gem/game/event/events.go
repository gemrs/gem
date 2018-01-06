//glua:bind module gem.game.event
package game_events

import (
	"github.com/gemrs/gem/gem/event"
)

//go:generate glua .

//glua:bind
var (
	PlayerLoadProfile      = event.NewEvent("PlayerLoadProfile")
	PlayerLogin            = event.NewEvent("PlayerLogin")
	PlayerLogout           = event.NewEvent("PlayerLogout")
	PlayerFinishLogin      = event.NewEvent("PlayerFinishLogin")
	EntitySectorChange     = event.NewEvent("EntitySectorChange")
	EntityRegionChange     = event.NewEvent("EntityRegionChange")
	PlayerAppearanceUpdate = event.NewEvent("PlayerAppearanceUpdate")
)
