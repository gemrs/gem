package player

import (
	"github.com/sinusoids/gem/gem/game/position"
)

type Rights int

const (
	RightsPlayer Rights = iota
	RightsModerator
	RightsAdmin
)

type Profile interface {
	Username() string
	Password() string
	Rights() Rights
	Position() *position.Absolute
	SetPosition(*position.Absolute)
	Skills() Skills
	Appearance() Appearance
	SetAppearance(Appearance)
	Animations() Animations
}

type Skills interface {
	CombatLevel() int
}
