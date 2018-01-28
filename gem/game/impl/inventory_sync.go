package impl

import (
	"github.com/gemrs/gem/gem/protocol"
)

func (p *Player) SyncInventories() {
	inventory := p.Profile().Inventory()
	updatedSlots := inventory.GetUpdatedSlots()
	// FIXME needs optimizing: there are batch inventory update packets
	for _, slot := range updatedSlots {
		item := inventory.Slot(slot)
		updatePacket := protocol.OutboundUpdateInventoryItem{
			InventoryID: p.CurrentFrame().InventoryInterface(),
			Slot:        slot,
		}
		if item != nil {
			updatePacket.ItemID = item.Definition().Id() + 1
			updatePacket.Count = item.Count()
		}

		p.Send(updatePacket)
	}
	inventory.ClearUpdatedSlots()
}
