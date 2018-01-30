package packet

import (
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol"
)

func init() {
	registerHandler((*protocol.InboundPlayerWalk)(nil), player_walk)
}

func player_walk(player protocol.Player, message server.Message) {
	walkPacket := message.(*protocol.InboundPlayerWalk)

	height := player.Position().Z()
	origin := position.NewAbsolute(player.Position().X(), player.Position().Y(), height)

	wpq := player.WaypointQueue()
	wpq.Clear()
	wpq.Push(origin)
	wpq.SetRunning(walkPacket.Running)
	/*	waypoints := make([]*position.Absolute, 1)
		for i, wp := range walkPacket.Waypoints {
			x1, y1 := wp.X, wp.Y
			x2, y2 := origin.X(), origin.Y()
			waypoints[i] = position.NewAbsolute(x1+x2, y1+y2, height)

			wpq.Push(waypoints[i])
		}*/

	wpq.Push(position.NewAbsolute(walkPacket.X, walkPacket.Y, height))

	if !wpq.Empty() {
		player.InteractionQueue().InterruptAndAppend(wpq)
	}
}
