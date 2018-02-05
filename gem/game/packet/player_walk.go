package packet

import (
	"fmt"

	astar "github.com/beefsack/go-astar"
	"github.com/gemrs/gem/gem/game/data"
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

	destination := position.NewAbsolute(walkPacket.X, walkPacket.Y, height)

	startTile := data.GetCollisionTile(origin.X(), origin.Y(), height)
	endTile := data.GetCollisionTile(destination.X(), destination.Y(), height)

	fmt.Printf("startTile is %#v %p, endTile is %#v %p\n", startTile, startTile, endTile, endTile)

	if endTile.Blocked() {
		return
	}

	path, _, found := astar.Path(startTile, endTile)
	fmt.Printf("found path? %v %v\n", found, path)
	if !found {
		return
	}

	for i := len(path) - 1; i >= 0; i-- {
		wp := path[i].(*data.CollisionTile)

		pos := position.NewAbsolute(wp.AbsX, wp.AbsY, height)
		wpq.Push(pos)
	}

	if !wpq.Empty() {
		player.InteractionQueue().InterruptAndAppend(wpq)
	}
}
