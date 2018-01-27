package packet

import (
	"github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol"
)

func init() {
	registerHandler((*protocol.InboundCommand)(nil), player_command)
}

func player_command(p protocol.Player, message server.Message) {
	commandPacket := message.(*protocol.InboundCommand)
	command := commandPacket.Command
	game_event.PlayerCommand.NotifyObservers(p, string(command))
}
