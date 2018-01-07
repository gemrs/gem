package world

import (
	"github.com/gemrs/gem/gem/game/interface/entity"
	"github.com/gemrs/gem/gem/game/position"
)

type Instance struct {
	sectors map[position.SectorHash]*Sector
}

func NewInstance() *Instance {
	return &Instance{
		sectors: make(map[position.SectorHash]*Sector),
	}
}

// Sector returns a sector instance for a given position.Sector.
// If the sector is not currently active, the sector is instantiated first.
func (i *Instance) Sector(s *position.Sector) *Sector {
	hash := s.HashCode()
	if sector, ok := i.sectors[hash]; ok {
		return sector
	}

	// Sector not yet tracked; Create a new sector
	i.sectors[hash] = NewSector(s)
	return i.sectors[hash]
}

func (i *Instance) UpdateEntityCollections() {
	for _, sector := range i.sectors {
		sector.Update()
	}
}

func (instance *Instance) Sectors(s []*position.Sector) []*Sector {
	list := make([]*Sector, len(s))
	for i, s := range s {
		list[i] = instance.Sector(s)
	}
	return list
}

func (instance *Instance) ForEachPlayer(fn func(p entity.Entity)) {
	for _, s := range instance.sectors {
		for _, e := range s.Entities().Filter(entity.PlayerType).Slice() {
			fn(e)
		}
	}
}
