package event

import (
	"github.com/qur/gopy/lib"
	"github.com/sinusoids/gem/gem/python"
)

type Event struct {
	py.BaseObject

	key       string
	observers map[int]Observer
}

func (e *Event) Init(key string) {
	e.key = key
	e.observers = make(map[int]Observer)
}

func (e *Event) Key() string {
	return e.key
}

func (e *Event) Register(o Observer) {
	e.observers[o.Id()] = o
}

func (e *Event) Unregister(o Observer) {
	delete(e.observers, o.Id())
}

func (e *Event) NotifyObservers(args ...interface{}) {
	for _, observer := range e.observers {
		observer.Notify(e, args...)
	}
}

// Py_NotifyObservers is a manual python wrapper of NotifyObservers, because
// pybind doesn't support ellipsis args
func (e *Event) Py_NotifyObservers(argsTuple *py.Tuple) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	args := []interface{}{}
	if argsTuple.Size() > 1 {
		for _, a := range argsTuple.Slice() {
			a.Incref()
			argIn, err := python.TypeConvIn(a, "")
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
