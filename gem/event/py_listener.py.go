package event

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/pybind"
)

var PyListenerDef = pybind.Define("PyListener", (*PyListener)(nil))
var RegisterPyListener = pybind.GenerateRegisterFunc(PyListenerDef)
var NewPyListener = pybind.GenerateConstructor(PyListenerDef).(func(py.Object) *PyListener)

func (l *PyListener) Py_Id(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(l.Id)
	return fn(args, kwds)
}
