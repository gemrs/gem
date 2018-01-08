package player

import (
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/position"
)

const PlayerViewDistance = 1

func (player *Player) SyncEntityList() {
	sectorPositions := player.sector.Position().SurroundingSectors(PlayerViewDistance)
	sectors := player.world.Sectors(sectorPositions)
	allAdded := entity.NewSet()
	allRemoved := entity.NewSet()

	for _, s := range sectors {
		player.visibleEntities.AddAll(s.Collection)
		allAdded.AddAll(s.Adding())
		allRemoved.AddAll(s.Removing())
	}

	for _, e := range allRemoved.Slice() {
		// Entity was both added and removed, probably went from one sector to another.
		if allAdded.Contains(e) {
			continue
		}

		player.visibleEntities.Remove(e)
	}
}

func (player *Player) UpdateVisibleEntities() {
	player.visibleEntities.Update()
}

// Called when the player's movement has caused them to enter a new sector
func (player *Player) onSectorChange() {
	oldList := player.sector.Position().SurroundingSectors(PlayerViewDistance)

	player.sector.Remove(player)
	player.sector = player.world.Sector(player.Position().Sector())
	player.sector.Add(player)

	newList := player.sector.Position().SurroundingSectors(PlayerViewDistance)

	removedPositions, addedPositions := position.SectorListDelta(oldList, newList)
	removed, added := player.world.Sectors(removedPositions), player.world.Sectors(addedPositions)

	for _, s := range removed {
		player.visibleEntities.RemoveAll(s.Collection)
	}

	for _, s := range added {
		player.visibleEntities.AddAll(s.Collection)
	}
}
