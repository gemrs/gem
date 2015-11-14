package world

import (
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/pybind"
)

var SectorDef = pybind.Define("Sector", (*Sector)(nil))
var RegisterSector = pybind.GenerateRegisterFunc(SectorDef)
var NewSector = pybind.GenerateConstructor(SectorDef).(func(*position.Sector) *Sector)
