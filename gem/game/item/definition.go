package item

import (
	"encoding/json"
	"io/ioutil"
)

var definitions []Definition

type definitionModel struct {
	Id        int
	Name      string
	Examine   string
	Noted     bool
	Notable   bool
	Stackable bool
	ParentId  int
	NotedId   int
	Members   bool
	ShopValue int
}

//glua:bind
type Definition struct {
	data definitionModel
}

//glua:bind constructor Definition
func NewDefinition(id int) *Definition {
	return &definitions[id]
}

func (d *Definition) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &d.data)
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
func (d *Definition) Examine() string {
	return d.data.Examine
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

//glua:bind
func LoadDefinitions(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &definitions)
	if err != nil {
		return err
	}

	return nil
}
