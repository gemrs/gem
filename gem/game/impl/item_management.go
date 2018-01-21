package impl

import (
	"github.com/gemrs/gem/gem/core/event"
	"github.com/gemrs/gem/gem/game/data"
	"github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/item"
	"github.com/gemrs/gem/gem/protocol"
	"github.com/gemrs/gem/gem/util/expire"
)

func init() {
	game_event.PlayerInventoryAction.Register(event.NewObserver(expire.NewNonExpirable(), playerInventoryAction))
}

func (p *Player) SyncInventories() {
	p.syncContainer(p.Profile().Inventory())
	p.syncContainer(p.Profile().Equipment().Container())
}

func (p *Player) syncContainer(container *item.Container) {
	updatedSlots := container.GetUpdatedSlots()
	// FIXME needs optimizing: there are batch container update packets
	for _, slot := range updatedSlots {
		updatePacket := &protocol.OutboundUpdateInventoryItem{
			Container: container,
			Slot:      slot,
		}

		p.Send(updatePacket)
	}
	container.ClearUpdatedSlots()
}

func playerInventoryAction(event *event.Event, args ...interface{}) {
	p := args[0].(*Player)
	stack := args[1].(*item.Stack)
	slot := args[2].(int)
	action := args[3].(string)

	switch action {
	case "Drop":
		stack := p.Profile().Inventory().RemoveAllFromSlot(slot)
		NewGroundItem(stack, p.Position(), p.world)

	case "Wield":
		fallthrough
	case "Wear":
		equipment := p.Profile().Equipment()
		inventory := p.Profile().Inventory()

		equipmentDef := data.Equipment[stack.Definition().Id()]
		// FIXME probably need some kind of transaction system for stuff like this
		oldEquipment := equipment.Equip(equipmentDef.Slot(), stack)
		inventory.RemoveAllFromSlot(slot)
		if oldEquipment != nil {
			inventory.Add(oldEquipment)
		}
	}

}
