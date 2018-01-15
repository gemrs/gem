package packet

import (
	"fmt"

	"github.com/gemrs/gem/gem/encoding"
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/item"
	"github.com/gemrs/gem/gem/game/player"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/protocol/game_protocol"
)

func init() {
	registerHandler((*game_protocol.InboundInventorySwapItem)(nil), player_inv_swap)
	registerHandler((*game_protocol.InboundInventoryAction1)(nil), player_inv_action)
	registerHandler((*game_protocol.InboundInventoryAction2)(nil), player_inv_action)
	registerHandler((*game_protocol.InboundInventoryAction3)(nil), player_inv_action)
	registerHandler((*game_protocol.InboundInventoryAction4)(nil), player_inv_action)
	registerHandler((*game_protocol.InboundInventoryAction5)(nil), player_inv_action)
	registerHandler((*game_protocol.InboundTakeGroundItem)(nil), player_take_ground_item)
}

func player_inv_swap(p *player.Player, packet encoding.Decodable) {
	swapItemPacket := packet.(*game_protocol.InboundInventorySwapItem)
	switch int(swapItemPacket.InterfaceID) {
	case player.RevisionConstants.InventoryInterfaceId:
		p.Profile().Inventory().SwapSlots(int(swapItemPacket.FromSlot), int(swapItemPacket.ToSlot))

	default:
		panic(fmt.Sprintf("unknown inventory interface id: %v", swapItemPacket.InterfaceID))
	}

}

func player_inv_action(p *player.Player, packet encoding.Decodable) {
	var slot, interfaceId, itemId, actionIndex int
	switch packet := packet.(type) {
	case *game_protocol.InboundInventoryAction1:
		slot = int(packet.Slot)
		interfaceId = int(packet.InterfaceID)
		itemId = int(packet.ItemID)
		actionIndex = 0

	case *game_protocol.InboundInventoryAction2:
		slot = int(packet.Slot)
		interfaceId = int(packet.InterfaceID)
		itemId = int(packet.ItemID)
		actionIndex = 1

	case *game_protocol.InboundInventoryAction3:
		slot = int(packet.Slot)
		interfaceId = int(packet.InterfaceID)
		itemId = int(packet.ItemID)
		actionIndex = 2

	case *game_protocol.InboundInventoryAction4:
		slot = int(packet.Slot)
		interfaceId = int(packet.InterfaceID)
		itemId = int(packet.ItemID)
		actionIndex = 3

	case *game_protocol.InboundInventoryAction5:
		slot = int(packet.Slot)
		interfaceId = int(packet.InterfaceID)
		itemId = int(packet.ItemID)
		actionIndex = 4

	default:
		panic("Unknown inventory action packet")
	}

	switch interfaceId {
	case player.RevisionConstants.InventoryInterfaceId:
		stack := p.Profile().Inventory().Slot(slot)

		if stack.Definition().Id() != itemId {
			panic("inventory action validation failed")
		}

		actions := stack.Definition().Actions()
		actionString := actions[actionIndex]
		if actionIndex == 4 {
			actionString = "Drop"
		}

		game_event.PlayerInventoryAction.NotifyObservers(p, stack, slot, actionString)

	default:
		panic(fmt.Sprintf("unknown inventory interface id: %v", interfaceId))
	}
}

type TakeGroundItemInteraction struct {
	item *item.GroundItem
}

func (i *TakeGroundItemInteraction) Tick(e entity.Entity) bool {
	p := e.(*player.Player)
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

func player_take_ground_item(p *player.Player, packet encoding.Decodable) {
	takeItemPacket := packet.(*game_protocol.InboundTakeGroundItem)
	itemPos := position.NewAbsolute(int(takeItemPacket.X), int(takeItemPacket.Y), p.Position().Z())
	entities := p.WorldInstance().EntitiesOnTile(itemPos)

	var groundItem *item.GroundItem
	for _, entity := range entities {
		if item, ok := entity.(*item.GroundItem); ok {
			if item.Definition().Id() == int(takeItemPacket.ItemID) {
				groundItem = item
				break
			}
		}
	}

	if groundItem == nil {
		return
	}

	p.InteractionQueue().Append(&TakeGroundItemInteraction{
		item: groundItem,
	})
}
