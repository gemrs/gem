package player

type Anim int

const (
	AnimIdle Anim = iota
	AnimSpotRotate
	AnimWalk
	AnimRotate180
	AnimRotateCCW
	AnimRotateCW
	AnimRun
	AnimMax
)

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

func (a *Animations) Animation(anim Anim) int {
	switch anim {
	case AnimIdle:
		return a.idle
	case AnimSpotRotate:
		return a.spotRotate
	case AnimWalk:
		return a.walk
	case AnimRotate180:
		return a.rotate180
	case AnimRotateCCW:
		return a.rotateCCW
	case AnimRotateCW:
		return a.rotateCW
	case AnimRun:
		return a.run
	}
	panic("not reached")
}
