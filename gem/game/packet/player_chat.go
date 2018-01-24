package packet

import (
	"github.com/gemrs/gem/gem/game/player"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol"
)

func init() {
	registerHandler((*protocol.InboundChatMessage)(nil), player_chat)
}

func player_chat(p *player.Player, message server.Message) {
	chatPacket := message.(*protocol.InboundChatMessage)
	p.AppendChatMessage(*chatPacket)
}
