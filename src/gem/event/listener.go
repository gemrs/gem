package event

import (
	"github.com/qur/gopy/lib"
	"github.com/tgascoigne/gopygen/gopygen"
)

type Listener func(Event, ...interface{})

type Listeners []Listener

func (l Listeners) Dispatch(event Event, args ...interface{}) {
	for _, listener := range l {
		listener(event, args...)
	}
}

func PythonListener(callback py.Object) Listener {
	callback.Incref()
	return func(ev Event, args ...interface{}) {
		argsIn := []interface{}{string(ev)}
		argsIn = append(argsIn, args...)

		argsOut := []py.Object{}
		for _, a := range argsIn {
			converted, err := gopygen.TypeConvOut(a, "")
			if err != nil {
				panic(err)
			}
			converted.Incref()
			argsOut = append(argsOut, converted)
		}

		lock := py.NewLock()
		_, err := callback.Base().CallFunctionObjArgs(argsOut...)
		if err != nil {
			panic(err)
		}
		lock.Unlock()
	}
}
