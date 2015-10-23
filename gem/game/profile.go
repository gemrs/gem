package game

import (
	"fmt"

	"github.com/qur/gopy/lib"

	"gem/game/player"
	"gem/game/position"
)

//go:generate gopygen -type Profile -type Skills $GOFILE

// Profile represents the saved state of a user
type Profile struct {
	py.BaseObject

	username string
	password string
	rights   player.Rights
	position *position.Absolute

	skills     *Skills
	appearance *Appearance
	animations *Animations
}

func (p *Profile) Init(username, password string) (err error) {
	p.username = username
	p.password = password

	p.skills, err = NewSkills()
	if err != nil {
		return err
	}

	p.appearance, err = NewAppearance()
	if err != nil {
		return err
	}

	p.animations, err = NewAnimations()
	if err != nil {
		return err
	}

	return nil
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

func (p *Profile) Animations() player.Animations {
	return p.animations
}

func (p *Profile) String() string {
	return fmt.Sprintf("Username: %v", p.username)
}

type Skills struct {
	py.BaseObject

	combatLevel int
}

func (s *Skills) Init() error {
	return nil
}

func (s *Skills) CombatLevel() int {
	return s.combatLevel
}
