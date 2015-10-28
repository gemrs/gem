package position

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/pybind"
)

var AbsoluteDef = pybind.Define("Absolute", (*Absolute)(nil))
var RegisterAbsolute = pybind.GenerateRegisterFunc(AbsoluteDef)
var NewAbsolute = pybind.GenerateConstructor(AbsoluteDef).(func(int, int, int) *Absolute)

var SectorDef = pybind.Define("Sector", (*Sector)(nil))
var RegisterSector = pybind.GenerateRegisterFunc(SectorDef)
var NewSector = pybind.GenerateConstructor(SectorDef).(func(int, int, int) *Sector)

var RegionDef = pybind.Define("Region", (*Region)(nil))
var RegisterRegion = pybind.GenerateRegisterFunc(RegionDef)
var NewRegion = pybind.GenerateConstructor(RegionDef).(func(*Sector) *Region)

var LocalDef = pybind.Define("Local", (*Local)(nil))
var RegisterLocal = pybind.GenerateRegisterFunc(LocalDef)
var NewLocal = pybind.GenerateConstructor(LocalDef).(func(int, int, int, *Region) *Local)

func (pos *Absolute) Py_X(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(pos.X)
	return fn(args, kwds)
}

func (pos *Absolute) Py_Y(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(pos.Y)
	return fn(args, kwds)
}

func (pos *Absolute) Py_Z(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(pos.Z)
	return fn(args, kwds)
}

func (pos *Absolute) PyStr() string {
	return pos.String()
}

func (pos *Absolute) Py_Sector(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(pos.Sector)
	return fn(args, kwds)
}

func (pos *Absolute) Py_RegionOf(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(pos.RegionOf)
	return fn(args, kwds)
}

func (pos *Absolute) Py_LocalTo(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(pos.LocalTo)
	return fn(args, kwds)
}

func (sector *Sector) Py_X(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(sector.X)
	return fn(args, kwds)
}

func (sector *Sector) Py_Y(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(sector.Y)
	return fn(args, kwds)
}

func (sector *Sector) Py_Z(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(sector.Z)
	return fn(args, kwds)
}

func (region *Region) Py_Origin(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(region.Origin)
	return fn(args, kwds)
}

func (region *Region) Py_Rebase(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(region.Rebase)
	return fn(args, kwds)
}

func (region *Region) Py_SectorDelta(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(region.SectorDelta)
	return fn(args, kwds)
}

func (local *Local) Py_X(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(local.X)
	return fn(args, kwds)
}

func (local *Local) Py_Y(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(local.Y)
	return fn(args, kwds)
}

func (local *Local) Py_Z(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(local.Z)
	return fn(args, kwds)
}

func (local *Local) Py_Region(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(local.Region)
	return fn(args, kwds)
}

func (local *Local) Py_Absolute(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(local.Absolute)
	return fn(args, kwds)
}
