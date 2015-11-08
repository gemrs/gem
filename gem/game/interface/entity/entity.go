package entity

import (
	"github.com/sinusoids/gem/gem/game/position"
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

type EntityType string

const (
	IncompleteType EntityType = "incomplete"
	PlayerType                = "player"
)

// An Entity is a 'thing' within the world, with a position, and an index.
type Entity interface {
	position.Positionable
	Region() *position.Region
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
