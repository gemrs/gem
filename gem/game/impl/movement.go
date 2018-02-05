package impl

import (
	astar "github.com/beefsack/go-astar"
	"github.com/gemrs/gem/gem/game/data"
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/protocol"
)

// SetPosition warps the mob to a given location
func (player *Player) SetPosition(pos *position.Absolute) {
	oldSector := player.sector.Position()
	player.GenericMob.SetPosition(pos)

	dx, dy, dz := oldSector.Delta(pos.Sector())

	if dx >= 1 || dy >= 1 || dz >= 1 {
		player.onSectorChange()
	}

	loadedRegion := player.LoadedRegion()
	dx, dy, dz = loadedRegion.SectorDelta(pos.RegionOf())

	player.Profile().SetPosition(player.Position())

	if dx >= 5 || dy >= 5 || dz >= 1 {
		player.RegionChange()
	}
}

func (player *Player) UpdateInteractionQueue() {
	player.InteractionQueue().Tick(player)
}

// Called by the waypoint processor to indicate the player's next step
func (player *Player) SetNextStep(next *position.Absolute) {
	player.SetPosition(next)
	player.GenericMob.SetNextStep(next)
}

// RegionUpdate is called when the player enters a new region
// It sends the region update packet and sets the correct update flags
func (player *Player) RegionChange() {
	player.loadedRegion = player.Position().RegionOf()

	player.Send(protocol.OutboundRegionUpdate{
		ProtoData: player.protoData,
		Player:    player,
	})

	player.SetFlags(entity.MobFlagRegionUpdate)
}

func (player *Player) SetRunning(bool) {
	wpq := player.WaypointQueue()
	wpq.SetRunning(true)
}

//glua:bind
func (player *Player) SetWalkDestination(destination *position.Absolute) bool {
	origin := player.Position()

	wpq := player.WaypointQueue()
	wpq.Clear()
	wpq.Push(origin)

	startTile := data.GetCollisionTile(origin.X(), origin.Y(), origin.Z())
	endTile := data.GetCollisionTile(destination.X(), destination.Y(), destination.Z())

	if startTile == nil || endTile == nil {
		return false
	}

	if endTile.Blocked() {
		return false
	}

	path, _, found := astar.Path(startTile, endTile)
	if !found {
		return false
	}

	for i := len(path) - 1; i >= 0; i-- {
		wp := path[i].(*data.CollisionTile)

		pos := position.NewAbsolute(wp.AbsX, wp.AbsY, wp.Z)
		wpq.Push(pos)
	}

	if !wpq.Empty() {
		player.InteractionQueue().InterruptAndAppend(wpq)
	}

	return true
}
