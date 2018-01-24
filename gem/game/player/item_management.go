package player

import (
	"github.com/gemrs/gem/gem/core/event"
	"github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/item"
	"github.com/gemrs/gem/gem/util/expire"
)

func init() {
	game_event.PlayerInventoryAction.Register(event.NewObserver(expire.NewNonExpirable(), PlayerDropItem))
}

func PlayerDropItem(event *event.Event, args ...interface{}) {
	p := args[0].(*Player)
	//stack := args[1].(*item.Stack)
	slot := args[2].(int)
	action := args[3].(string)
	if action != "Drop" {
		return
	}

	stack := p.Profile().Inventory().RemoveAllFromSlot(slot)
	item.NewGroundItem(stack, p.Position(), p.world)
}
