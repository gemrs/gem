package packet

import (
	"github.com/gemrs/gem/gem/game/player"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol"
)

func init() {
	registerHandler((*protocol.InboundPlayerWalk)(nil), player_walk)
}

func player_walk(player *player.Player, message server.Message) {
	walkPacket := message.(*protocol.InboundPlayerWalk)

	height := player.Position().Z()
	origin := position.NewAbsolute(walkPacket.OriginX, walkPacket.OriginY, height)

	waypoints := make([]*position.Absolute, len(walkPacket.Waypoints))

	wpq := player.WaypointQueue()
	wpq.Clear()
	wpq.Push(origin)
	for i, wp := range walkPacket.Waypoints {
		x1, y1 := wp.X, wp.Y
		x2, y2 := origin.X(), origin.Y()
		waypoints[i] = position.NewAbsolute(x1+x2, y1+y2, height)

		wpq.Push(waypoints[i])
	}

	if !wpq.Empty() {
		player.InteractionQueue().InterruptAndAppend(wpq)
	}
}
