package world

import (
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/protocol"
)

type Instance struct {
	sectors map[position.SectorHash]protocol.Sector
	players [protocol.MaxPlayers]protocol.Player
}

func NewInstance() *Instance {
	return &Instance{
		sectors: make(map[position.SectorHash]protocol.Sector),
	}
}

func (i *Instance) FindPlayerSlot() int {
	for i, p := range i.players {
		if i > 0 && p == nil {
			return i
		}
	}
	return -1
}

func (i *Instance) SetPlayerSlot(slot int, p protocol.Player) {
	i.players[slot] = p
}

func (i *Instance) GetPlayers() []protocol.Player {
	return i.players[:]
}

// Sector returns a sector instance for a given position.Sector.
// If the sector is not currently active, the sector is instantiated first.
func (i *Instance) Sector(s *position.Sector) protocol.Sector {
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
		sector.Collection().Update()
	}
}

func (instance *Instance) Sectors(s []*position.Sector) []protocol.Sector {
	list := make([]protocol.Sector, len(s))
	for i, s := range s {
		list[i] = instance.Sector(s)
	}
	return list
}

func (instance *Instance) EntitiesOnTile(p *position.Absolute) []entity.Entity {
	entities := make([]entity.Entity, 0)
	sectorPos := p.Sector()
	sector := instance.Sector(sectorPos)

	for _, entity := range sector.Collection().Entities().Slice() {
		if entity.Position().Compare(p) {
			entities = append(entities, entity)
		}
	}

	return entities
}

func (instance *Instance) AllEntities(typ entity.EntityType) []entity.Entity {
	entities := make([]entity.Entity, 0)
	for _, s := range instance.sectors {
		entities = append(entities, s.Collection().Entities().Filter(typ).Slice()...)
	}
	return entities
}
