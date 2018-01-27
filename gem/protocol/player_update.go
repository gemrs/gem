package protocol

import (
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/position"
)

type PlayerUpdateData struct {
	Index            int
	Username         string
	Flags            entity.Flags
	Position         *position.Absolute
	LoadedRegion     *position.Region
	WaypointQueue    entity.WaypointQueue
	ChatMessageQueue []InboundChatMessage
	Rights           Rights
	Appearance       Appearance
	Animations       Animations
	Skills           Skills

	// Some protocol specific data
	ProtoData interface{}
}

// +gen pack_outgoing
type PlayerUpdate struct {
	Me       PlayerUpdateData
	Visible  []PlayerUpdateData
	Others   map[int]PlayerUpdateData
	Removing map[int]bool
	Updating []int
	Adding   []int
}

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

type Appearance interface {
	Gender() int
	HeadIcon() int
	Model(b BodyPart) int
	SetModel(b BodyPart, model int)
	Color(b BodyPart) int
}

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

type Animations interface {
	Animation(anim Anim) int
}

//glua:bind
type SkillId int

//glua:bind
const (
	SkillAttack SkillId = iota
	SkillDefence
	SkillStrength
	SkillHitpoints
	SkillRange
	SkillPrayer
	SkillMagic
	SkillCooking
	SkillWoodcutting
	SkillFletching
	SkillFishing
	SkillFiremaking
	SkillCrafting
	SkillSmithing
	SkillMining
	SkillHerblore
	SkillAgility
	SkillThieving
	SkillSlayer
	SkillFarming
	SkillRunecrafting
)

type Skills interface {
	CombatLevel() int
	//	Skill(id SkillId) *Skill
}
