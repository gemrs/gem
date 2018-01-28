package world

import (
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/willow/log"
)

type Sector struct {
	collection *entity.Collection
	position   *position.Sector
	logger     log.Log
}

func NewSector(position *position.Sector) *Sector {
	return &Sector{
		collection: entity.NewCollection(),
		position:   position,
		logger:     log.New("world/sector", log.MapContext{"position": position}),
	}
}

func (s *Sector) Collection() *entity.Collection {
	return s.collection
}

func (s *Sector) Add(entity entity.Entity) {
	s.collection.Add(entity)
}

func (s *Sector) Remove(entity entity.Entity) {
	s.collection.Remove(entity)
}

func (s *Sector) Position() *position.Sector {
	return s.position
}
