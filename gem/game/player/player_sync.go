package player

import (
	"github.com/gemrs/gem/gem/encoding"
	game_event "github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/interface/entity"
	"github.com/gemrs/gem/gem/game/position"
	game_protocol "github.com/gemrs/gem/gem/protocol/game"
)

const PlayerViewDistance = 1

func (client *Player) LoadProfile() {
	profile := client.Profile().(*Profile)
	profile.setPlayer(client)
	client.SetPosition(profile.Position())

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
	sectorPositions := client.sector.Position().SurroundingSectors(PlayerViewDistance)
	sectors := client.world.Sectors(sectorPositions)
	allAdded := entity.NewSet()
	allRemoved := entity.NewSet()

	for _, s := range sectors {
		client.visibleEntities.AddAll(s.Collection)
		allAdded.AddAll(s.Adding())
		allRemoved.AddAll(s.Removing())
	}

	for _, e := range allRemoved.Slice() {
		// Entity was both added and removed, probably went from one sector to another.
		if allAdded.Contains(e) {
			continue
		}

		client.visibleEntities.Remove(e)
	}
}

func (client *Player) SectorChange() {
	oldList := client.sector.Position().SurroundingSectors(PlayerViewDistance)

	client.sector.Remove(client)
	client.sector = client.world.Sector(client.Position().Sector())
	client.sector.Add(client)

	newList := client.sector.Position().SurroundingSectors(PlayerViewDistance)

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

func (client *Player) ProcessChatQueue() {
	if len(client.chatQueue) > 0 {
		client.chatQueue = client.chatQueue[1:]
	}
}

func (client *Player) SendPlayerSync() {
	client.Conn().Write <- game_protocol.NewPlayerUpdateBlock(client)
}

func (client *Player) UpdateVisibleEntities() {
	client.visibleEntities.Update()
}
