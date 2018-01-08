package world

import (
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/willow/log"
)

type Sector struct {
	*entity.Collection
	position *position.Sector
	logger   log.Log
}

func NewSector(position *position.Sector) *Sector {
	return &Sector{
		Collection: entity.NewCollection(),
		position:   position,
		logger:     log.New("world/sector", log.MapContext{"position": position}),
	}
}

func (s *Sector) Add(entity entity.Entity) {
	s.Collection.Add(entity)
}

func (s *Sector) Remove(entity entity.Entity) {
	s.Collection.Remove(entity)
}

func (s *Sector) Position() *position.Sector {
	return s.position
}
