package packet

import (
	"fmt"

	"gem/encoding"
	"gem/game/player"
	"gem/game/position"
	"gem/protocol"
)

func init() {
	registerHandler((*protocol.InboundPlayerWalk)(nil), player_walk)
	registerHandler((*protocol.InboundPlayerWalkMap)(nil), player_walk)
}

func player_walk(player player.Player, packet encoding.Decodable) {
	var walkPacket *protocol.InboundPlayerWalkBlock
	switch p := packet.(type) {
	case *protocol.InboundPlayerWalk:
		walkPacket = (*protocol.InboundPlayerWalkBlock)(p)
	case *protocol.InboundPlayerWalkMap:
		walkPacket = (*protocol.InboundPlayerWalkBlock)(p)
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
