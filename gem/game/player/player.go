package player

import (
	"github.com/sinusoids/gem/gem/encoding"
	entityimpl "github.com/sinusoids/gem/gem/game/entity"
	game_event "github.com/sinusoids/gem/gem/game/event"
	"github.com/sinusoids/gem/gem/game/interface/entity"
	"github.com/sinusoids/gem/gem/game/interface/player"
	"github.com/sinusoids/gem/gem/game/position"
	"github.com/sinusoids/gem/gem/game/server"
	game_protocol "github.com/sinusoids/gem/gem/protocol/game"

	"github.com/qur/gopy/lib"
)

// GameClient is a client which serves players
type Player struct {
	py.BaseObject

	*server.Connection
	*entityimpl.GenericMob
	decode player.DecodeFunc

	session *Session
	profile *Profile
}

// NewGameClient constructs a new GameClient
func (client *Player) Init(conn *server.Connection) {
	session := NewSession()

	client.Connection = conn
	client.session = session

	wpq := entityimpl.NewSimpleWaypointQueue()

	client.GenericMob = entityimpl.NewGenericMob(wpq)
}

func (client *Player) SetDecodeFunc(d player.DecodeFunc) {
	client.decode = d
}

// Conn returns the underlying Connection
func (client *Player) Conn() *server.Connection {
	return client.Connection
}

// Encode writes encoding.Encodables to the client's buffer using the session's outbound rand generator
func (client *Player) Encode(codable encoding.Encodable) error {
	session := client.Session().(*Session)
	return codable.Encode(client.Conn().WriteBuffer, session.ISAACOut())
}

// Decode processes incoming packets and adds them to the read queue
func (client *Player) Decode() error {
	return client.decode(client)
}

// Session returns the player's session
func (client *Player) Session() player.Session {
	return client.session
}

// Profile returns the player's profile
func (client *Player) Profile() player.Profile {
	return client.profile
}

// SetProfile sets the player's profile
func (client *Player) SetProfile(profile player.Profile) {
	client.profile = profile.(*Profile)
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
	oldRegion := client.Region()
	client.GenericMob.SetPosition(pos)

	dx, dy, dz := client.GenericMob.Region().SectorDelta(oldRegion)

	eventArgs := map[string]interface{}{
		"entity":   client,
		"position": pos,
	}

	if dx >= 1 || dy >= 1 || dz >= 1 {
		game_event.EntitySectorChange.NotifyObservers(eventArgs)
	}

	if dx >= 5 || dy >= 5 || dz >= 1 {
		game_event.EntityRegionChange.NotifyObservers(eventArgs)
	}
}

// EntityType identifies what kind of entity this entity is
func (client *Player) EntityType() entity.EntityType {
	return entity.PlayerType
}
