package world

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/game/interface/entity"
	"github.com/sinusoids/gem/gem/game/position"
	"github.com/sinusoids/gem/gem/log"
)

type Sector struct {
	py.BaseObject

	*entity.Collection
	position *position.Sector
	logger   log.Log
}

func (s *Sector) Init(position *position.Sector) {
	s.Collection = entity.NewCollection()
	s.position = position
	s.logger = log.New("world/sector", log.MapContext{"position": position})
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

// Active determines if this sector is currently in use.
// An 'in use' sector is one which contains entities or other information.
// If a sector is not in use, it can usually be garbage collected.
func (s *Sector) Active() bool {
	return s.Collection.Size() > 0
}
