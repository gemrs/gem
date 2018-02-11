//glua:bind module gem.game.item
package item

import "encoding/json"

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

type jsonStack struct {
	Id    int
	Count int
}

func (s *Stack) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonStack{
		Id:    s.definition.Id(),
		Count: s.count,
	})
}

func (s *Stack) UnmarshalJSON(d []byte) error {
	var stack jsonStack
	err := json.Unmarshal(d, &stack)
	if err == nil {
		s.definition = NewDefinition(stack.Id)
		s.count = stack.Count
	}
	return err
}

//glua:bind
func (i *Stack) Definition() *Definition {
	return i.definition
}

//glua:bind
func (i *Stack) Count() int {
	return i.count
}
