package packet

import (
	"gem/encoding"
	"gem/game/server"
	"gem/protocol"
)

func init() {
	registerHandler((*protocol.InboundPing)(nil), ping)
}

func ping(player server.Player, packet encoding.Decodable) {
	player.Log().Debugf("Pinged!")
}
