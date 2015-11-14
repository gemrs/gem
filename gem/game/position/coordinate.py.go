package position

import (
	"github.com/qur/gopy/lib"

	"github.com/gemrs/gem/pybind"
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

var AreaDef = pybind.Define("Area", (*Area)(nil))
var RegisterArea = pybind.GenerateRegisterFunc(AreaDef)
var NewArea = pybind.GenerateConstructor(AreaDef).(func(*Sector, int, int) *Area)

func (pos *Absolute) PyGet_x() (py.Object, error) {
	fn := pybind.Wrap(pos.X)
	return fn(nil, nil)
}

func (pos *Absolute) PyGet_y() (py.Object, error) {
	fn := pybind.Wrap(pos.Y)
	return fn(nil, nil)
}

func (pos *Absolute) PyGet_z() (py.Object, error) {
	fn := pybind.Wrap(pos.Z)
	return fn(nil, nil)
}

func (pos *Absolute) PyStr() string {
	return pos.String()
}

func (pos *Absolute) PyGet_sector() (py.Object, error) {
	fn := pybind.Wrap(pos.Sector)
	return fn(nil, nil)
}

func (pos *Absolute) PyGet_region() (py.Object, error) {
	fn := pybind.Wrap(pos.RegionOf)
	return fn(nil, nil)
}

func (pos *Absolute) Py_local_to(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(pos.LocalTo)
	return fn(args, kwds)
}

func (sector *Sector) PyGet_x() (py.Object, error) {
	fn := pybind.Wrap(sector.X)
	return fn(nil, nil)
}

func (sector *Sector) PyGet_y() (py.Object, error) {
	fn := pybind.Wrap(sector.Y)
	return fn(nil, nil)
}

func (sector *Sector) PyGet_z() (py.Object, error) {
	fn := pybind.Wrap(sector.Z)
	return fn(nil, nil)
}

func (region *Region) PyGet_origin() (py.Object, error) {
	fn := pybind.Wrap(region.Origin)
	return fn(nil, nil)
}

func (region *Region) Py_rebase(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(region.Rebase)
	return fn(args, kwds)
}

func (local *Local) PyGet_x() (py.Object, error) {
	fn := pybind.Wrap(local.X)
	return fn(nil, nil)
}

func (local *Local) PyGet_y() (py.Object, error) {
	fn := pybind.Wrap(local.Y)
	return fn(nil, nil)
}

func (local *Local) PyGet_z() (py.Object, error) {
	fn := pybind.Wrap(local.Z)
	return fn(nil, nil)
}

func (local *Local) PyGet_region() (py.Object, error) {
	fn := pybind.Wrap(local.Region)
	return fn(nil, nil)
}

func (local *Local) PyGet_absolute() (py.Object, error) {
	fn := pybind.Wrap(local.Absolute)
	return fn(nil, nil)
}

func (area *Area) PyGet_min_sector() (py.Object, error) {
	fn := pybind.Wrap(area.MinSector)
	return fn(nil, nil)
}

func (area *Area) PyGet_max_sector() (py.Object, error) {
	fn := pybind.Wrap(area.MaxSector)
	return fn(nil, nil)
}

func (area *Area) PyGet_min() (py.Object, error) {
	fn := pybind.Wrap(area.Min)
	return fn(nil, nil)
}

func (area *Area) PyGet_max() (py.Object, error) {
	fn := pybind.Wrap(area.Max)
	return fn(nil, nil)
}

func (area *Area) Py_contains(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(area.Contains)
	return fn(args, kwds)
}
