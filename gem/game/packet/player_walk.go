package packet

import (
	"fmt"

	"github.com/gemrs/gem/gem/encoding"
	"github.com/gemrs/gem/gem/game/player"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/protocol/game_protocol"
)

func init() {
	registerHandler((*game_protocol.InboundPlayerWalk)(nil), player_walk)
	registerHandler((*game_protocol.InboundPlayerWalkMap)(nil), player_walk)
}

func player_walk(player *player.Player, packet encoding.Decodable) {
	var walkPacket *game_protocol.InboundPlayerWalkBlock
	switch p := packet.(type) {
	case *game_protocol.InboundPlayerWalk:
		walkPacket = (*game_protocol.InboundPlayerWalkBlock)(p)
	case *game_protocol.InboundPlayerWalkMap:
		walkPacket = (*game_protocol.InboundPlayerWalkBlock)(p)
	default:
		panic(fmt.Sprintf("got invalid walk packet: %T", p))
	}

	height := player.Position().Z()
	origin := position.NewAbsolute(int(walkPacket.OriginX), int(walkPacket.OriginY), height)

	waypoints := make([]*position.Absolute, len(walkPacket.Waypoints))

	wpq := player.WaypointQueue()
	wpq.Clear()
	wpq.Push(origin)
	for i, wp := range walkPacket.Waypoints {
		x1, y1 := int(wp.X), int(wp.Y)
		x2, y2 := int(origin.X()), int(origin.Y())
		waypoints[i] = position.NewAbsolute(int(x1+x2), int(y1+y2), height)

		wpq.Push(waypoints[i])
	}
}
