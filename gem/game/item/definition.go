package item

import (
	"github.com/gemrs/gem/gem/game/data"
)

//glua:bind
type Definition struct {
	data data.ItemDefinition
}

//glua:bind constructor Definition
func NewDefinition(id int) *Definition {
	return &Definition{
		data: data.Item(id),
	}
}

//glua:bind
func (d *Definition) Id() int {
	return d.data.Id
}

//glua:bind
func (d *Definition) Name() string {
	return d.data.Name
}

//glua:bind
func (d *Definition) Description() string {
	return "no description!"
}

func (d *Definition) Actions() []string {
	return d.data.Actions
}

func (d *Definition) GroundActions() []string {
	return d.data.GroundActions
}

//glua:bind
func (d *Definition) Stackable() bool {
	return d.data.Stackable == 1 || d.data.NotedTemplate >= 0
}

//glua:bind
func (d *Definition) NotedId() int {
	return d.data.NotedId
}

//glua:bind
func (d *Definition) ShopValue() int {
	return d.data.ShopValue
}
