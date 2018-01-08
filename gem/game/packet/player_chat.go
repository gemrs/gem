package packet

import (
	"github.com/gemrs/gem/gem/encoding"
	"github.com/gemrs/gem/gem/game/player"
	"github.com/gemrs/gem/gem/protocol/game_protocol"
)

func init() {
	registerHandler((*game_protocol.InboundChatMessage)(nil), player_chat)
}

func player_chat(p *player.Player, packet encoding.Decodable) {
	chatPacket := packet.(*game_protocol.InboundChatMessage)
	data := ([]byte)(chatPacket.EncodedMessage)
	size := len(data)
	decoded := make([]byte, size)
	for i, _ := range data {
		decoded[i] = byte(data[size-i-1] - 128)
	}
	message := encoding.ChatTextSanitize(encoding.ChatTextUnpack(decoded))
	p.AppendChatMessage(&player.ChatMessage{
		Effects:       uint8(chatPacket.Effects),
		Colour:        uint8(chatPacket.Colour),
		Message:       message,
		PackedMessage: encoding.ChatTextPack(message),
	})
}
