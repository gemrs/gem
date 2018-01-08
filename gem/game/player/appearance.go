package player

type BodyPart int

const (
	BodyPartTorso BodyPart = iota
	BodyPartArms
	BodyPartLegs
	BodyPartHead
	BodyPartHands
	BodyPartFeet
	BodyPartBeard
	BodyPartHair
	BodyPartSkin
	BodyPartMax
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

func (a *Appearance) Model(b BodyPart) int {
	switch b {
	case BodyPartTorso:
		return a.torsoModel
	case BodyPartArms:
		return a.armsModel
	case BodyPartLegs:
		return a.legsModel
	case BodyPartHead:
		return a.headModel
	case BodyPartHands:
		return a.handsModel
	case BodyPartFeet:
		return a.feetModel
	case BodyPartBeard:
		return a.beardModel
	}
	return -1
}

func (a *Appearance) SetModel(b BodyPart, model int) {
	switch b {
	case BodyPartTorso:
		a.torsoModel = model
	case BodyPartArms:
		a.armsModel = model
	case BodyPartLegs:
		a.legsModel = model
	case BodyPartHead:
		a.headModel = model
	case BodyPartHands:
		a.handsModel = model
	case BodyPartFeet:
		a.feetModel = model
	case BodyPartBeard:
		a.beardModel = model
	}
	a.signalUpdate()
}

func (a *Appearance) Color(b BodyPart) int {
	switch b {
	case BodyPartHair:
		return a.hairColor
	case BodyPartTorso:
		return a.torsoColor
	case BodyPartLegs:
		return a.legsColor
	case BodyPartFeet:
		return a.feetColor
	case BodyPartSkin:
		return a.skinColor
	}
	return -1
}

func (a *Appearance) SetColor(b BodyPart, color int) {
	switch b {
	case BodyPartHair:
		a.hairColor = color
	case BodyPartTorso:
		a.torsoColor = color
	case BodyPartLegs:
		a.legsColor = color
	case BodyPartFeet:
		a.feetColor = color
	case BodyPartSkin:
		a.skinColor = color
	}
	a.signalUpdate()
}
