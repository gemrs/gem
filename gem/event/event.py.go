package event

import (
	"github.com/qur/gopy/lib"

	"github.com/gemrs/gem/pybind"
)

var EventDef = pybind.Define("Event", (*Event)(nil))
var RegisterEvent = pybind.GenerateRegisterFunc(EventDef)
var NewEvent = pybind.GenerateConstructor(EventDef).(func(string) *Event)

func (e *Event) PyGet_key() (py.Object, error) {
	fn := pybind.Wrap(e.Key)
	return fn(nil, nil)
}

func (e *Event) Py_register(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(e.Register)
	return fn(args, kwds)
}

func (e *Event) Py_unregister(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(e.Unregister)
	return fn(args, kwds)
}

// Py_notify_observers is a manual python wrapper of notify_observers, because
// pybind doesn't support ellipsis args
func (e *Event) Py_notify_observers(argsTuple *py.Tuple) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	args := []interface{}{}
	if argsTuple.Size() > 1 {
		for _, a := range argsTuple.Slice() {
			a.Incref()
			argIn, err := pybind.TypeConvIn(a, "")
			if err != nil {
				py.None.Incref()
				return py.None, nil
			}

			args = append(args, argIn)
		}
	}

	e.Incref()
	e.NotifyObservers(args...)
	e.Decref()

	for _, a := range argsTuple.Slice() {
		a.Decref()
	}

	py.None.Incref()
	return py.None, nil
}

func (e *Event) PyRepr() string {
	return e.Key()
}
