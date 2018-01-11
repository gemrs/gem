//glua:bind module gem.game.item
package item

//go:generate glua .

//glua:bind
type ItemStack struct {
	id    int
	count int
}

//glua:bind constructor ItemStack
func NewItemStack(id, count int) *ItemStack {
	return &ItemStack{
		id:    id,
		count: count,
	}
}

//glua:bind
func (i *ItemStack) Id() int {
	return i.id
}

//glua:bind
func (i *ItemStack) Count() int {
	return i.count
}
