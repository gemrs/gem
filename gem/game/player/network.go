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
func (player *Player) Conn() *server.Connection {
	return player.Connection
}

// Encode writes encoding.Encodables to the player's buffer using the outbound rand generator
func (player *Player) Encode(codable encoding.Encodable) error {
	return codable.Encode(player.Conn().WriteBuffer, &player.randOut)
}

// Decode processes incoming packets and adds them to the read queue
func (player *Player) Decode() error {
	return player.decode(player)
}

func (player *Player) ServerISAACSeed() []uint32 {
	return player.serverRandKey
}

func (player *Player) InitISAAC(inSeed, outSeed []uint32) {
	player.randIn.SeedArray(inSeed)
	player.randOut.SeedArray(outSeed)
}

func (player *Player) ISAACIn() *isaac.ISAAC {
	return &player.randIn
}
