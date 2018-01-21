package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//go:generate glua .

type equipmentData struct {
	Id           int
	Slot         int
	Bonuses      []int
	Requirements []int
	CoversFace   bool `json:"covers_face"`
	CoversHair   bool `json:"covers_hair"`
	CoversArms   bool `json:"covers_arms"`
}

type EquipmentDef struct {
	data equipmentData
}

func (d *EquipmentDef) Id() int {
	return d.data.Id
}

func (d *EquipmentDef) Slot() int {
	return d.data.Slot
}

func (d *EquipmentDef) Bonuses() []int {
	return d.data.Bonuses
}

func (d *EquipmentDef) Requirements() []int {
	return d.data.Requirements
}

func (d *EquipmentDef) CoversFace() bool {
	return d.data.CoversFace
}

func (d *EquipmentDef) CoversHair() bool {
	return d.data.CoversHair
}

func (d *EquipmentDef) CoversArms() bool {
	return d.data.CoversArms
}

var Equipment = map[int]EquipmentDef{}

//glua:bind
func LoadEquipmentData(path string) error {
	var data []equipmentData
	fd, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fd.Close()

	jsonData, err := ioutil.ReadAll(fd)
	if err != nil {
		return err
	}

	json.Unmarshal(jsonData, &data)
	for _, d := range data {
		Equipment[d.Id] = EquipmentDef{d}
	}

	return nil
}
