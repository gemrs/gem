package event

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/pybind"
)

var EventDef = pybind.Define("Event", (*Event)(nil))
var RegisterEvent = pybind.GenerateRegisterFunc(EventDef)
var NewEvent = pybind.GenerateConstructor(EventDef).(func(string) *Event)

func (e *Event) Py_Key(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(e.Key)
	return fn(args, kwds)
}

func (e *Event) Py_Register(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(e.Register)
	return fn(args, kwds)
}

func (e *Event) Py_Unregister(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(e.Unregister)
	return fn(args, kwds)
}
