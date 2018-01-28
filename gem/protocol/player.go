package protocol

import (
	"github.com/gemrs/gem/fork/github.com/gtank/isaac"
	"github.com/gemrs/gem/gem/core/crypto"
	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/item"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/willow/log"
)

type LoginHandler interface {
	SetRsaKeypair(keypair *crypto.Keypair)
	SetServerIsaacSeed(seed []uint32)
	SetCompleteCallback(cb func(LoginHandler) error)
	Perform(server.GameClient)

	/* results */
	Username() string
	Password() string
	IsaacSeeds() (in []uint32, out []uint32)
}

type Profile interface {
	Username() string
	Rights() Rights
	Inventory() *item.Container
	Appearance() Appearance
	Skills() Skills
	Position() *position.Absolute
	SetPosition(*position.Absolute)
	SetPlayer(Player)
}

type Player interface {
	AppendChatMessage(m InboundChatMessage)
	Profile() Profile
	WorldInstance() World
	Position() *position.Absolute
	InteractionQueue() *entity.InteractionQueue
	WaypointQueue() entity.WaypointQueue
	Log() log.Log
	CurrentFrame() FrameType
	SetCurrentFrame(FrameType)
	SetAppearanceChanged()

	Conn() *server.Connection
	Decode() error
	Encode(encoding.Encodable) error
	Disconnect()
	SetProtoData(d interface{})
	SetDecodeFunc(d server.DecodeFunc)
	IsaacIn() *isaac.ISAAC

	SendSkill(id, level, experience int)
	SendMessage(message string)

	UpdateInteractionQueue()
	SyncInventories()
	SyncEntityList()
	SendPlayerSync()
	SendGroundItemSync()
	UpdateVisibleEntities()
	ProcessChatQueue()
	ClearFlags()
}

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
	Log              log.Log

	// Some protocol specific data
	ProtoData interface{}
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
