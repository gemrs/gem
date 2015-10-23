package packet

import (
	"fmt"

	"github.com/sinusoids/gem/gem/encoding"
	"github.com/sinusoids/gem/gem/game/player"
	"github.com/sinusoids/gem/gem/game/position"
	game_protocol "github.com/sinusoids/gem/gem/protocol/game"
)

func init() {
	registerHandler((*game_protocol.InboundPlayerWalk)(nil), player_walk)
	registerHandler((*game_protocol.InboundPlayerWalkMap)(nil), player_walk)
}

func player_walk(player player.Player, packet encoding.Decodable) {
	var walkPacket *game_protocol.InboundPlayerWalkBlock
	switch p := packet.(type) {
	case *game_protocol.InboundPlayerWalk:
		walkPacket = (*game_protocol.InboundPlayerWalkBlock)(p)
	case *game_protocol.InboundPlayerWalkMap:
		walkPacket = (*game_protocol.InboundPlayerWalkBlock)(p)
	default:
		panic(fmt.Sprintf("got invalid walk packet: %T", p))
	}

	height := player.Position().Z
	origin, err := position.NewAbsolute(int(walkPacket.OriginX), int(walkPacket.OriginY), height)
	if err != nil {
		panic(err)
	}

	waypoints := make([]*position.Absolute, len(walkPacket.Waypoints))
	player.Log().Debugf("Origin %v", origin)
	for i, wp := range walkPacket.Waypoints {
		waypoints[i], err = position.NewAbsolute(int(wp.X)+origin.X, int(wp.Y)+origin.Y, height)
		if err != nil {
			panic(err)
		}

		player.Log().Debugf("Waypoint %v %v", i, waypoints[i])
	}
}
