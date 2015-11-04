package player

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/game/interface/player"
)

type Animations struct {
	py.BaseObject

	idle       int
	spotRotate int
	walk       int
	rotate180  int
	rotateCCW  int
	rotateCW   int
	run        int
}

func NewAnimations() *Animations {
	return &Animations{
		idle:       0x328,
		spotRotate: 0x337,
		walk:       0x333,
		rotate180:  0x334,
		rotateCCW:  0x335,
		rotateCW:   0x336,
		run:        0x338,
	}
}

func (a *Animations) Animation(anim player.Anim) int {
	switch anim {
	case player.AnimIdle:
		return a.idle
	case player.AnimSpotRotate:
		return a.spotRotate
	case player.AnimWalk:
		return a.walk
	case player.AnimRotate180:
		return a.rotate180
	case player.AnimRotateCCW:
		return a.rotateCCW
	case player.AnimRotateCW:
		return a.rotateCW
	case player.AnimRun:
		return a.run
	}
	panic("not reached")
}
