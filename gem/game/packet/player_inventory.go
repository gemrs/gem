package packet

import (
	"fmt"

	"github.com/gemrs/gem/gem/encoding"
	"github.com/gemrs/gem/gem/game/player"
	"github.com/gemrs/gem/gem/protocol/game_protocol"
)

func init() {
	registerHandler((*game_protocol.InboundInventorySwapItem)(nil), player_inv_swap)
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
