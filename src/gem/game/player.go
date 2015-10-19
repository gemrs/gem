package game

import (
	"gem"
	"gem/encoding"
	"gem/event"
	"gem/game/player"
	"gem/game/position"
	"gem/game/server"

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
	flags   int
}

// NewGameClient constructs a new GameClient
func (client *Player) Init(conn *server.Connection, svc *GameService) error {
	session, err := player.NewSession()
	if err != nil {
		return err
	}

	session.SetTarget(conn)

	client.Connection = conn
	client.service = svc
	client.decode = svc.handshake
	client.session = session

	client.region, err = position.NewRegion(nil)
	if err != nil {
		return err
	}

	PlayerRegionChangeEvent.Register(event.NewListener(client.RegionUpdate))

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
		PlayerSectorChangeEvent.NotifyObservers(pos)
	}

	if dx >= 5 || dy >= 5 || dz >= 1 {
		PlayerRegionChangeEvent.NotifyObservers(pos)
	}

	client.Log().Debugf("Warping to %v", pos)
}

// Encode writes encoding.Encodables to the client's buffer using the session's outbound rand generator
func (client *Player) Encode(codable encoding.Encodable) error {
	return codable.Encode(client.Conn().WriteBuffer, &client.Session().RandOut)
}
