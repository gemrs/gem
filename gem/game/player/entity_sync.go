package player

import (
	"github.com/gemrs/gem/gem/encoding"
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/item"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/protocol/game_protocol"
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

func (player *Player) SendGroundItemSync() {
	newItems := player.visibleEntities.Adding().Filter(entity.GroundItemType)
	for _, entity := range newItems.Slice() {
		player.setUpdatingSector(entity.Position().Sector())
		itemPos := entity.Position()

		item := entity.(*item.GroundItem)
		stack := item.Item()

		dx, dy := itemPos.SectorLocal()

		player.Conn().Write <- &game_protocol.OutboundCreateGroundItem{
			ItemID:         encoding.Uint16(stack.Definition().Id()),
			PositionOffset: encoding.Uint8((dx << 4) + dy),
			Count:          encoding.Uint16(stack.Count()),
		}
	}

	removingItems := player.visibleEntities.Removing().Filter(entity.GroundItemType)
	for _, entity := range removingItems.Slice() {
		player.setUpdatingSector(entity.Position().Sector())
		itemPos := entity.Position()

		item := entity.(*item.GroundItem)
		stack := item.Item()

		dx, dy := itemPos.SectorLocal()

		player.Conn().Write <- &game_protocol.OutboundRemoveGroundItem{
			ItemID:         encoding.Uint16(stack.Definition().Id()),
			PositionOffset: encoding.Uint8((dx << 4) + dy),
		}
	}
}

func (player *Player) setUpdatingSector(s *position.Sector) {
	region := player.loadedRegion
	regionOrigin := region.Origin()
	dx, dy, _ := s.Min().Delta(regionOrigin.Min())
	player.Conn().Write <- &game_protocol.OutboundSetUpdatingSector{
		PositionX: encoding.Uint8(dx),
		PositionY: encoding.Uint8(dy),
	}
}
