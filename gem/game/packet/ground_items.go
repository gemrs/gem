package packet

import (
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol"
)

func init() {
	registerHandler((*protocol.InboundGroundItemAction)(nil), player_ground_item_action)
}

func player_ground_item_action(p protocol.Player, message server.Message) {
	action := message.(*protocol.InboundGroundItemAction)
	itemPos := position.NewAbsolute(action.X, action.Y, p.Position().Z())
	entities := p.WorldInstance().EntitiesOnTile(itemPos)

	var groundItem entity.GroundItem
	for _, e := range entities {
		if item, ok := e.(entity.GroundItem); ok {
			if item.Definition().Id() == action.ItemID {
				groundItem = item
				break
			}
		}
	}

	if groundItem == nil {
		return
	}

	actions := groundItem.Definition().GroundActions()
	actionString := actions[action.Action]

	game_event.PlayerGroundItemAction.NotifyObservers(p, actionString, itemPos, groundItem)
}
