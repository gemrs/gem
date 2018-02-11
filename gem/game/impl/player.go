package impl

import (
	"math/rand"

	"github.com/gemrs/gem/fork/github.com/gtank/isaac"
	"github.com/gemrs/gem/gem/auth"
	"github.com/gemrs/gem/gem/core/log"
	"github.com/gemrs/gem/gem/protocol"

	"github.com/gemrs/gem/gem/game/entity"
	entityimpl "github.com/gemrs/gem/gem/game/entity"
	game_event "github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/game/world"
)

// GameClient is a client which serves players
//glua:bind
type Player struct {
	index            int
	sector           protocol.Sector
	auth             auth.Provider
	world            protocol.World
	loadedRegion     *position.Region
	previousPosition *position.Absolute
	protoData        interface{}

	visibleEntities *entity.Collection
	chatQueue       []protocol.InboundChatMessage

	*server.Connection
	*entityimpl.GenericMob
	decode server.DecodeFunc

	randIn          isaac.ISAAC
	randOut         isaac.ISAAC
	serverRandKey   []uint32
	secureBlockSize int

	profile      protocol.Profile
	clientConfig *ClientConfig
	currentFrame protocol.FrameType
}

// NewGameClient constructs a new GameClient
//glua:bind constructor Player
func NewPlayer(index int, conn *server.Connection, worldInst *world.Instance, auth auth.Provider) *Player {
	player := &Player{}
	player.index = index
	player.Connection = conn
	player.world = worldInst
	player.auth = auth
	player.serverRandKey = []uint32{
		uint32(rand.Int31()), uint32(rand.Int31()),
	}

	nilPosition := position.NewAbsolute(0, 0, 0)
	player.sector = worldInst.Sector(nilPosition.Sector())
	player.loadedRegion = nilPosition.RegionOf()

	wpq := entityimpl.NewSimpleWaypointQueue()
	player.GenericMob = entityimpl.NewGenericMob(wpq)

	player.visibleEntities = entity.NewCollection()
	player.clientConfig = NewClientConfig(player)
	return player
}

//glua:bind
func (player *Player) Disconnect() {
	player.auth.SaveProfile(player.Profile())
	worldSector := player.world.Sector(player.Position().Sector())
	worldSector.Remove(player)
	player.world.SetPlayerSlot(player.Index(), nil)
	game_event.PlayerLogout.NotifyObservers(player)

	player.Conn().Disconnect()
}

func (player *Player) CurrentFrame() protocol.FrameType {
	return player.currentFrame
}

func (player *Player) SetCurrentFrame(frame protocol.FrameType) {
	player.currentFrame = frame
}

func (player *Player) WorldInstance() protocol.World {
	return player.world
}

//glua:bind
func (player *Player) Index() int {
	return player.index
}

// EntityType identifies what kind of entity this entity is
func (player *Player) EntityType() entity.EntityType {
	return entity.PlayerType
}

func (player *Player) ProtoData() interface{} {
	return player.protoData
}

func (player *Player) SetProtoData(d interface{}) {
	player.protoData = d
}

func (player *Player) SetDecodeFunc(d server.DecodeFunc) {
	player.decode = d
}

//glua:bind
func (player *Player) Logger() *log.Module {
	return &log.Module{player.Conn().Log()}
}

func (player *Player) SecureBlockSize() int {
	return player.secureBlockSize
}

func (player *Player) SetSecureBlockSize(size int) {
	player.secureBlockSize = size
}

func (player *Player) LoadedRegion() *position.Region {
	return player.loadedRegion
}

func (player *Player) VisibleEntities() *entity.Collection {
	return player.visibleEntities
}

// Profile returns the player's profile
//glua:bind
func (player *Player) Profile() protocol.Profile {
	return player.profile
}

// SetProfile sets the player's profile
func (player *Player) SetProfile(profile protocol.Profile) {
	player.profile = profile
}

//glua:bind
func (player *Player) ClientConfig() *ClientConfig {
	return player.clientConfig
}

// FinishInit is called once the player has finished the low level login sequence
func (player *Player) FinishInit() {
	player.Send(protocol.OutboundPlayerInit{
		Membership: 1,
		Index:      player.Index(),
	})

	profile := player.profile
	player.SetPosition(profile.Position())

	player.Send(protocol.OutboundRegionUpdate{
		ProtoData: player.protoData,
		Player:    player,
	})

	player.Send(protocol.OutboundInitInterface{
		ProtoData: player.protoData,
		Inventory: player.Profile().Inventory(),
	})

	// This triggers sync of inventory, skills etc, so it needs to happen
	// after the first region update (init gpi)
	profile.SetPlayer(player)

	player.Send(protocol.OutboundUpdateAllInventoryItems{
		Container: player.Profile().Inventory(),
	})

	game_event.PlayerLoadProfile.NotifyObservers(player, player.Profile())
}
