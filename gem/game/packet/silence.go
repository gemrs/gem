package packet

import (
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol"
)

func init() {
	// Silence some packets we dont care about
	registerHandler((*protocol.InboundPing)(nil), nullHandler)
	registerHandler((*protocol.InboundMouseMovement)(nil), nullHandler)
	registerHandler((*protocol.InboundMouseClick)(nil), nullHandler)
	registerHandler((*protocol.InboundWindowFocus)(nil), nullHandler)
	registerHandler((*protocol.InboundKeyPress)(nil), nullHandler)
	registerHandler((*protocol.InboundCameraMovement)(nil), nullHandler)
}

func nullHandler(player protocol.Player, message server.Message) {
	//	player.Log().Debugf("Pinged!")
}
