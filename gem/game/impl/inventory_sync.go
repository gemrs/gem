package impl

import (
	"github.com/gemrs/gem/gem/protocol"
)

func (p *Player) SyncInventories() {
	inventory := p.Profile().Inventory()
	updatedSlots := inventory.GetUpdatedSlots()
	// FIXME needs optimizing: there are batch inventory update packets
	for _, slot := range updatedSlots {
		updatePacket := protocol.OutboundUpdateInventoryItem{
			Container: inventory,
			Slot:      slot,
		}

		p.Send(updatePacket)
	}
	inventory.ClearUpdatedSlots()
}
