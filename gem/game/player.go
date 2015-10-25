package game

import (
	"github.com/sinusoids/gem/gem/encoding"
	engine_event "github.com/sinusoids/gem/gem/engine/event"
	"github.com/sinusoids/gem/gem/event"
	game_event "github.com/sinusoids/gem/gem/game/event"
	"github.com/sinusoids/gem/gem/game/interface/entity"
	"github.com/sinusoids/gem/gem/game/interface/player"
	"github.com/sinusoids/gem/gem/game/position"
	"github.com/sinusoids/gem/gem/game/server"
	game_protocol "github.com/sinusoids/gem/gem/protocol/game"

	"github.com/qur/gopy/lib"
)

//go:generate gopygen -type Player -excfield "^[a-z].*" $GOFILE

// decodeFunc is used for parsing the read stream and dealing with the incoming data.
// If io.EOF is returned, it is assumed that we didn't have enough data, and
// the underlying buffer's read pointer is not altered.
type decodeFunc func(*Player) error

// GameClient is a client which serves players
type Player struct {
	py.BaseObject

	*server.Connection
	decode decodeFunc

	session *Session
	profile *Profile
}

// NewGameClient constructs a new GameClient
func (client *Player) Init(conn *server.Connection) error {
	session, err := NewSession()
	if err != nil {
		return err
	}

	client.Connection = conn
	client.session = session

	client.session.region, err = position.NewRegion(nil)
	if err != nil {
		return err
	}

	game_event.PlayerRegionChange.Register(event.NewListener(client.RegionUpdate))
	game_event.PlayerAppearanceUpdate.Register(event.NewListener(client.AppearanceUpdate))

	return nil
}

// finishLogin calls PlayerInit and registers tick event callbacks for various things
func finishLogin(_ *event.Event, args ...interface{}) {
	client := args[0].(*Player)
	client.PlayerInit()
	engine_event.Tick.Register(event.NewListener(client.PlayerUpdate))
	engine_event.PostTick.Register(event.NewListener(client.ClearUpdateFlags))
}

func (client *Player) SetDecodeFunc(d decodeFunc) {
	client.decode = d
}

// Session returns the player's session
func (client *Player) Session() player.Session {
	return client.session
}

// Profile returns the player's profile
func (client *Player) Profile() player.Profile {
	return client.profile
}

// Conn returns the underlying Connection
func (client *Player) Conn() *server.Connection {
	return client.Connection
}

// Decode processes incoming packets and adds them to the read queue
func (client *Player) Decode() error {
	return client.decode(client)
}

// Position returns the absolute position of the player
func (client *Player) Position() *position.Absolute {
	return client.Profile().Position()
}

// WalkDirection returns the player's current and (in the case of running) last walking direction
func (client *Player) WalkDirection() (current int, last int) {
	return client.Session().WalkDirection()
}

// Flags returns the mob update flags for this player
func (client *Player) Flags() entity.Flags {
	return client.Session().Flags()
}

// Region returns the player's current surrounding region
func (client *Player) Region() *position.Region {
	return client.Session().Region()
}

// SetPosition warps the player to a given location
func (client *Player) SetPosition(pos *position.Absolute) {
	profile := client.Profile().(*Profile)
	session := client.Session().(*Session)

	profile.SetPosition(pos)
	oldRegion := session.Region()
	session.SetRegion(pos.RegionOf())
	dx, dy, dz := session.Region().SectorDelta(oldRegion)

	if dx >= 1 || dy >= 1 || dz >= 1 {
		game_event.PlayerSectorChange.NotifyObservers(client, pos)
	}

	if dx >= 5 || dy >= 5 || dz >= 1 {
		game_event.PlayerRegionChange.NotifyObservers(client, pos)
	}

	client.Log().Debugf("Warping to %v", pos)
}

// SetAppearance modifies the player's appearance
func (client *Player) SetAppearance(a player.Appearance) {
	profile := client.Profile().(*Profile)
	profile.SetAppearance(a)
	client.AppearanceUpdated()
}

// AppearanceUpdated signals that the player's appearance should be re-synchronized
func (client *Player) AppearanceUpdated() {
	game_event.PlayerAppearanceUpdate.NotifyObservers(client)
}

// Encode writes encoding.Encodables to the client's buffer using the session's outbound rand generator
func (client *Player) Encode(codable encoding.Encodable) error {
	session := client.Session().(*Session)
	return codable.Encode(client.Conn().WriteBuffer, &session.RandOut)
}

// SendMessage puts a message to the player's chat window
func (client *Player) SendMessage(message string) {
	client.Conn().Write <- &game_protocol.OutboundChatMessage{
		Message: encoding.JString(message),
	}
}
