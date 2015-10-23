package player

import (
	"gem/game/entity"
	"gem/game/position"
)

type Session interface {
	Flags() entity.Flags
	Region() *position.Region
	WalkDirection() (current int, last int)
}
