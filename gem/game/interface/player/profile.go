package player

import (
	"github.com/gemrs/gem/gem/game/position"
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
}

type Skills interface {
	CombatLevel() int
}
