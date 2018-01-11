package packet

import (
	"github.com/gemrs/gem/gem/encoding"
	"github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/player"
	"github.com/gemrs/gem/gem/protocol/game_protocol"
)

func init() {
	registerHandler((*game_protocol.InboundCommand)(nil), player_command)
}

func player_command(p *player.Player, packet encoding.Decodable) {
	commandPacket := packet.(*game_protocol.InboundCommand)
	command := commandPacket.Command
	game_event.PlayerCommand.NotifyObservers(p, string(command))
}
