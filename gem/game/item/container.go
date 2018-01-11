package item

import "fmt"

//glua:bind
type Container struct {
	capacity int
	slots    []*ItemStack
	listener ContainerListener
}

//glua:bind constructor Container
func NewContainer(capacity int) *Container {
	return &Container{
		capacity: capacity,
		slots:    make([]*ItemStack, capacity),
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

//glua:bind
func (c *Container) Slot(i int) *ItemStack {
	c.assertSize(i)
	c.assertSlotPopulated(true, i)
	return c.slots[i]
}

//glua:bind
func (c *Container) SetSlot(i int, item *ItemStack) {
	c.assertSize(i)
	c.assertSlotPopulated(false, i)
	c.slots[i] = item
	if c.listener != nil {
		c.listener.ContainerSlotSet(i, item)
	}
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

	if c.SlotPopulated(i) == populated {
		panic(fmt.Sprintf("expected slot %v to be %v", i, stateStr))
	}
}

//glua:bind
func (c *Container) SwapSlots(a, b int) {
	// Don't assert slots are populated, because we could be moving to an empty slot
	c.assertSize(a)
	c.assertSize(b)
	c.slots[a], c.slots[b] = c.slots[b], c.slots[a]
	if c.listener != nil {
		c.listener.ContainerSlotsSwapped(a, b)
	}
}

//glua:bind
func (c *Container) FindStackOf(id int) int {
	for i, stack := range c.slots {
		if stack.Id() == id {
			return i
		}
	}
	return -1
}

//glua:bind
func (c *Container) RemoveFromSlot(slot, count int) *ItemStack {
	stack := c.Slot(slot)
	if count > stack.Count() {
		panic(fmt.Sprintf("tried to remove %v from item stack of size %v", count, stack.Count()))
	}

	stack.count -= count
	if stack.Count() == 0 {
		c.slots[slot] = nil
	}

	newStack := NewItemStack(stack.Id(), count)
	return newStack
}

//glua:bind
func (c *Container) RemoveAllFromSlot(slot int) *ItemStack {
	stack := c.Slot(slot)
	return c.RemoveFromSlot(slot, stack.Count())
}
