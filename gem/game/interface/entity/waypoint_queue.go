package entity

import (
	"github.com/gemrs/gem/gem/game/position"
)

type WaypointQueue interface {
	Empty() bool
	Clear()
	Push(point *position.Absolute)
	Tick(mob Movable)
	WalkDirection() (current int, last int)
}
