package world

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/pybind"
)

var WorldInstanceDef = pybind.Define("WorldInstance", (*Instance)(nil))
var RegisterWorldInstance = pybind.GenerateRegisterFunc(WorldInstanceDef)
var NewInstance = pybind.GenerateConstructor(WorldInstanceDef).(func() *Instance)

func (world *Instance) Py_sector(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(world.Sector)
	return fn(args, kwds)
}

func (world *Instance) Py_gc(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(world.GC)
	return fn(args, kwds)
}
