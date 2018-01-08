package player

import "github.com/gemrs/gem/gem/game/entity"

type ChatMessage struct {
	Effects       uint8
	Colour        uint8
	Message       string
	PackedMessage []byte
}

func (player *Player) AppendChatMessage(m *ChatMessage) {
	player.chatQueue = append(player.chatQueue, m)
	player.SetFlags(player.Flags() | entity.MobFlagChatUpdate)
}

func (player *Player) ChatMessageQueue() []*ChatMessage {
	return player.chatQueue
}

func (player *Player) ProcessChatQueue() {
	if len(player.chatQueue) > 0 {
		player.chatQueue = player.chatQueue[1:]
	}
}
