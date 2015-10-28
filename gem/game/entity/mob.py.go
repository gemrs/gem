package entity

import (
	"github.com/sinusoids/gem/gem/game/interface/entity"
	"github.com/sinusoids/gem/pybind"
)

var GenericMobDef = pybind.Define("GenericMob", (*GenericMob)(nil))
var RegisterGenericMob = pybind.GenerateRegisterFunc(GenericMobDef)
var NewGenericMob = pybind.GenerateConstructor(GenericMobDef).(func(entity.WaypointQueue) *GenericMob)
