//glua:bind module gem.game.item
package item

//go:generate glua .

//glua:bind
type Stack struct {
	definition *Definition
	count      int
}

//glua:bind constructor Stack
func NewStack(definition *Definition, count int) *Stack {
	return &Stack{
		definition: definition,
		count:      count,
	}
}

//glua:bind
func (i *Stack) Definition() *Definition {
	return i.definition
}

//glua:bind
func (i *Stack) Count() int {
	return i.count
}
