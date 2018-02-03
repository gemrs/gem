package impl

import (
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/protocol"
)

func (player *Player) SyncEntityList() {
	sectorPositions := player.sector.Position().SurroundingSectors(protocol.PlayerViewDistance)
	sectors := player.world.Sectors(sectorPositions)
	allAdded := entity.NewSet()
	allRemoved := entity.NewSet()

	for _, s := range sectors {
		//		fmt.Printf("sector %v has %#v\n", s.Position(), s.Collection().Entities().Slice())
		player.visibleEntities.AddAll(s.Collection())
		allAdded.AddAll(s.Collection().Adding())
		allRemoved.AddAll(s.Collection().Removing())
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
	oldList := player.sector.Position().SurroundingSectors(protocol.PlayerViewDistance)

	player.sector.Remove(player)
	player.sector = player.world.Sector(player.Position().Sector())
	player.sector.Add(player)

	newList := player.sector.Position().SurroundingSectors(protocol.PlayerViewDistance)

	removedPositions, addedPositions := position.SectorListDelta(oldList, newList)
	removed, added := player.world.Sectors(removedPositions), player.world.Sectors(addedPositions)

	for _, s := range removed {
		player.visibleEntities.RemoveAll(s.Collection())
	}

	for _, s := range added {
		player.visibleEntities.AddAll(s.Collection())
	}
}

func (player *Player) SendGroundItemSync() {
	newItems := player.visibleEntities.Adding().Filter(entity.GroundItemType)

	for _, entity := range newItems.Slice() {
		player.setUpdatingSector(entity.Position().Sector())
		itemPos := entity.Position()

		item := entity.(*GroundItem)
		stack := item.Item()

		dx, dy := itemPos.SectorLocal()

		player.Send(protocol.OutboundCreateGroundItem{
			ItemID:         stack.Definition().Id(),
			PositionOffset: (dx << 4) + dy,
			Count:          stack.Count(),
		})
	}

	removingItems := player.visibleEntities.Removing().Filter(entity.GroundItemType)
	for _, entity := range removingItems.Slice() {
		player.setUpdatingSector(entity.Position().Sector())
		itemPos := entity.Position()

		item := entity.(*GroundItem)
		stack := item.Item()

		dx, dy := itemPos.SectorLocal()

		player.Send(protocol.OutboundRemoveGroundItem{
			ItemID:         stack.Definition().Id(),
			PositionOffset: (dx << 4) + dy,
		})
	}
}

func (player *Player) setUpdatingSector(s *position.Sector) {
	region := player.loadedRegion
	regionOrigin := region.Origin()
	dx, dy, _ := s.Min().Delta(regionOrigin.Min())
	player.Send(protocol.OutboundSetUpdatingTile{
		PositionX: dx,
		PositionY: dy,
	})
}
