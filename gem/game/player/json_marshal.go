package player

import (
	"github.com/gemrs/gem/gem/game/interface/player"
	"github.com/gemrs/gem/gem/game/position"
)

type jsonPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type jsonProfile struct {
	Username   string         `json:"username"`
	Password   string         `json:"password"`
	Rights     player.Rights  `json:"rights"`
	Position   jsonPosition   `json:"position"`
	Skills     jsonSkills     `json:"skills"`
	Appearance jsonAppearance `json:"appearance"`
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

			TorsoModel: appearance.Model(player.Torso),
			ArmsModel:  appearance.Model(player.Arms),
			LegsModel:  appearance.Model(player.Legs),
			HeadModel:  appearance.Model(player.Head),
			HandsModel: appearance.Model(player.Hands),
			FeetModel:  appearance.Model(player.Feet),
			BeardModel: appearance.Model(player.Beard),

			HairColor:  appearance.Color(player.Hair),
			TorsoColor: appearance.Color(player.Torso),
			LegsColor:  appearance.Color(player.Legs),
			FeetColor:  appearance.Color(player.Feet),
			SkinColor:  appearance.Color(player.Skin),
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

	skills := p.Skills().(*Skills)
	skills.SetCombatLevel(js.Skills.CombatLevel)

	appearance := p.Appearance().(*Appearance)
	appearance.SetGender(js.Appearance.Gender)
	appearance.SetHeadIcon(js.Appearance.HeadIcon)

	appearance.SetModel(player.Torso, js.Appearance.TorsoModel)
	appearance.SetModel(player.Arms, js.Appearance.ArmsModel)
	appearance.SetModel(player.Legs, js.Appearance.LegsModel)
	appearance.SetModel(player.Head, js.Appearance.HeadModel)
	appearance.SetModel(player.Hands, js.Appearance.HandsModel)
	appearance.SetModel(player.Feet, js.Appearance.FeetModel)
	appearance.SetModel(player.Beard, js.Appearance.BeardModel)

	appearance.SetColor(player.Hair, js.Appearance.HairColor)
	appearance.SetColor(player.Torso, js.Appearance.TorsoColor)
	appearance.SetColor(player.Legs, js.Appearance.LegsColor)
	appearance.SetColor(player.Feet, js.Appearance.FeetColor)
	appearance.SetColor(player.Skin, js.Appearance.SkinColor)
}
