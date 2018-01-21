package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type weaponData struct {
	Id           int
	TwoHanded    bool `json:"two_handed"`
	Type         string
	AttackAnims  []int      `json:"attack_anims"`
	AttackSpeeds []int      `json:"attack_speeds"`
	CharAnims    Animations `json:"char_anims"`
	TabInterface int        `json:"tab_interface"`
}

type WeaponDef struct {
	data weaponData
}

func (d WeaponDef) CharAnimations() *Animations {
	return &d.data.CharAnims
}

func (d WeaponDef) AttackTabInterface() int {
	return d.data.TabInterface
}

type Animations struct {
	Idle       int `json:"idle"`
	SpotRotate int `json:"spot_rotate"`
	Walk       int `json:"walk"`
	Rotate180  int `json:"rotate_180"`
	RotateCCW  int `json:"rotate_ccw"`
	RotateCW   int `json:"rotate_cw"`
	Run        int `json:"run"`
}

func NewAnimations() *Animations {
	return &Animations{
		Idle:       0x328,
		SpotRotate: 0x337,
		Walk:       0x333,
		Rotate180:  0x334,
		RotateCCW:  0x335,
		RotateCW:   0x336,
		Run:        0x338,
	}
}

var DefaultAnimations = NewAnimations()

var Weapons = map[int]WeaponDef{}

//glua:bind
func LoadWeaponData(path string) error {
	var data []weaponData
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
		Weapons[d.Id] = WeaponDef{d}
	}

	return nil
}
