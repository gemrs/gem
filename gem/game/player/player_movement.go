package player

import (
	"github.com/gemrs/gem/gem/encoding"
	"github.com/gemrs/gem/gem/game/interface/entity"
	"github.com/gemrs/gem/gem/game/position"
	game_protocol "github.com/gemrs/gem/gem/protocol/game"
)

// SetPosition warps the mob to a given location
func (client *Player) SetPosition(pos *position.Absolute) {
	oldSector := client.sector.Position()
	client.GenericMob.SetPosition(pos)

	dx, dy, dz := oldSector.Delta(pos.Sector())

	if dx >= 1 || dy >= 1 || dz >= 1 {
		client.onSectorChange()
	}

	loadedRegion := client.LoadedRegion()
	dx, dy, dz = loadedRegion.SectorDelta(pos.RegionOf())

	if dx >= 5 || dy >= 5 || dz >= 1 {
		client.RegionChange()
	}
}

func (client *Player) UpdateWaypointQueue() {
	client.WaypointQueue().Tick(client)
	// FIXME why update the profile here? why not SetPosition
	client.Profile().SetPosition(client.Position())
}

// Called by the waypoint processor to indicate the player's next step
func (client *Player) SetNextStep(next *position.Absolute) {
	client.SetPosition(next)
	client.GenericMob.SetNextStep(next)
}

// RegionUpdate is called when the player enters a new region
// It sends the region update packet and sets the correct update flags
func (client *Player) RegionChange() {
	client.loadedRegion = client.Position().RegionOf()

	sector := client.Position().Sector()
	client.Conn().Write <- &game_protocol.OutboundRegionUpdate{
		SectorX: encoding.Uint16(sector.X()),
		SectorY: encoding.Uint16(sector.Y()),
	}

	client.SetFlags(entity.MobFlagRegionUpdate)
}
