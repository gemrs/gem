package gem

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/pybind"
)

var EngineDef = pybind.Define("Engine", (*Engine)(nil))
var RegisterEngine = pybind.GenerateRegisterFunc(EngineDef)
var NewEngine = pybind.GenerateConstructor(EngineDef).(func() *Engine)

func (e *Engine) Py_Start(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(e.Start)
	return fn(args, kwds)
}

func (e *Engine) Py_Join(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(e.Join)
	return fn(args, kwds)
}

func (e *Engine) Py_Stop(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(e.Stop)
	return fn(args, kwds)
}
