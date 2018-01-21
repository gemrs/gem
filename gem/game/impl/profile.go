//glua:bind module gem.game.impl
package impl

import (
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
	username string
	password string
	rights   protocol.Rights
	position *position.Absolute

	skills     *Skills
	appearance *Appearance
	inventory  *item.Container
	equipment  *Equipment
}

//glua:bind constructor Profile
func NewProfile(username, password string) *Profile {
	profile := &Profile{
		username:   username,
		password:   password,
		skills:     NewSkills(),
		appearance: NewAppearance(),
		inventory:  item.NewContainer(data.Int("inventory.inventory_size")),
		equipment:  NewEquipment(),
	}

	profile.inventory.SetInterfaceLocation(
		data.Int("widget.inventory_group_id"), 0,
		data.Int("inventory.inventory"))

	return profile
}

func (p *Profile) SetPlayer(player protocol.Player) {
	p.skills.setPlayer(player)
	p.appearance.setPlayer(player)
	p.equipment.setPlayer(player)
}

//glua:bind
func (p *Profile) Username() string {
	return p.username
}

//glua:bind
func (p *Profile) Password() string {
	return p.password
}

//glua:bind
func (p *Profile) Rights() protocol.Rights {
	return p.rights
}

func (p *Profile) setRights(rights protocol.Rights) {
	p.rights = rights
}

//glua:bind accessor
func (p *Profile) Position() *position.Absolute {
	return p.position
}

func (p *Profile) SetPosition(pos *position.Absolute) {
	p.position = pos
}

//glua:bind
func (p *Profile) Skills() protocol.Skills {
	return p.skills
}

//glua:bind
func (p *Profile) Inventory() *item.Container {
	return p.inventory
}

//glua:bind
func (p *Profile) Equipment() protocol.Equipment {
	return p.equipment
}

func (p *Profile) Appearance() protocol.Appearance {
	return p.appearance
}

func (p *Profile) String() string {
	return fmt.Sprintf("Username: %v", p.username)
}
