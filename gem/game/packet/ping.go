package packet

import (
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol"
)

func init() {
	registerHandler((*protocol.InboundPing)(nil), ping)
}

func ping(player protocol.Player, message server.Message) {
	//	player.Log().Debugf("Pinged!")
}
