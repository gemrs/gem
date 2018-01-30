package entity

import (
	"github.com/gemrs/gem/gem/game/position"
)

type WaypointQueue interface {
	Interaction
	Empty() bool
	Clear()
	Push(point *position.Absolute)
	SetRunning(running bool)
	WalkDirection() (current int, last int)
}
