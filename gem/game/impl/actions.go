package impl

import (
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/protocol"
)

// SendPlayerSync sends the player update block
func (player *Player) SendPlayerSync() {
	player.Send(buildPlayerUpdate(player))
}

// SendMessage puts a message to the player's chat window
//glua:bind
func (player *Player) SendMessage(message string) {
	player.Send(protocol.OutboundChatMessage{
		Message: message,
	})
}

// SendMessage puts a message to the player's chat window
//glua:bind
func (player *Player) SendSkill(id, level, experience int) {
	player.Send(protocol.OutboundSkill{
		Skill:      id,
		Level:      level,
		Experience: experience,
	})
}

func (player *Player) sendTabInterface(tab protocol.InterfaceTab, id int) {
	player.Send(protocol.OutboundTabInterface{
		ProtoData:   player.protoData,
		Tab:         tab,
		InterfaceID: id,
	})
}

// Ask the player to log out
//glua:bind
func (player *Player) SendForceLogout() {
	player.Send(protocol.OutboundLogout{})
}

func (player *Player) ClearFlags() {
	player.GenericMob.ClearFlags()
	// Don't clear the chat flag if there are still messages queued
	if len(player.chatQueue) > 0 {
		player.SetFlags(player.Flags() | entity.MobFlagChatUpdate)
	}
}
