package packet

import (
	"fmt"

	"github.com/gemrs/gem/gem/encoding"
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
		actionIndex = 1

	case *game_protocol.InboundInventoryAction2:
		slot = int(packet.Slot)
		interfaceId = int(packet.InterfaceID)
		itemId = int(packet.ItemID)
		actionIndex = 2

	case *game_protocol.InboundInventoryAction3:
		slot = int(packet.Slot)
		interfaceId = int(packet.InterfaceID)
		itemId = int(packet.ItemID)
		actionIndex = 3

	case *game_protocol.InboundInventoryAction4:
		slot = int(packet.Slot)
		interfaceId = int(packet.InterfaceID)
		itemId = int(packet.ItemID)
		actionIndex = 4

	case *game_protocol.InboundInventoryAction5:
		slot = int(packet.Slot)
		interfaceId = int(packet.InterfaceID)
		itemId = int(packet.ItemID)
		actionIndex = 5

	default:
		panic("Unknown inventory action packet")
	}

	// FIXME validate
	// FIXME work out actual action by loading from obj.dat
	actionString := ""
	if actionIndex == 5 {
		actionString = "Drop"
	}

	_ = itemId

	switch interfaceId {
	case player.RevisionConstants.InventoryInterfaceId:
		stack := p.Profile().Inventory().Slot(slot)
		game_event.PlayerInventoryAction.NotifyObservers(p, stack, slot, actionString)

	default:
		panic(fmt.Sprintf("unknown inventory interface id: %v", interfaceId))
	}
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
		panic("Player asked to pick up item that doesn't exist")
	}

	fmt.Printf("picking up %v\n", groundItem)

	if p.Profile().Inventory().Add(groundItem.Item()) != nil {
		p.SendMessage("You don't have enough inventory space to hold that item")
		return
	}

	groundItem.Expire()
}
