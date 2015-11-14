package entity

import (
	"github.com/gemrs/gem/gem/game/position"
)

type Flags int

const (
	MobFlagRegionUpdate   Flags = (1 << 0)
	MobFlagWalkUpdate     Flags = (1 << 1)
	MobFlagRunUpdate      Flags = (1 << 2)
	MobFlagIdentityUpdate Flags = (1 << 4)
	MobFlagChatUpdate     Flags = (1 << 7)
	MobFlagMovementUpdate Flags = (MobFlagRegionUpdate | MobFlagWalkUpdate | MobFlagRunUpdate)
)

// An Entity is a 'thing' within the world, with a position, and an index.
type Entity interface {
	position.Positionable
	EntityType() EntityType
	RegionChange()
	SectorChange()
	Index() int
}

type Movable interface {
	Entity
	Flags() Flags
	SetFlags(Flags)
	ClearFlags()
	SetNextStep(*position.Absolute)
	AppearanceChange()
	WaypointQueue() WaypointQueue
}
