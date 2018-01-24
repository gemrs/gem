package player

import "github.com/gemrs/gem/gem/protocol"

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

func (a *Appearance) Model(b protocol.BodyPart) int {
	switch b {
	case protocol.BodyPartTorso:
		return a.torsoModel
	case protocol.BodyPartArms:
		return a.armsModel
	case protocol.BodyPartLegs:
		return a.legsModel
	case protocol.BodyPartHead:
		return a.headModel
	case protocol.BodyPartHands:
		return a.handsModel
	case protocol.BodyPartFeet:
		return a.feetModel
	case protocol.BodyPartBeard:
		return a.beardModel
	}
	return -1
}

func (a *Appearance) SetModel(b protocol.BodyPart, model int) {
	switch b {
	case protocol.BodyPartTorso:
		a.torsoModel = model
	case protocol.BodyPartArms:
		a.armsModel = model
	case protocol.BodyPartLegs:
		a.legsModel = model
	case protocol.BodyPartHead:
		a.headModel = model
	case protocol.BodyPartHands:
		a.handsModel = model
	case protocol.BodyPartFeet:
		a.feetModel = model
	case protocol.BodyPartBeard:
		a.beardModel = model
	}
	a.signalUpdate()
}

func (a *Appearance) Color(b protocol.BodyPart) int {
	switch b {
	case protocol.BodyPartHair:
		return a.hairColor
	case protocol.BodyPartTorso:
		return a.torsoColor
	case protocol.BodyPartLegs:
		return a.legsColor
	case protocol.BodyPartFeet:
		return a.feetColor
	case protocol.BodyPartSkin:
		return a.skinColor
	}
	return -1
}

func (a *Appearance) SetColor(b protocol.BodyPart, color int) {
	switch b {
	case protocol.BodyPartHair:
		a.hairColor = color
	case protocol.BodyPartTorso:
		a.torsoColor = color
	case protocol.BodyPartLegs:
		a.legsColor = color
	case protocol.BodyPartFeet:
		a.feetColor = color
	case protocol.BodyPartSkin:
		a.skinColor = color
	}
	a.signalUpdate()
}
