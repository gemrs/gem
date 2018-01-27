package impl

import (
	"github.com/gemrs/gem/fork/github.com/gtank/isaac"
	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/game/server"
)

// Conn returns the underlying Connection
func (player *Player) Conn() *server.Connection {
	return player.Connection
}

// Encode writes encoding.Encodables to the player's buffer using the outbound rand generator
func (player *Player) Encode(codable encoding.Encodable) error {
	return encoding.TryEncode(codable, player.Conn().WriteBuffer, &player.randOut)
}

// Decode processes incoming packets and adds them to the read queue
func (player *Player) Decode() error {
	return player.decode(player)
}

func (player *Player) ServerIsaacSeed() []uint32 {
	return player.serverRandKey
}

func (player *Player) InitIsaac(inSeed, outSeed []uint32) {
	player.randIn.SeedArray(inSeed)
	player.randOut.SeedArray(outSeed)
}

func (player *Player) IsaacIn() *isaac.ISAAC {
	return &player.randIn
}
