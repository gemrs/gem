package player

import (
	"github.com/gemrs/gem/gem/game/interface/entity"
	"github.com/gemrs/gem/gem/game/position"
)

const PlayerViewDistance = 1

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

func (client *Player) UpdateVisibleEntities() {
	client.visibleEntities.Update()
}

// Called when the player's movement has caused them to enter a new sector
func (client *Player) onSectorChange() {
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
