package event

import (
	"github.com/qur/gopy/lib"
)

func Py_RegisterListener(args *py.Tuple) (py.Object, error) {
	var event string
	var callback py.Object
	err := py.ParseTuple(args, "sO", &event, &callback)
	if err != nil {
		return nil, err
	}

	listener := PythonListener(callback)
	Dispatcher.Register(Event(event), listener)
	py.None.Incref()
	return py.None, nil
}

func Py_RaiseEvent(args *py.Tuple) (py.Object, error) {
	var event_ string
	err := py.ParseTuple(args, "s", &event_)
	if err != nil {
		return nil, err
	}

	Dispatcher.Raise(Event(event_))
	py.None.Incref()
	return py.None, nil
}

func Py_Clear(args *py.Tuple) (py.Object, error) {
	Dispatcher.Clear()
	py.None.Incref()
	return py.None, nil
}
