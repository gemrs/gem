package protocol

import (
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/position"
)

const PlayerViewDistance = 1

type World interface {
	EntitiesOnTile(p *position.Absolute) []entity.Entity
	Sector(s *position.Sector) Sector
	Sectors(s []*position.Sector) []Sector
	FindPlayerSlot() int
	SetPlayerSlot(i int, p Player)
	GetPlayers() []Player
}

type Sector interface {
	Collection() *entity.Collection
	Add(entity entity.Entity)
	Remove(entity entity.Entity)
	Position() *position.Sector
}
