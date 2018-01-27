package player

import (
	"math/rand"

	"github.com/gemrs/gem/fork/github.com/gtank/isaac"
	"github.com/gemrs/gem/gem/core/log"
	"github.com/gemrs/gem/gem/protocol"

	"github.com/gemrs/gem/gem/game/entity"
	entityimpl "github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/game/world"
)

// GameClient is a client which serves players
//glua:bind
type Player struct {
	index        int
	sector       *world.Sector
	world        *world.Instance
	loadedRegion *position.Region
	protoData    interface{}

	visibleEntities *entity.Collection
	chatQueue       []protocol.InboundChatMessage

	*server.Connection
	*entityimpl.GenericMob
	decode server.DecodeFunc

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

func (player *Player) WorldInstance() *world.Instance {
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
	player.Conn().Write <- protocol.OutboundPlayerInit{
		Membership: 1,
		Index:      player.Index(),
	}
	/*
		player.Conn().Write <- protocol.OutboundResetCamera{}
		player.Conn().Write <- protocol.OutboundDnsLookup{
			Host: "127.0.0.1",
		}
	*/

	player.Conn().Write <- protocol.OutboundInitInterface{
		ProtoData: player.protoData,
	}
	/*
		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     chatBox,
			InterfaceID: 162,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     expDisplay,
			InterfaceID: 163,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     dataOrbs,
			InterfaceID: 160,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     privateChat,
			InterfaceID: 122,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     29,
			InterfaceID: 378,
			Clickable:   false,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     28,
			InterfaceID: 50,
			Clickable:   false,
		}

		player.Conn().Write <- protocol.OutboundScriptEvent{
			ScriptID: 233,
			Params: []interface{}{
				3276804, 7085, 0, 0, 434, 1912, 0, 400, -1,
			},
		}

		player.Conn().Write <- protocol.OutboundScriptEvent{
			ScriptID: 233,
			Params: []interface{}{
				3276805, 32817, 0, 100, 93, 179, 0, 800, 820,
			},
		}

		player.Conn().Write <- protocol.OutboundScriptEvent{
			ScriptID: 1080,
			Params:   []interface{}{},
		}

		attackStyleTab := 8
		skillsTab := 9
		journeyTab := 10
		itemBagTab := 11
		equipmentTab := 12
		prayerTab := 13
		spellBookTab := 14
		clanChatTab := 15
		friendListTab := 16
		ignoreListTab := 17
		logoutTab := 18
		gameSettingsTab := 19
		emotesTab := 20
		musicListTab := 21

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     attackStyleTab,
			InterfaceID: 593,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     skillsTab,
			InterfaceID: 320,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     journeyTab,
			InterfaceID: 399,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     itemBagTab,
			InterfaceID: 149,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     equipmentTab,
			InterfaceID: 387,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     prayerTab,
			InterfaceID: 541,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     spellBookTab,
			InterfaceID: 218,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     clanChatTab,
			InterfaceID: 7,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     friendListTab,
			InterfaceID: 429,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     ignoreListTab,
			InterfaceID: 432,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     logoutTab,
			InterfaceID: 182,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     gameSettingsTab,
			InterfaceID: 261,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     emotesTab,
			InterfaceID: 216,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundSetInterface{
			RootID:      rootIface,
			ChildID:     musicListTab,
			InterfaceID: 239,
			Clickable:   true,
		}

		player.Conn().Write <- protocol.OutboundScriptEvent{
			ScriptID: 2014,
			Params:   []interface{}{0, 0, 0, 0, 0, 0},
		}

		player.Conn().Write <- protocol.OutboundScriptEvent{
			ScriptID: 2015,
			Params:   []interface{}{0},
		}
	*/
}
