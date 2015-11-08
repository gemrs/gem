package world

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/game/position"
	"github.com/sinusoids/gem/pybind"
)

var SectorDef = pybind.Define("Sector", (*Sector)(nil))
var RegisterSector = pybind.GenerateRegisterFunc(SectorDef)
var NewSector = pybind.GenerateConstructor(SectorDef).(func(*position.Sector) *Sector)

func (sector *Sector) PyGet_active() (py.Object, error) {
	fn := pybind.Wrap(sector.Active)
	return fn(nil, nil)
}
