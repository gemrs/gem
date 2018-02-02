package item

import (
	"errors"
	"fmt"
)

var ErrNoFreeSlots = errors.New("No free slots in inventory")

//glua:bind
type Container struct {
	capacity        int
	slots           []*Stack
	slotsUpdated    map[int]interface{}
	rootId, childId int
	interfaceId     int
}

//glua:bind constructor Container
func NewContainer(capacity int) *Container {
	container := &Container{
		capacity:     capacity,
		slots:        make([]*Stack, capacity),
		slotsUpdated: make(map[int]interface{}),
	}

	// All containers start with all slots updated to force a clear
	for i := 0; i < capacity; i++ {
		container.slotsUpdated[i] = 0
	}

	return container
}

func (c *Container) SetInterfaceLocation(root, child, iface int) {
	c.rootId = root
	c.childId = child
	c.interfaceId = iface
}

func (c *Container) InterfaceLocation() (root, child, iface int) {
	return c.rootId, c.childId, c.interfaceId
}

func (c *Container) setSlotUpdated(slot int) {
	c.slotsUpdated[slot] = 0
}

func (c *Container) GetUpdatedSlots() []int {
	slots := make([]int, 0)
	for s, _ := range c.slotsUpdated {
		slots = append(slots, s)
	}
	return slots
}

func (c *Container) ClearUpdatedSlots() {
	if len(c.slotsUpdated) > 0 {
		c.slotsUpdated = make(map[int]interface{})
	}
}

//glua:bind
func (c *Container) Capacity() int {
	return c.capacity
}

func (c *Container) assertSize(i int) {
	if len(c.slots) < i {
		panic(fmt.Sprintf("no slot at %v in container of size %v", i, c.capacity))
	}
}

func (c *Container) updateStackSize(slot, count int) {
	c.assertSlotPopulated(true, slot)
	stack := c.Slot(slot)
	newCount := stack.Count() + count
	if newCount < 0 {
		panic(fmt.Sprintf("tried to remove %v from item stack of size %v", count, stack.Count()))
	}

	if newCount == 0 {
		c.slots[slot] = nil
	} else {
		c.slots[slot].count = newCount
	}

	c.setSlotUpdated(slot)
}

//glua:bind
func (c *Container) Add(item *Stack) error {
	slot := c.FindStackOf(item.Definition())
	if item.Definition().Stackable() && slot != -1 {
		c.updateStackSize(slot, item.Count())
		return nil
	}

	slot = c.findEmptySlot()
	if slot != -1 {
		c.SetSlot(slot, item)
		return nil
	}

	return ErrNoFreeSlots
}

func (c *Container) findEmptySlot() int {
	for i, stack := range c.slots {
		if stack == nil {
			return i
		}
	}
	return -1
}

//glua:bind
func (c *Container) Slot(i int) *Stack {
	c.assertSize(i)
	return c.slots[i]
}

//glua:bind
func (c *Container) SetSlot(i int, item *Stack) {
	c.assertSize(i)
	c.assertSlotPopulated(false, i)
	c.slots[i] = item
	c.setSlotUpdated(i)
}

func (c *Container) SlotPopulated(i int) bool {
	c.assertSize(i)
	return c.slots[i] != nil
}

func (c *Container) assertSlotPopulated(populated bool, i int) {
	stateStr := "populated"
	if !populated {
		stateStr = "unpopulated"
	}

	if c.SlotPopulated(i) != populated {
		panic(fmt.Sprintf("expected slot %v to be %v", i, stateStr))
	}
}

//glua:bind
func (c *Container) SwapSlots(a, b int) {
	// Don't assert slots are populated, because we could be moving to an empty slot
	c.assertSize(a)
	c.assertSize(b)
	c.slots[a], c.slots[b] = c.slots[b], c.slots[a]
	c.setSlotUpdated(a)
	c.setSlotUpdated(b)
}

//glua:bind
func (c *Container) FindStackOf(item *Definition) int {
	for i, stack := range c.slots {
		if stack != nil && stack.Definition() == item {
			return i
		}
	}
	return -1
}

//glua:bind
func (c *Container) RemoveFromSlot(slot, count int) *Stack {
	c.assertSlotPopulated(true, slot)
	stack := c.Slot(slot)
	c.updateStackSize(slot, -count)

	newStack := NewStack(stack.Definition(), count)
	return newStack
}

//glua:bind
func (c *Container) RemoveAllFromSlot(slot int) *Stack {
	c.assertSlotPopulated(true, slot)
	stack := c.Slot(slot)
	return c.RemoveFromSlot(slot, stack.Count())
}
