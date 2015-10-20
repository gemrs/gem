package entity

import (
	"gem/game/player"
	"gem/game/position"
	"gem/log"
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
	Region() *position.Region
}

type Mob interface {
	Entity
	Flags() Flags
	WalkDirection() (current int, last int)
}

// Player is an Entity representing a player
type Player interface {
	Mob
	Profile() *player.Profile
	Session() *player.Session
	Log() *log.Module
}
