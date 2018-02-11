//glua:bind module gem.game.impl
package impl

import (
	"encoding/json"
	"fmt"

	"github.com/gemrs/gem/gem/game/data"
	"github.com/gemrs/gem/gem/game/item"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/protocol"
)

//go:generate glua .

// Profile represents the saved state of a user
//glua:bind
type Profile struct {
	jsonProfile
}

type jsonProfile struct {
	Username string
	Password string
	Rights   protocol.Rights
	Position *position.Absolute

	Skills     *Skills
	Appearance *Appearance
	Inventory  *item.Container
	Equipment  *Equipment
}

//glua:bind constructor Profile
func NewProfile(username, password string) *Profile {
	profile := &Profile{
		jsonProfile: jsonProfile{
			Username:   username,
			Password:   password,
			Skills:     NewSkills(),
			Appearance: NewAppearance(),
			Inventory:  item.NewContainer(data.Int("inventory.inventory_size")),
			Equipment:  NewEquipment(),
		},
	}

	profile.jsonProfile.Inventory.SetInterfaceLocation(
		data.Int("widget.inventory_group_id"), 0,
		data.Int("inventory.inventory"))

	return profile
}

func (p *Profile) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.jsonProfile)
}

func (p *Profile) UnmarshalJSON(d []byte) error {
	return json.Unmarshal(d, &p.jsonProfile)
}

//glua:bind
func (p *Profile) Save() string {
	data, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(data)
}

//glua:bind
func (p *Profile) Load(data string) {
	err := json.Unmarshal([]byte(data), &p)
	if err != nil {
		panic(err)
	}
}

func (p *Profile) SetPlayer(player protocol.Player) {
	p.jsonProfile.Skills.setPlayer(player)
	p.jsonProfile.Appearance.setPlayer(player)
	p.jsonProfile.Equipment.setPlayer(player)
}

//glua:bind
func (p *Profile) Username() string {
	return p.jsonProfile.Username
}

//glua:bind
func (p *Profile) Password() string {
	return p.jsonProfile.Password
}

//glua:bind
func (p *Profile) Rights() protocol.Rights {
	return p.jsonProfile.Rights
}

func (p *Profile) setRights(rights protocol.Rights) {
	p.jsonProfile.Rights = rights
}

//glua:bind accessor
func (p *Profile) Position() *position.Absolute {
	return p.jsonProfile.Position
}

func (p *Profile) SetPosition(pos *position.Absolute) {
	p.jsonProfile.Position = pos
}

//glua:bind
func (p *Profile) Skills() protocol.Skills {
	return p.jsonProfile.Skills
}

//glua:bind
func (p *Profile) Inventory() *item.Container {
	return p.jsonProfile.Inventory
}

//glua:bind
func (p *Profile) Equipment() protocol.Equipment {
	return p.jsonProfile.Equipment
}

func (p *Profile) Appearance() protocol.Appearance {
	return p.jsonProfile.Appearance
}

func (p *Profile) String() string {
	return fmt.Sprintf("Username: %v", p.Username())
}
