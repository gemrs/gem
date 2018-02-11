package impl

import (
	"github.com/gemrs/gem/gem/core/event"
	"github.com/gemrs/gem/gem/game/data"
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/item"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/protocol"
	"github.com/gemrs/gem/gem/util/expire"
)

func init() {
	game_event.PlayerInventoryAction.Register(event.NewObserver(expire.NewNonExpirable(), playerInventoryAction))
	game_event.PlayerGroundItemAction.Register(event.NewObserver(expire.NewNonExpirable(), playerGroundItemAction))
	game_event.PlayerWidgetAction.Register(event.NewObserver(expire.NewNonExpirable(), playerUnequipItem))
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

func playerGroundItemAction(event *event.Event, args ...interface{}) {
	p := args[0].(*Player)
	action := args[1].(string)
	itemPos := args[2].(*position.Absolute)
	groundItem := args[3].(entity.GroundItem)

	switch action {
	case "Take":
		p.SetWalkDestination(itemPos)
		p.InteractionQueue().Append(&TakeGroundItemInteraction{
			item: groundItem,
		})

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

func playerUnequipItem(event *event.Event, args ...interface{}) {
	p := args[0].(*Player)
	action := args[1].(int)
	interfaceId := args[2].(int)
	widgetId := args[3].(int)

	if action != 0 || interfaceId != data.Int("widget.equipment_group_id") {
		return
	}

	slot := widgetId - 6

	equipment := p.Profile().Equipment()
	inventory := p.Profile().Inventory()

	stack := equipment.Unequip(slot)
	inventory.Add(stack)
}
