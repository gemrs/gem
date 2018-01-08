package player

import (
	"github.com/gemrs/gem/fork/github.com/gtank/isaac"
	"github.com/gemrs/gem/gem/encoding"
	"github.com/gemrs/gem/gem/game/server"
)

// DecodeFunc is used for parsing the read stream and dealing with the incoming data.
// If io.EOF is returned, it is assumed that we didn't have enough data, and
// the underlying buffer's read pointer is not altered.
type DecodeFunc func(*Player) error

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

func (client *Player) InitISAAC(inSeed, outSeed []uint32) {
	client.randIn.SeedArray(inSeed)
	client.randOut.SeedArray(outSeed)
}

func (client *Player) ISAACIn() *isaac.ISAAC {
	return &client.randIn
}

func (client *Player) ISAACOut() *isaac.ISAAC {
	return &client.randOut
}
