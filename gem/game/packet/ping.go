package packet

import (
	"gem/encoding"
	"gem/game/player"
	game_protocol "gem/protocol/game"
)

func init() {
	registerHandler((*game_protocol.InboundPing)(nil), ping)
}

func ping(player player.Player, packet encoding.Decodable) {
	player.Log().Debugf("Pinged!")
}
