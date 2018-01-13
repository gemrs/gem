package world

import (
	"github.com/gemrs/gem/gem/game/entity"
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

func (instance *Instance) EntitiesOnTile(p *position.Absolute) []entity.Entity {
	entities := make([]entity.Entity, 0)
	sectorPos := p.Sector()
	sector := instance.Sector(sectorPos)

	for _, entity := range sector.Entities().Slice() {
		if entity.Position().Compare(p) {
			entities = append(entities, entity)
		}
	}

	return entities
}

func (instance *Instance) AllEntities(typ entity.EntityType) []entity.Entity {
	entities := make([]entity.Entity, 0)
	for _, s := range instance.sectors {
		entities = append(entities, s.Entities().Filter(typ).Slice()...)
	}
	return entities
}
