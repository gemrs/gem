package packet

import (
	"github.com/gemrs/gem/gem/encoding"
	"github.com/gemrs/gem/gem/game/player"
	"github.com/gemrs/gem/gem/protocol/game_protocol"
)

func init() {
	registerHandler((*game_protocol.InboundPing)(nil), ping)
}

func ping(player *player.Player, packet encoding.Decodable) {
	//	player.Log().Debugf("Pinged!")
}
