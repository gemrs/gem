package packet

import (
	"gem/encoding"
	"gem/game/player"
	"gem/protocol"
)

func init() {
	registerHandler((*protocol.InboundPing)(nil), ping)
}

func ping(player player.Player, packet encoding.Decodable) {
	player.Log().Debugf("Pinged!")
}
