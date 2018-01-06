package player

import (
	"math/rand"

	"github.com/gtank/isaac"

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
type Player struct {
	index        int
	sector       *world.Sector
	world        *world.Instance
	loadedRegion *position.Region

	*server.Connection
	*entityimpl.GenericMob
	decode player.DecodeFunc

	randIn          isaac.ISAAC
	randOut         isaac.ISAAC
	serverRandKey   []uint32
	secureBlockSize int

	animations *Animations
	profile    *Profile
}

// NewGameClient constructs a new GameClient
func (client *Player) Init(conn *server.Connection, worldInst *world.Instance) {
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

	client.animations = NewAnimations()
	client.index = entity.NextIndex()
}

func (client *Player) Index() int {
	return client.index
}

func (client *Player) SetDecodeFunc(d player.DecodeFunc) {
	client.decode = d
}

// Conn returns the underlying Connection
func (client *Player) Conn() *server.Connection {
	return client.Connection
}

// Encode writes encoding.Encodables to the client's buffer using the outbound rand generator
func (client *Player) Encode(codable encoding.Encodable) error {
	return codable.Encode(client.Conn().WriteBuffer, client.ISAACOut())
}

// Decode processes incoming packets and adds them to the read queue
func (client *Player) Decode() error {
	return client.decode(client)
}

func (client *Player) ServerISAACSeed() []uint32 {
	return client.serverRandKey
}

func (client *Player) ISAACIn() *isaac.ISAAC {
	return &client.randIn
}

func (client *Player) ISAACOut() *isaac.ISAAC {
	return &client.randOut
}

func (client *Player) InitISAAC(inSeed, outSeed []uint32) {
	client.randIn.SeedArray(inSeed)
	client.randOut.SeedArray(outSeed)
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

func (client *Player) Animations() player.Animations {
	return client.animations
}

// Profile returns the player's profile
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

// SetAppearance modifies the player's appearance
func (client *Player) SetAppearance(a player.Appearance) {
	profile := client.Profile().(*Profile)
	profile.SetAppearance(a)
	client.AppearanceUpdated()
}

// AppearanceUpdated signals that the player's appearance should be re-synchronized
func (client *Player) AppearanceUpdated() {
	eventArgs := map[string]interface{}{
		"entity": client,
	}
	game_event.PlayerAppearanceUpdate.NotifyObservers(eventArgs)
}

// SendMessage puts a message to the player's chat window
func (client *Player) SendMessage(message string) {
	client.Conn().Write <- &game_protocol.OutboundChatMessage{
		Message: encoding.JString(message),
	}
}

func (client *Player) SetNextStep(next *position.Absolute) {
	client.SetPosition(next)
	client.GenericMob.SetNextStep(next)
}

// SetPosition warps the mob to a given location
func (client *Player) SetPosition(pos *position.Absolute) {
	oldSector := client.sector.Position()
	client.GenericMob.SetPosition(pos)

	dx, dy, dz := oldSector.Delta(pos.Sector())

	eventArgs := map[string]interface{}{
		"entity":   client,
		"position": pos,
	}

	if dx >= 1 || dy >= 1 || dz >= 1 {
		game_event.EntitySectorChange.NotifyObservers(eventArgs)
	}

	loadedRegion := client.LoadedRegion()
	dx, dy, dz = loadedRegion.SectorDelta(pos.RegionOf())

	if dx >= 5 || dy >= 5 || dz >= 1 {
		game_event.EntityRegionChange.NotifyObservers(eventArgs)
	}
}

// EntityType identifies what kind of entity this entity is
func (client *Player) EntityType() entity.EntityType {
	return entity.PlayerType
}
