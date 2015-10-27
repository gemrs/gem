package entity

import (
	"github.com/sinusoids/gem/gem/game/interface/entity"
	"github.com/sinusoids/gem/gem/game/position"

	"github.com/qur/gopy/lib"
)

//go:generate gopygen -type GenericMob -excfield "^[a-z].*" $GOFILE

type GenericMob struct {
	py.BaseObject

	waypointQueue entity.WaypointQueue
	position      *position.Absolute
	region        *position.Region
	flags         entity.Flags
}

func (mob *GenericMob) Init(wpq entity.WaypointQueue) error {
	var err error
	mob.region, err = position.NewRegion(nil)
	if err != nil {
		return err
	}

	mob.waypointQueue = wpq
	return nil
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

// Region returns the mob's current surrounding region
func (mob *GenericMob) Region() *position.Region {
	return mob.region
}

// SetPosition warps the mob to a given location
func (mob *GenericMob) SetPosition(pos *position.Absolute) {
	mob.position = pos

	newRegion := pos.RegionOf()
	dx, dy, dz := newRegion.SectorDelta(mob.Region())
	if dx >= 5 || dy >= 5 || dz >= 1 {
		mob.region = newRegion
	}
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
