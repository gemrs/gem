package player

import "github.com/gemrs/gem/gem/protocol"

type Animations struct {
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

func (a *Animations) Animation(anim protocol.Anim) int {
	switch anim {
	case protocol.AnimIdle:
		return a.idle
	case protocol.AnimSpotRotate:
		return a.spotRotate
	case protocol.AnimWalk:
		return a.walk
	case protocol.AnimRotate180:
		return a.rotate180
	case protocol.AnimRotateCCW:
		return a.rotateCCW
	case protocol.AnimRotateCW:
		return a.rotateCW
	case protocol.AnimRun:
		return a.run
	}
	panic("not reached")
}
