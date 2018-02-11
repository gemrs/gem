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
	Equipment() Equipment
	Appearance() Appearance
	Skills() Skills
	Position() *position.Absolute
	SetPosition(*position.Absolute)
	SetPlayer(Player)
	Save() string
}

type Player interface {
	Index() int
	Flags() entity.Flags
	Profile() Profile
	WorldInstance() World
	Position() *position.Absolute
	PreviousPosition() *position.Absolute
	LoadedRegion() *position.Region
	SetWalkDestination(pos *position.Absolute) bool
	SetRunning(bool)

	AppendChatMessage(m InboundChatMessage)
	ChatMessageQueue() []InboundChatMessage
	InteractionQueue() *entity.InteractionQueue
	WaypointQueue() entity.WaypointQueue
	CurrentFrame() FrameType
	SetCurrentFrame(FrameType)
	SetAppearanceChanged()

	Conn() *server.Connection
	Decode() error
	Send(encoding.Encodable) error
	Disconnect()
	ProtoData() interface{}
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

	Log() log.Log
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
	SkullIcon() int
	Model(b BodyPart) int
	SetModel(b BodyPart, model int)
	Color(b BodyPart) int
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
	SkillHunter
	SkillConstruction

	SkillMax
)

type Skills interface {
	CombatLevel() int
	//	Skill(id SkillId) *Skill
}

type Equipment interface {
	Equip(slot int, item *item.Stack) (oldEquipment *item.Stack)
	Container() *item.Container
	Slot(i int) *item.Stack
	Unequip(slot int) *item.Stack
}
