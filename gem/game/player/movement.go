package player

import (
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

	if dx >= 5 || dy >= 5 || dz >= 1 {
		player.RegionChange()
	}

	player.Profile().SetPosition(player.Position())
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

	player.Conn().Write <- protocol.OutboundRegionUpdate{
		Position:    player.Position(),
		PlayerIndex: player.Index(),
	}

	player.SetFlags(entity.MobFlagRegionUpdate)
}
