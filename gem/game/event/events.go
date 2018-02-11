//glua:bind module gem.game.event
package game_event

import (
	"github.com/gemrs/gem/gem/core/event"
)

//go:generate glua .

//glua:bind
var (
	PlayerLoadProfile      = event.NewEvent("PlayerLoadProfile")
	PlayerLogin            = event.NewEvent("PlayerLogin")
	PlayerLogout           = event.NewEvent("PlayerLogout")
	PlayerCommand          = event.NewEvent("PlayerCommand")
	PlayerInventoryAction  = event.NewEvent("PlayerInventoryAction")
	PlayerGroundItemAction = event.NewEvent("PlayerGroundItemAction")
	PlayerWidgetAction     = event.NewEvent("PlayerWigetAction")
)
