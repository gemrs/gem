package player

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/game/interface/player"
)

type Appearance struct {
	py.BaseObject

	gender   int
	headIcon int

	torsoModel int
	armsModel  int
	legsModel  int
	headModel  int
	handsModel int
	feetModel  int
	beardModel int

	hairColor  int
	torsoColor int
	legsColor  int
	feetColor  int
	skinColor  int
}

func (a *Appearance) Init() {
	a.gender = 0
	a.headIcon = 0

	a.torsoModel = 19
	a.armsModel = 29
	a.legsModel = 39
	a.headModel = 3
	a.handsModel = 35
	a.feetModel = 44
	a.beardModel = 10

	a.hairColor = 7
	a.torsoColor = 8
	a.legsColor = 9
	a.feetColor = 5
	a.skinColor = 0
}

func (a *Appearance) Gender() int {
	return a.gender
}

func (a *Appearance) setGender(gender int) {
	a.gender = gender
}

func (a *Appearance) HeadIcon() int {
	return a.headIcon
}

func (a *Appearance) setHeadIcon(headIcon int) {
	a.headIcon = headIcon
}

func (a *Appearance) Model(b player.BodyPart) int {
	switch b {
	case player.Torso:
		return a.torsoModel
	case player.Arms:
		return a.armsModel
	case player.Legs:
		return a.legsModel
	case player.Head:
		return a.headModel
	case player.Hands:
		return a.handsModel
	case player.Feet:
		return a.feetModel
	case player.Beard:
		return a.beardModel
	}
	return -1
}

func (a *Appearance) setModel(b player.BodyPart, model int) {
	switch b {
	case player.Torso:
		a.torsoModel = model
	case player.Arms:
		a.armsModel = model
	case player.Legs:
		a.legsModel = model
	case player.Head:
		a.headModel = model
	case player.Hands:
		a.handsModel = model
	case player.Feet:
		a.feetModel = model
	case player.Beard:
		a.beardModel = model
	}
}

func (a *Appearance) Color(b player.BodyPart) int {
	switch b {
	case player.Hair:
		return a.hairColor
	case player.Torso:
		return a.torsoColor
	case player.Legs:
		return a.legsColor
	case player.Feet:
		return a.feetColor
	case player.Skin:
		return a.skinColor
	}
	return -1
}

func (a *Appearance) setColor(b player.BodyPart, color int) {
	switch b {
	case player.Hair:
		a.hairColor = color
	case player.Torso:
		a.torsoColor = color
	case player.Legs:
		a.legsColor = color
	case player.Feet:
		a.feetColor = color
	case player.Skin:
		a.skinColor = color
	}
}

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

func (a *Animations) Init() {
	a.idle = 0x328
	a.spotRotate = 0x337
	a.walk = 0x333
	a.rotate180 = 0x334
	a.rotateCCW = 0x335
	a.rotateCW = 0x336
	a.run = 0x338
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
