package player

import (
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/protocol"
)

type jsonPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type jsonProfile struct {
	Username   string          `json:"username"`
	Password   string          `json:"password"`
	Rights     protocol.Rights `json:"rights"`
	Position   jsonPosition    `json:"position"`
	Skills     jsonSkills      `json:"skills"`
	Appearance jsonAppearance  `json:"appearance"`
}

type jsonSkills struct {
	CombatLevel int `json:"combat_level"`
}

type jsonAppearance struct {
	Gender   int `json:"gender"`
	HeadIcon int `json:"head_icon"`

	TorsoModel int `json:"model_torso"`
	ArmsModel  int `json:"model_arms"`
	LegsModel  int `json:"model_legs"`
	HeadModel  int `json:"model_head"`
	HandsModel int `json:"model_hands"`
	FeetModel  int `json:"model_feet"`
	BeardModel int `json:"model_beard"`

	HairColor  int `json:"color_hair"`
	TorsoColor int `json:"color_torso"`
	LegsColor  int `json:"color_legs"`
	FeetColor  int `json:"color_feet"`
	SkinColor  int `json:"color_skin"`
}

func jsonObjForProfile(p *Profile) jsonProfile {
	pos := p.Position()
	skills := p.Skills()
	appearance := p.Appearance()
	obj := jsonProfile{
		Username: p.Username(),
		Password: p.Password(),
		Rights:   p.Rights(),
		Position: jsonPosition{
			pos.X(), pos.Y(), pos.Z(),
		},
		Skills: jsonSkills{
			CombatLevel: skills.CombatLevel(),
		},
		Appearance: jsonAppearance{
			Gender:   appearance.Gender(),
			HeadIcon: appearance.HeadIcon(),

			TorsoModel: appearance.Model(protocol.BodyPartTorso),
			ArmsModel:  appearance.Model(protocol.BodyPartArms),
			LegsModel:  appearance.Model(protocol.BodyPartLegs),
			HeadModel:  appearance.Model(protocol.BodyPartHead),
			HandsModel: appearance.Model(protocol.BodyPartHands),
			FeetModel:  appearance.Model(protocol.BodyPartFeet),
			BeardModel: appearance.Model(protocol.BodyPartBeard),

			HairColor:  appearance.Color(protocol.BodyPartHair),
			TorsoColor: appearance.Color(protocol.BodyPartTorso),
			LegsColor:  appearance.Color(protocol.BodyPartLegs),
			FeetColor:  appearance.Color(protocol.BodyPartFeet),
			SkinColor:  appearance.Color(protocol.BodyPartSkin),
		},
	}

	return obj
}

func jsonObjToProfile(p *Profile, js jsonProfile) {
	/*
		Don't set the username or password, because the Profile will have been constructed with the correct values
			p.setUsername(js.Username)
			p.setPassword(js.Password)
	*/

	p.setRights(js.Rights)
	p.SetPosition(position.NewAbsolute(js.Position.X, js.Position.Y, js.Position.Z))

	skills := p.Skills()
	skills.SetCombatLevel(js.Skills.CombatLevel)

	appearance := p.Appearance()
	appearance.SetGender(js.Appearance.Gender)
	appearance.SetHeadIcon(js.Appearance.HeadIcon)

	appearance.SetModel(protocol.BodyPartTorso, js.Appearance.TorsoModel)
	appearance.SetModel(protocol.BodyPartArms, js.Appearance.ArmsModel)
	appearance.SetModel(protocol.BodyPartLegs, js.Appearance.LegsModel)
	appearance.SetModel(protocol.BodyPartHead, js.Appearance.HeadModel)
	appearance.SetModel(protocol.BodyPartHands, js.Appearance.HandsModel)
	appearance.SetModel(protocol.BodyPartFeet, js.Appearance.FeetModel)
	appearance.SetModel(protocol.BodyPartBeard, js.Appearance.BeardModel)

	appearance.SetColor(protocol.BodyPartHair, js.Appearance.HairColor)
	appearance.SetColor(protocol.BodyPartTorso, js.Appearance.TorsoColor)
	appearance.SetColor(protocol.BodyPartLegs, js.Appearance.LegsColor)
	appearance.SetColor(protocol.BodyPartFeet, js.Appearance.FeetColor)
	appearance.SetColor(protocol.BodyPartSkin, js.Appearance.SkinColor)
}
