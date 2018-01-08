package player

import (
	"github.com/gemrs/gem/gem/encoding"
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/protocol/game_protocol"
)

// SendPlayerSync sends the player update block
func (player *Player) SendPlayerSync() {
	player.Conn().Write <- NewPlayerUpdateBlock(player)
}

// SendMessage puts a message to the player's chat window
//glua:bind
func (player *Player) SendMessage(message string) {
	player.Conn().Write <- &game_protocol.OutboundChatMessage{
		Message: encoding.JString(message),
	}
}

func (player *Player) sendTabInterface(tab, id int) {
	player.Conn().Write <- &game_protocol.OutboundTabInterface{
		Tab:         encoding.Uint8(tab),
		InterfaceID: encoding.Uint16(id),
	}
}

// Ask the player to log out
//glua:bind
func (player *Player) SendForceLogout() {
	player.Conn().Write <- &game_protocol.OutboundLogout{}
}

func (player *Player) ClearFlags() {
	player.GenericMob.ClearFlags()
	// Don't clear the chat flag if there are still messages queued
	if len(player.chatQueue) > 0 {
		player.SetFlags(player.Flags() | entity.MobFlagChatUpdate)
	}
}
