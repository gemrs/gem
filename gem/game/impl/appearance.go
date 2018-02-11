package impl

import (
	"encoding/json"

	"github.com/gemrs/gem/gem/protocol"
)

type Appearance struct {
	player protocol.Player

	jsonAppearance
}

type jsonAppearance struct {
	Gender    int
	headIcon  int
	skullIcon int

	TorsoModel int
	ArmsModel  int
	LegsModel  int
	HeadModel  int
	HandsModel int
	FeetModel  int
	BeardModel int

	HairColor  int
	TorsoColor int
	LegsColor  int
	FeetColor  int
	SkinColor  int
}

func NewAppearance() *Appearance {
	a := &Appearance{}
	a.jsonAppearance.Gender = 0
	a.jsonAppearance.headIcon = -1
	a.jsonAppearance.skullIcon = -1

	a.TorsoModel = 19
	a.ArmsModel = 29
	a.LegsModel = 39
	a.HeadModel = 3
	a.HandsModel = 35
	a.FeetModel = 44
	a.BeardModel = 10

	a.HairColor = 7
	a.TorsoColor = 8
	a.LegsColor = 9
	a.FeetColor = 5
	a.SkinColor = 0
	return a
}

func (a *Appearance) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.jsonAppearance)
}

func (a *Appearance) UnmarshalJSON(d []byte) error {
	return json.Unmarshal(d, &a.jsonAppearance)
}

func (a *Appearance) setPlayer(p protocol.Player) {
	a.player = p
	a.signalUpdate()
}

func (a *Appearance) signalUpdate() {
	if a.player != nil {
		a.player.SetAppearanceChanged()
	}
}

func (a *Appearance) Gender() int {
	return a.jsonAppearance.Gender
}

func (a *Appearance) SetGender(gender int) {
	a.jsonAppearance.Gender = gender
	a.signalUpdate()
}

func (a *Appearance) SkullIcon() int {
	return a.jsonAppearance.skullIcon
}

func (a *Appearance) SetSkullIcon(skullIcon int) {
	a.skullIcon = skullIcon
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
		return a.TorsoModel
	case protocol.BodyPartArms:
		return a.ArmsModel
	case protocol.BodyPartLegs:
		return a.LegsModel
	case protocol.BodyPartHead:
		return a.HeadModel
	case protocol.BodyPartHands:
		return a.HandsModel
	case protocol.BodyPartFeet:
		return a.FeetModel
	case protocol.BodyPartBeard:
		return a.BeardModel
	}
	return -1
}

func (a *Appearance) SetModel(b protocol.BodyPart, model int) {
	switch b {
	case protocol.BodyPartTorso:
		a.TorsoModel = model
	case protocol.BodyPartArms:
		a.ArmsModel = model
	case protocol.BodyPartLegs:
		a.LegsModel = model
	case protocol.BodyPartHead:
		a.HeadModel = model
	case protocol.BodyPartHands:
		a.HandsModel = model
	case protocol.BodyPartFeet:
		a.FeetModel = model
	case protocol.BodyPartBeard:
		a.BeardModel = model
	}
	a.signalUpdate()
}

func (a *Appearance) Color(b protocol.BodyPart) int {
	switch b {
	case protocol.BodyPartHair:
		return a.HairColor
	case protocol.BodyPartTorso:
		return a.TorsoColor
	case protocol.BodyPartLegs:
		return a.LegsColor
	case protocol.BodyPartFeet:
		return a.FeetColor
	case protocol.BodyPartSkin:
		return a.SkinColor
	}
	return -1
}

func (a *Appearance) SetColor(b protocol.BodyPart, color int) {
	switch b {
	case protocol.BodyPartHair:
		a.HairColor = color
	case protocol.BodyPartTorso:
		a.TorsoColor = color
	case protocol.BodyPartLegs:
		a.LegsColor = color
	case protocol.BodyPartFeet:
		a.FeetColor = color
	case protocol.BodyPartSkin:
		a.SkinColor = color
	}
	a.signalUpdate()
}
