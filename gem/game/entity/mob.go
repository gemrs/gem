package entity

import (
	"github.com/gemrs/gem/gem/game/position"
)

type GenericMob struct {
	waypointQueue WaypointQueue
	position      *position.Absolute
	flags         Flags
}

func NewGenericMob(wpq WaypointQueue) *GenericMob {
	return &GenericMob{
		waypointQueue: wpq,
	}
}

func (mob *GenericMob) SetNextStep(next *position.Absolute) {
	if (mob.flags & MobFlagWalkUpdate) != 0 { // Called twice in one cycle - mob is running
		mob.flags &= ^MobFlagWalkUpdate
		mob.flags |= MobFlagRunUpdate
	} else {
		mob.flags |= MobFlagWalkUpdate
	}
}

// Position returns the absolute position of the mob
func (mob *GenericMob) Position() *position.Absolute {
	return mob.position
}

// SetPosition warps the mob to a given location
func (mob *GenericMob) SetPosition(pos *position.Absolute) {
	mob.position = pos
}

// Flags returns the mob update flags for this mob
func (mob *GenericMob) Flags() Flags {
	return mob.flags
}

// SetFlags ORs the given flags with the player's current update flags
func (mob *GenericMob) SetFlags(f Flags) {
	mob.flags |= f
}

func (mob *GenericMob) ClearFlags() {
	mob.flags = 0
}

// AppearanceUpdate is called when the player's appearance changes
// It ensures the player's appearance is synced at next update
func (mob *GenericMob) SetAppearanceChanged() {
	mob.SetFlags(MobFlagIdentityUpdate)
}

func (mob *GenericMob) WaypointQueue() WaypointQueue {
	return mob.waypointQueue
}

// EntityType identifies what kind of entity this entity is
func (mob *GenericMob) EntityType() EntityType {
	return IncompleteType
}
