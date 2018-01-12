package packet

import (
	"github.com/gemrs/gem/gem/encoding"
	"github.com/gemrs/gem/gem/game/player"
	"github.com/gemrs/gem/gem/protocol/game_protocol"
)

func init() {
	registerHandler((*game_protocol.InboundInventorySwapItem)(nil), player_inv_swap)
}

func player_inv_swap(p *player.Player, packet encoding.Decodable) {
	swapItemPacket := packet.(*game_protocol.InboundInventorySwapItem)
	// FIXME do something with InterfaceID
	p.Profile().Inventory().SwapSlots(int(swapItemPacket.FromSlot), int(swapItemPacket.ToSlot))
}
