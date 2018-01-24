package player

import (
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/protocol"
)

func buildPlayerUpdate(player *Player) protocol.PlayerUpdate {
	block := protocol.PlayerUpdate{
		Others:   make(map[int]protocol.PlayerUpdateData),
		Removing: make(map[int]bool),
	}
	block.Me = buildPlayerUpdateData(player)

	visibleEntities := player.VisibleEntities()
	updatingPlayers := visibleEntities.Entities().Clone()
	updatingPlayers.RemoveAll(visibleEntities.Adding())
	updatingPlayers = updatingPlayers.Filter(entity.PlayerType)
	updatingPlayers.Remove(player)

	for _, other := range updatingPlayers.Slice() {
		block.Others[other.Index()] = buildPlayerUpdateData(other.(*Player))
		block.Updating = append(block.Updating, other.Index())
		block.Removing[other.Index()] = visibleEntities.Removing().Contains(other)
	}

	for _, other := range visibleEntities.Adding().Filter(entity.PlayerType).Slice() {
		if player.Index() == other.Index() {
			continue
		}
		block.Adding = append(block.Adding, other.Index())
		block.Others[other.Index()] = buildPlayerUpdateData(other.(*Player))
	}

	return block
}

func buildPlayerUpdateData(player *Player) protocol.PlayerUpdateData {
	return protocol.PlayerUpdateData{
		Index:            player.Index(),
		Username:         player.Profile().Username(),
		Flags:            player.Flags(),
		Position:         player.Position(),
		LoadedRegion:     player.LoadedRegion(),
		WaypointQueue:    player.WaypointQueue(),
		ChatMessageQueue: player.ChatMessageQueue(),
		Rights:           player.Profile().Rights(),
		Appearance:       player.Profile().Appearance(),
		Animations:       player.Animations(),
		Skills:           player.Profile().Skills(),
	}
}
