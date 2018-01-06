package entity

import (
	"github.com/gemrs/gem/gem/game/interface/entity"
	"github.com/gemrs/gem/gem/game/position"
)

type GenericMob struct {
	waypointQueue entity.WaypointQueue
	position      *position.Absolute
	flags         entity.Flags
}

func NewGenericMob(wpq entity.WaypointQueue) *GenericMob {
	return &GenericMob{
		waypointQueue: wpq,
	}
}

func (mob *GenericMob) SetNextStep(next *position.Absolute) {
	if (mob.flags & entity.MobFlagWalkUpdate) != 0 { // Called twice in one cycle - mob is running
		mob.flags &= ^entity.MobFlagWalkUpdate
		mob.flags |= entity.MobFlagRunUpdate
	} else {
		mob.flags |= entity.MobFlagWalkUpdate
	}
}

// Position returns the absolute position of the mob
func (mob *GenericMob) Position() *position.Absolute {
	return mob.position
}

// Flags returns the mob update flags for this mob
func (mob *GenericMob) Flags() entity.Flags {
	return mob.flags
}

// SetFlags ORs the given flags with the player's current update flags
func (mob *GenericMob) SetFlags(f entity.Flags) {
	mob.flags |= f
}

func (mob *GenericMob) ClearFlags() {
	mob.flags = 0
}

// SetPosition warps the mob to a given location
func (mob *GenericMob) SetPosition(pos *position.Absolute) {
	mob.position = pos
}

// AppearanceUpdate is called when the player's appearance changes
// It ensures the player's appearance is synced at next update
func (mob *GenericMob) AppearanceChange() {
	mob.SetFlags(entity.MobFlagIdentityUpdate)
}

func (mob *GenericMob) WaypointQueue() entity.WaypointQueue {
	return mob.waypointQueue
}

// EntityType identifies what kind of entity this entity is
func (mob *GenericMob) EntityType() entity.EntityType {
	return entity.IncompleteType
}

func (mob *GenericMob) SectorChange() {
	panic("not implemented")
}

func (mob *GenericMob) RegionChange() {
	panic("not implemented")
}
