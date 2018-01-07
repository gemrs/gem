package player

import (
	"github.com/gemrs/gem/gem/encoding"
	game_event "github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/interface/entity"
	"github.com/gemrs/gem/gem/game/interface/player"
	"github.com/gemrs/gem/gem/game/position"
	game_protocol "github.com/gemrs/gem/gem/protocol/game"
)

func (client *Player) LoadProfile() {
	profile := client.Profile().(*Profile)
	client.SetPosition(profile.Position())
	client.SetAppearance(profile.Appearance())

	game_event.PlayerLoadProfile.NotifyObservers(client, client.Profile().(*Profile))
}

// FinishInit is called once the player has finished the low level login sequence
func (client *Player) FinishInit() {
	client.Conn().Write <- &game_protocol.OutboundPlayerInit{
		Membership: encoding.Uint8(1),
		Index:      encoding.Uint16(client.Index()),
	}
}

// Cleanup is called when the player logs out
func (client *Player) Cleanup() {

}

func (client *Player) SyncEntityList() {
	sectorPositions := client.sector.Position().SurroundingSectors(1)
	sectors := client.world.Sectors(sectorPositions)
	for _, s := range sectors {
		for _, e := range s.Adding().Slice() {
			client.visibleEntities.Add(e)
		}

		for _, e := range s.Removing().Slice() {
			client.visibleEntities.Remove(e)
		}
	}
}

func (client *Player) SectorChange() {
	oldList := client.sector.Position().SurroundingSectors(1)
	client.sector.Remove(client)
	client.sector = client.world.Sector(client.Position().Sector())
	client.sector.Add(client)
	newList := client.sector.Position().SurroundingSectors(1)

	removedPositions, addedPositions := position.SectorListDelta(oldList, newList)
	removed, added := client.world.Sectors(removedPositions), client.world.Sectors(addedPositions)

	for _, s := range removed {
		client.visibleEntities.RemoveAll(s.Collection)
	}

	for _, s := range added {
		client.visibleEntities.AddAll(s.Collection)
	}
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

func (client *Player) UpdateWaypointQueue() {
	client.WaypointQueue().Tick(client)
	client.Profile().SetPosition(client.Position())
}

func (client *Player) SendPlayerSync() {
	client.Conn().Write <- &game_protocol.PlayerUpdate{
		OurPlayer: player.Snapshot(client),
	}
}

func (client *Player) UpdateVisibleEntities() {
	client.visibleEntities.Update()
}
