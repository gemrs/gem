package packet

import (
	"gem/encoding"
	"gem/game/entity"
	"gem/protocol"
)

func init() {
	registerHandler((*protocol.InboundPing)(nil), ping)
}

func ping(player entity.Player, packet encoding.Decodable) {
	player.Log().Debugf("Pinged!")
}
