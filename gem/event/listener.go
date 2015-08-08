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
	return func(ev Event, args ...interface{}) {
		argsIn := []interface{}{string(ev)}
		argsIn = append(argsIn, args...)

		argsOut := []py.Object{}
		for _, a := range argsIn {
			converted, err := gopygen.TypeConvOut(a, "")
			if err != nil {
				panic(err)
			}
			argsOut = append(argsOut, converted)
		}

		argTuple, err := py.PackTuple(argsOut...)
		if err != nil {
			panic(err)
		}

		_, err = callback.Base().CallObject(argTuple)
		if err != nil {
			panic(err)
		}
	}
}
