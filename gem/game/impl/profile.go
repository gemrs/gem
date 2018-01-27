//glua:bind module gem.game.impl
package impl

import (
	"fmt"

	game_event "github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/item"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/protocol"
)

func (player *Player) LoadProfile() {
	profile := player.profile
	profile.SetPlayer(player)
	player.SetPosition(profile.Position())

	game_event.PlayerLoadProfile.NotifyObservers(player, player.Profile())
}

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
}

//glua:bind constructor Profile
func NewProfile(username, password string) *Profile {
	return &Profile{
		username:   username,
		password:   password,
		skills:     NewSkills(),
		appearance: NewAppearance(),
		inventory:  item.NewContainer(28),
	}
}

func (p *Profile) SetPlayer(player protocol.Player) {
	p.skills.setPlayer(player)
	p.appearance.setPlayer(player)
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

func (p *Profile) Appearance() protocol.Appearance {
	return p.appearance
}

func (p *Profile) String() string {
	return fmt.Sprintf("Username: %v", p.username)
}
