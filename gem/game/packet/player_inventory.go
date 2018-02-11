package packet

import (
	"fmt"

	"github.com/gemrs/gem/gem/game/data"
	"github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol"
)

func init() {
	registerHandler((*protocol.InboundInventorySwapItem)(nil), player_inv_swap)
	registerHandler((*protocol.InboundInventoryAction)(nil), player_inv_action)
}

func player_inv_swap(p protocol.Player, message server.Message) {
	swapItemPacket := message.(*protocol.InboundInventorySwapItem)
	switch swapItemPacket.InterfaceID {
	case data.Int("widget.inventory_group_id"):
		p.Profile().Inventory().SwapSlots(swapItemPacket.From, swapItemPacket.To)

	default:
		panic(fmt.Sprintf("unknown inventory interface id: %v", swapItemPacket.InterfaceID))
	}

}

func player_inv_action(p protocol.Player, message server.Message) {
	action := message.(*protocol.InboundInventoryAction)

	switch action.InterfaceID {
	case data.Int("widget.inventory_group_id"):
		stack := p.Profile().Inventory().Slot(action.Slot)

		if stack.Definition().Id() != action.ItemID {
			panic("inventory action validation failed")
		}

		actions := stack.Definition().Actions()
		actionString := actions[action.Action]

		game_event.PlayerInventoryAction.NotifyObservers(p, stack, action.Slot, actionString)

	default:
		panic(fmt.Sprintf("unknown inventory interface id: %v", action.InterfaceID))
	}
}
