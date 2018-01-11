package item

type ContainerListener interface {
	ContainerSlotSet(i int, item *ItemStack)
	ContainerSlotsSwapped(a, b int)
}
