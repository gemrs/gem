package player

import (
	"math/rand"

	"github.com/gemrs/gem/fork/github.com/gtank/isaac"

	"github.com/gemrs/gem/gem/encoding"
	entityimpl "github.com/gemrs/gem/gem/game/entity"
	game_event "github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/interface/entity"
	"github.com/gemrs/gem/gem/game/interface/player"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/game/world"
	game_protocol "github.com/gemrs/gem/gem/protocol/game"
)

// GameClient is a client which serves players
//glua:bind
type Player struct {
	index        int
	sector       *world.Sector
	world        *world.Instance
	loadedRegion *position.Region

	visibleEntities *entity.Collection
	chatQueue       []*player.ChatMessage

	*server.Connection
	*entityimpl.GenericMob
	decode DecodeFunc

	randIn          isaac.ISAAC
	randOut         isaac.ISAAC
	serverRandKey   []uint32
	secureBlockSize int

	animations   *Animations
	profile      *Profile
	clientConfig *ClientConfig
}

// NewGameClient constructs a new GameClient
//glua:bind constructor Player
func NewPlayer(conn *server.Connection, worldInst *world.Instance) *Player {
	client := &Player{}
	client.Connection = conn
	client.world = worldInst
	client.serverRandKey = []uint32{
		uint32(rand.Int31()), uint32(rand.Int31()),
	}

	nilPosition := position.NewAbsolute(0, 0, 0)
	client.sector = worldInst.Sector(nilPosition.Sector())
	client.loadedRegion = nilPosition.RegionOf()

	wpq := entityimpl.NewSimpleWaypointQueue()
	client.GenericMob = entityimpl.NewGenericMob(wpq)

	client.visibleEntities = entity.NewCollection()
	client.animations = NewAnimations()
	client.index = entity.NextIndex()
	client.clientConfig = NewClientConfig(client)
	return client
}

//glua:bind
func (client *Player) Index() int {
	return client.index
}

// EntityType identifies what kind of entity this entity is
func (client *Player) EntityType() entity.EntityType {
	return entity.PlayerType
}

func (client *Player) SetDecodeFunc(d DecodeFunc) {
	client.decode = d
}

func (client *Player) SecureBlockSize() int {
	return client.secureBlockSize
}

func (client *Player) SetSecureBlockSize(size int) {
	client.secureBlockSize = size
}

func (client *Player) LoadedRegion() *position.Region {
	return client.loadedRegion
}

func (client *Player) VisibleEntities() *entity.Collection {
	return client.visibleEntities
}

func (client *Player) Animations() player.Animations {
	return client.animations
}

// Profile returns the player's profile
//glua:bind
func (client *Player) Profile() player.Profile {
	return client.profile
}

// SetProfile sets the player's profile
func (client *Player) SetProfile(profile player.Profile) {
	client.profile = profile.(*Profile)
}

// Appearance returns the player's appearance
func (client *Player) Appearance() player.Appearance {
	profile := client.Profile().(*Profile)
	return profile.Appearance()
}

//glua:bind
func (client *Player) ClientConfig() player.ClientConfig {
	return client.clientConfig
}

// SendPlayerSync sends the player update block
func (client *Player) SendPlayerSync() {
	client.Conn().Write <- game_protocol.NewPlayerUpdateBlock(client)
}

// SendMessage puts a message to the player's chat window
//glua:bind
func (client *Player) SendMessage(message string) {
	client.Conn().Write <- &game_protocol.OutboundChatMessage{
		Message: encoding.JString(message),
	}
}

func (client *Player) sendTabInterface(tab, id int) {
	client.Conn().Write <- &game_protocol.OutboundTabInterface{
		Tab:         encoding.Uint8(tab),
		InterfaceID: encoding.Uint16(id),
	}
}

// Ask the client to log out
//glua:bind
func (client *Player) SendForceLogout() {
	client.Conn().Write <- &game_protocol.OutboundLogout{}
}

func (client *Player) AppendChatMessage(m *player.ChatMessage) {
	client.chatQueue = append(client.chatQueue, m)
	client.SetFlags(client.Flags() | entity.MobFlagChatUpdate)
}

func (client *Player) ChatMessageQueue() []*player.ChatMessage {
	return client.chatQueue
}

func (client *Player) ProcessChatQueue() {
	if len(client.chatQueue) > 0 {
		client.chatQueue = client.chatQueue[1:]
	}
}

func (client *Player) ClearFlags() {
	client.GenericMob.ClearFlags()
	// Don't clear the chat flag if there are still messages queued
	if len(client.chatQueue) > 0 {
		client.SetFlags(client.Flags() | entity.MobFlagChatUpdate)
	}
}

func (client *Player) LoadProfile() {
	profile := client.Profile().(*Profile)
	profile.setPlayer(client)
	client.SetPosition(profile.Position())

	game_event.PlayerLoadProfile.NotifyObservers(client, client.Profile().(*Profile))
}

// FinishInit is called once the player has finished the low level login sequence
func (client *Player) FinishInit() {
	client.Conn().Write <- &game_protocol.OutboundPlayerInit{
		Membership: encoding.Uint8(1),
		Index:      encoding.Uint16(client.Index()),
	}
}
