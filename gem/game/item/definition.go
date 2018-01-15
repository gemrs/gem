package item

import (
	"github.com/gemrs/gem/gem/runite"
	"github.com/gemrs/gem/gem/runite/format/rt3"
)

var definitions []Definition

//glua:bind
type Definition struct {
	data rt3.ItemDefinition
}

//glua:bind constructor Definition
func NewDefinition(id int) *Definition {
	return &definitions[id]
}

func LoadDefinitions(ctx *runite.Context) {
	definitions = make([]Definition, len(ctx.ItemDefinitions))
	for i, def := range ctx.ItemDefinitions {
		definitions[i].data = *def
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
	return d.data.Description
}

func (d *Definition) Actions() []string {
	return d.data.Actions
}

func (d *Definition) GroundActions() []string {
	return d.data.GroundActions
}

//glua:bind
func (d *Definition) Noted() bool {
	return d.data.Noted
}

//glua:bind
func (d *Definition) Notable() bool {
	return d.data.Notable
}

//glua:bind
func (d *Definition) Stackable() bool {
	return d.data.Stackable
}

//glua:bind
func (d *Definition) ParentId() int {
	return d.data.ParentId
}

//glua:bind
func (d *Definition) NotedId() int {
	return d.data.NotedId
}

//glua:bind
func (d *Definition) Members() bool {
	return d.data.Members
}

//glua:bind
func (d *Definition) ShopValue() int {
	return d.data.ShopValue
}
