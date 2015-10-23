package player

import (
	"github.com/sinusoids/gem/gem/game/entity"
	"github.com/sinusoids/gem/gem/game/position"
)

type Session interface {
	Flags() entity.Flags
	Region() *position.Region
	WalkDirection() (current int, last int)
}
