package impl

import (
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/protocol"
)

func buildPlayerUpdate(player *Player) protocol.PlayerUpdate {
	block := protocol.PlayerUpdate{
		Others:   make(map[int]protocol.Player),
		Removing: make(map[int]bool),
	}
	block.Me = player

	visibleEntities := player.VisibleEntities()

	for _, e := range visibleEntities.Entities().Slice() {
		if e.EntityType() != entity.PlayerType {
			continue
		}
		block.Visible = append(block.Visible, e.Index())
		block.Others[e.Index()] = e.(protocol.Player)
	}

	updatingPlayers := visibleEntities.Entities().Clone()
	updatingPlayers.RemoveAll(visibleEntities.Adding())
	updatingPlayers = updatingPlayers.Filter(entity.PlayerType)
	updatingPlayers.Remove(player)

	for _, other := range updatingPlayers.Slice() {
		block.Updating = append(block.Updating, other.Index())
		block.Removing[other.Index()] = visibleEntities.Removing().Contains(other)
	}

	for _, other := range visibleEntities.Adding().Filter(entity.PlayerType).Slice() {
		if player.Index() == other.Index() {
			continue
		}
		block.Adding = append(block.Adding, other.Index())
	}

	return block
}
