package packet

import (
	"fmt"

	"github.com/gemrs/gem/gem/game/data"
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol"
)

func init() {
	registerHandler((*protocol.InboundInventorySwapItem)(nil), player_inv_swap)
	registerHandler((*protocol.InboundInventoryAction)(nil), player_inv_action)
	registerHandler((*protocol.InboundTakeGroundItem)(nil), player_take_ground_item)
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

type TakeGroundItemInteraction struct {
	item entity.GroundItem
}

func (i *TakeGroundItemInteraction) Tick(e entity.Entity) bool {
	p := e.(protocol.Player)
	groundItem := i.item

	select {
	case <-groundItem.Expired():
		return true
	default:
	}

	if p.Profile().Inventory().Add(groundItem.Item()) != nil {
		p.SendMessage("You don't have enough inventory space to hold that item")
		return true
	}

	groundItem.Expire()
	return true
}

func (i *TakeGroundItemInteraction) Interruptible() bool {
	return true
}

func player_take_ground_item(p protocol.Player, message server.Message) {
	takeItemPacket := message.(*protocol.InboundTakeGroundItem)
	itemPos := position.NewAbsolute(takeItemPacket.X, takeItemPacket.Y, p.Position().Z())
	entities := p.WorldInstance().EntitiesOnTile(itemPos)

	var groundItem entity.GroundItem
	for _, e := range entities {
		if item, ok := e.(entity.GroundItem); ok {
			if item.Definition().Id() == takeItemPacket.ItemID {
				groundItem = item
				break
			}
		}
	}

	if groundItem == nil {
		return
	}

	p.SetWalkDestination(itemPos)
	p.InteractionQueue().Append(&TakeGroundItemInteraction{
		item: groundItem,
	})
}
