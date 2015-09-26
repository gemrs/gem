package event

import (
	"github.com/qur/gopy/lib"
	"github.com/tgascoigne/gopygen/gopygen"
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

func Py_RaiseEvent(argsTuple *py.Tuple) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	temp, err := argsTuple.GetItem(0)
	if err != nil {
		py.None.Incref()
		return py.None, nil
	}
	event := Event(temp.(*py.String).String())

	args := []interface{}{}

	if argsTuple.Size() > 1 {
		for _, a := range argsTuple.Slice()[1:] {
			argIn, err := gopygen.TypeConvIn(a, "")
			if err != nil {
				py.None.Incref()
				return py.None, nil
			}

			args = append(args, argIn)
		}
	}

	Dispatcher.Raise(event, args...)
	py.None.Incref()
	return py.None, nil
}

func Py_Clear(args *py.Tuple) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	Dispatcher.Clear()
	py.None.Incref()
	return py.None, nil
}
