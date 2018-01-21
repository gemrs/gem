package impl

import (
	"github.com/gemrs/gem/gem/game/data"
	"github.com/gemrs/gem/gem/game/item"
	"github.com/gemrs/gem/gem/protocol"
)

type Equipment struct {
	container *item.Container
	player    protocol.Player
}

func NewEquipment() *Equipment {
	equipment := &Equipment{
		container: item.NewContainer(data.Int("inventory.equipment_size")),
	}

	equipment.container.SetInterfaceLocation(
		// It looks like we actually want the widget id to be anything < 0,
		// the official server sends this seemingly arbitrary number, so we will too.
		0xFFFF, 0xFAD0,
		data.Int("inventory.equipment"))

	return equipment
}

func (e *Equipment) Equip(slot int, stack *item.Stack) (oldEquipment *item.Stack) {
	if e.container.SlotPopulated(slot) {
		oldEquipment = e.container.RemoveAllFromSlot(slot)
		e.container.Add(oldEquipment)
	}

	e.container.SetSlot(slot, stack)
	e.signalUpdate()
	return
}

func (e *Equipment) Slot(i int) *item.Stack {
	return e.Container().Slot(i)
}

func (e *Equipment) Container() *item.Container {
	return e.container
}

func (e *Equipment) setPlayer(p protocol.Player) {
	e.player = p
	e.signalUpdate()
}

func (e *Equipment) signalUpdate() {
	if e.player != nil {
		e.player.SetAppearanceChanged()
	}
}
