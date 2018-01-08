package player

import (
	"math/rand"

	"github.com/gemrs/gem/fork/github.com/gtank/isaac"

	"github.com/gemrs/gem/gem/encoding"
	"github.com/gemrs/gem/gem/game/entity"
	entityimpl "github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/game/world"
	"github.com/gemrs/gem/gem/protocol/game_protocol"
)

// GameClient is a client which serves players
//glua:bind
type Player struct {
	index        int
	sector       *world.Sector
	world        *world.Instance
	loadedRegion *position.Region

	visibleEntities *entity.Collection
	chatQueue       []*ChatMessage

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
	player := &Player{}
	player.Connection = conn
	player.world = worldInst
	player.serverRandKey = []uint32{
		uint32(rand.Int31()), uint32(rand.Int31()),
	}

	nilPosition := position.NewAbsolute(0, 0, 0)
	player.sector = worldInst.Sector(nilPosition.Sector())
	player.loadedRegion = nilPosition.RegionOf()

	wpq := entityimpl.NewSimpleWaypointQueue()
	player.GenericMob = entityimpl.NewGenericMob(wpq)

	player.visibleEntities = entity.NewCollection()
	player.animations = NewAnimations()
	player.index = entity.NextIndex()
	player.clientConfig = NewClientConfig(player)
	return player
}

//glua:bind
func (player *Player) Index() int {
	return player.index
}

// EntityType identifies what kind of entity this entity is
func (player *Player) EntityType() entity.EntityType {
	return entity.PlayerType
}

func (player *Player) SetDecodeFunc(d DecodeFunc) {
	player.decode = d
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

func (player *Player) Animations() *Animations {
	return player.animations
}

// Profile returns the player's profile
//glua:bind
func (player *Player) Profile() *Profile {
	return player.profile
}

// SetProfile sets the player's profile
func (player *Player) SetProfile(profile *Profile) {
	player.profile = profile
}

// Appearance returns the player's appearance
func (player *Player) Appearance() *Appearance {
	profile := player.Profile()
	return profile.Appearance()
}

//glua:bind
func (player *Player) ClientConfig() *ClientConfig {
	return player.clientConfig
}

// FinishInit is called once the player has finished the low level login sequence
func (player *Player) FinishInit() {
	player.Conn().Write <- &game_protocol.OutboundPlayerInit{
		Membership: encoding.Uint8(1),
		Index:      encoding.Uint16(player.Index()),
	}
}
