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

// Send writes encoding.Encodables to the player's buffer using the outbound rand generator
func (player *Player) Send(codable encoding.Encodable) error {
	var buf encoding.Encoded
	err := encoding.TryEncode(codable, &buf, &player.randOut)
	if err != nil {
		return err
	}
	player.Write <- buf
	return nil
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
