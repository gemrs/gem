package player

import (
	"github.com/gemrs/gem/gem/encoding"
	"github.com/gemrs/gem/gem/protocol/game_protocol"
)

func (p *Player) SyncInventories() {
	inventory := p.Profile().Inventory()
	updatedSlots := inventory.GetUpdatedSlots()
	// FIXME needs optimizing: there are batch inventory update packets
	for _, slot := range updatedSlots {
		item := inventory.Slot(slot)
		updatePacket := &game_protocol.OutboundUpdateInventoryItem{
			InventoryID: encoding.Uint16(RevisionConstants.InventoryInterfaceId),
			Slot:        encoding.Uint8(slot),
		}
		if item != nil {
			updatePacket.ItemID = encoding.Uint16(item.Definition().Id() + 1)
			updatePacket.Count = encoding.Uint8(item.Count())
		}

		p.Conn().Write <- updatePacket
	}
	inventory.ClearUpdatedSlots()
}
