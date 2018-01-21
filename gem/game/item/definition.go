package item

import (
	"github.com/gemrs/gem/gem/game/data"
	"github.com/gemrs/gem/gem/runite/format/rt3"
)

//glua:bind
type Definition struct {
	itemData      *rt3.ItemDefinition
	weaponData    data.WeaponDef
	equipmentData data.EquipmentDef
}

//glua:bind constructor Definition
func NewDefinition(id int) *Definition {
	itemData, err := data.Config.Item(id)
	if err != nil {
		return nil
	}

	definition := &Definition{
		itemData: itemData,
	}

	if weaponData, ok := data.Weapons[id]; ok {
		definition.weaponData = weaponData
	}

	if equipmentData, ok := data.Equipment[id]; ok {
		definition.equipmentData = equipmentData
	}

	return definition
}

//glua:bind
func (d *Definition) Id() int {
	return d.itemData.Id
}

//glua:bind
func (d *Definition) Name() string {
	return d.itemData.Name
}

//glua:bind
func (d *Definition) Description() string {
	return "no description!"
}

func (d *Definition) Actions() []string {
	return d.itemData.InventoryActions[:]
}

func (d *Definition) GroundActions() []string {
	return d.itemData.GroundActions[:]
}

//glua:bind
func (d *Definition) Stackable() bool {
	return d.itemData.Stackable || d.itemData.NotedTemplate >= 0
}

//glua:bind
func (d *Definition) NotedId() int {
	return d.itemData.NotedId
}

//glua:bind
func (d *Definition) ShopValue() int {
	return d.itemData.ShopValue
}

func (d *Definition) WeaponData() *data.WeaponDef {
	return &d.weaponData
}

func (d *Definition) EquipmentData() *data.EquipmentDef {
	return &d.equipmentData
}
