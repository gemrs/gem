package game

import (
	"gem"
	"gem/encoding"
	"gem/event"
	"gem/game/entity"
	"gem/game/player"
	"gem/game/position"
	"gem/game/server"
	"gem/protocol"

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
	service *GameService
	decode  decodeFunc

	session *player.Session
	profile *player.Profile
	region  *position.Region
	flags   entity.Flags
}

// NewGameClient constructs a new GameClient
func (client *Player) Init(conn *server.Connection, svc *GameService) error {
	session, err := player.NewSession()
	if err != nil {
		return err
	}

	client.Connection = conn
	client.service = svc
	client.decode = svc.handshake
	client.session = session

	client.region, err = position.NewRegion(nil)
	if err != nil {
		return err
	}

	PlayerRegionChangeEvent.Register(event.NewListener(client.RegionUpdate))
	PlayerAppearanceUpdateEvent.Register(event.NewListener(client.AppearanceUpdate))

	return nil
}

func finishLogin(_ *event.Event, args ...interface{}) {
	client := args[0].(*Player)
	client.PlayerInit()
	gem.TickEvent.Register(event.NewListener(client.PlayerUpdate))
	gem.PostTickEvent.Register(event.NewListener(client.ClearUpdateFlags))
}

func (client *Player) Session() *player.Session {
	return client.session
}

func (client *Player) Profile() *player.Profile {
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
	return client.Profile().Pos
}

// SetPosition warps the player to a given location
func (client *Player) SetPosition(pos *position.Absolute) {
	client.Profile().Pos = pos
	oldRegion := client.region
	client.region = pos.RegionOf()
	dx, dy, dz := client.region.SectorDelta(oldRegion)

	if dx >= 1 || dy >= 1 || dz >= 1 {
		PlayerSectorChangeEvent.NotifyObservers(client, pos)
	}

	if dx >= 5 || dy >= 5 || dz >= 1 {
		PlayerRegionChangeEvent.NotifyObservers(client, pos)
	}

	client.Log().Debugf("Warping to %v", pos)
}

// SetAppearance modifies the player's appearance
func (client *Player) SetAppearance(a *player.Appearance) {
	client.Profile().Appearance = a
	client.AppearanceUpdated()
}

// AppearanceUpdated signals that the player's appearance should be re-synchronized
func (client *Player) AppearanceUpdated() {
	PlayerAppearanceUpdateEvent.NotifyObservers(client)
}

// Flags returns the mob update flags for this player
func (client *Player) Flags() entity.Flags {
	return client.flags
}

// Region returns the player's current surrounding region
func (client *Player) Region() *position.Region {
	return client.region
}

// Encode writes encoding.Encodables to the client's buffer using the session's outbound rand generator
func (client *Player) Encode(codable encoding.Encodable) error {
	return codable.Encode(client.Conn().WriteBuffer, &client.Session().RandOut)
}

// WalkDirection returns the current and previous walking directions
func (client *Player) WalkDirection() (current, last int) {
	return 0, 0
}

// SendMessage puts a message to the player's chat window
func (client *Player) SendMessage(message string) {
	client.Conn().Write <- &protocol.OutboundChatMessage{
		Message: encoding.JString(message),
	}
}
