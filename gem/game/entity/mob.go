package entity

import (
	"github.com/sinusoids/gem/gem/game/interface/entity"
	"github.com/sinusoids/gem/gem/game/position"

	"github.com/qur/gopy/lib"
)

//go:generate gopygen -type GenericMob -excfield "^[a-z].*" $GOFILE

type GenericMob struct {
	py.BaseObject

	position *position.Absolute
	region   *position.Region
	flags    entity.Flags
}

func (mob *GenericMob) Init() error {
	var err error
	mob.region, err = position.NewRegion(nil)
	if err != nil {
		return err
	}
	return nil
}

func (mob *GenericMob) SetNextStep(pos *position.Absolute) {

}

// Position returns the absolute position of the mob
func (mob *GenericMob) Position() *position.Absolute {
	return mob.position
}

// WalkDirection returns the mob's current and (in the case of running) last walking direction
func (mob *GenericMob) WalkDirection() (current int, last int) {
	return 0, 0
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
	mob.region = pos.RegionOf()
}
