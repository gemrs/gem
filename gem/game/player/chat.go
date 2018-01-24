package player

import (
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/protocol"
)

func (player *Player) AppendChatMessage(m protocol.InboundChatMessage) {
	player.chatQueue = append(player.chatQueue, m)
	player.SetFlags(player.Flags() | entity.MobFlagChatUpdate)
}

func (player *Player) ChatMessageQueue() []protocol.InboundChatMessage {
	return player.chatQueue
}

func (player *Player) ProcessChatQueue() {
	if len(player.chatQueue) > 0 {
		player.chatQueue = player.chatQueue[1:]
	}
}
