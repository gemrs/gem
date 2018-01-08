package player

import (
	"github.com/gemrs/gem/gem/game/interface/player"
)

type Appearance struct {
	player *Player

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

func NewAppearance() *Appearance {
	a := &Appearance{}
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
	return a
}

func (a *Appearance) setPlayer(p *Player) {
	a.player = p
	a.signalUpdate()
}

func (a *Appearance) signalUpdate() {
	if a.player != nil {
		a.player.SetAppearanceChanged()
	}
}

func (a *Appearance) Gender() int {
	return a.gender
}

func (a *Appearance) SetGender(gender int) {
	a.gender = gender
	a.signalUpdate()
}

func (a *Appearance) HeadIcon() int {
	return a.headIcon
}

func (a *Appearance) SetHeadIcon(headIcon int) {
	a.headIcon = headIcon
	a.signalUpdate()
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

func (a *Appearance) SetModel(b player.BodyPart, model int) {
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
	a.signalUpdate()
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

func (a *Appearance) SetColor(b player.BodyPart, color int) {
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
	a.signalUpdate()
}
