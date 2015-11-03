package player

import (
	"fmt"

	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/game/interface/player"
	"github.com/sinusoids/gem/gem/game/position"
)

// Profile represents the saved state of a user
type Profile struct {
	py.BaseObject `json:"-"`

	username string             `json:"username"`
	password string             `json:"password"`
	rights   player.Rights      `json:"rights"`
	position *position.Absolute `json:"position"`

	skills     *Skills     `json:"skills"`
	appearance *Appearance `json:"appearance"`
}

func (p *Profile) Init(username, password string) {
	p.username = username
	p.password = password

	p.skills = NewSkills()
	p.appearance = NewAppearance()
}

func (p *Profile) Username() string {
	return p.username
}

func (p *Profile) Password() string {
	return p.password
}

func (p *Profile) Rights() player.Rights {
	return p.rights
}

func (p *Profile) Position() *position.Absolute {
	return p.position
}

func (p *Profile) SetPosition(pos *position.Absolute) {
	p.position = pos
}

func (p *Profile) Skills() player.Skills {
	return p.skills
}

func (p *Profile) Appearance() player.Appearance {
	return p.appearance
}

func (p *Profile) SetAppearance(appearance player.Appearance) {
	p.appearance = appearance.(*Appearance)
}

func (p *Profile) String() string {
	return fmt.Sprintf("Username: %v", p.username)
}

type Skills struct {
	py.BaseObject

	combatLevel int
}

func (s *Skills) Init() {}

func (s *Skills) CombatLevel() int {
	return s.combatLevel
}
