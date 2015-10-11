package game

import (
	"gem/encoding"
	"gem/game/player"
	"gem/log"

	"github.com/qur/gopy/lib"
)

//go:generate gopygen -type GameClient -excfield "^[a-z].*" $GOFILE

// decodeFunc is used for parsing the read stream and dealing with the incoming data.
// If io.EOF is returned, it is assumed that we didn't have enough data, and
// the underlying buffer's read pointer is not altered.
type decodeFunc func(*GameClient) error

// GameClient is a client which serves players
type GameClient struct {
	py.BaseObject

	*Connection
	service *GameService
	decode  decodeFunc

	Log     *log.Module
	Session *player.Session
	Profile *player.Profile
}

// NewGameClient constructs a new GameClient
func NewGameClient(conn *Connection, svc *GameService) *GameClient {
	session, err := player.Session{}.Alloc()
	if err != nil {
		panic(err)
	}

	session.SetTarget(conn)

	client, err := GameClient{
		Connection: conn,
		service:    svc,
		decode:     svc.handshake,
		Session:    session,
		Log:        conn.Log,
	}.Alloc()
	if err != nil {
		panic(err)
	}

	// gopygen doesn't populate embedded fields
	client.Connection = conn

	return client
}

// Conn returns the underlying Connection
func (client *GameClient) Conn() *Connection {
	return client.Connection
}

// Decode processes incoming packets and adds them to the read queue
func (client *GameClient) Decode() error {
	return client.decode(client)
}

// Encode writes encoding.Encodables to the client's buffer using the session's outbound rand generator
func (client *GameClient) Encode(codable encoding.Encodable) error {
	return codable.Encode(client.Conn().writeBuffer, &client.Session.RandOut)
}
