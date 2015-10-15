package event

import (
	"gem/log"
	"gem/safe"

	"github.com/qur/gopy/lib"
	"github.com/tgascoigne/gopygen/gopygen"
)

var logger *log.Module

type Listener func(Event, ...interface{})

type Listeners []Listener

func (l Listeners) Dispatch(event Event, args ...interface{}) {
	for _, listener := range l {
		listener(event, args...)
	}
}

func PythonListener(callback py.Object) Listener {
	lock := py.NewLock()
	defer lock.Unlock()

	callback.Incref()

	if logger == nil {
		logger = log.New("python_event")
	}

	return func(ev Event, args ...interface{}) {
		argsIn := []interface{}{string(ev)}
		argsIn = append(argsIn, args...)

		lock := py.NewLock()
		defer lock.Unlock()
		defer safe.Recover(logger)

		argsOut := []py.Object{}
		for _, a := range argsIn {
			converted, err := gopygen.TypeConvOut(a, "")
			if err != nil {
				panic(err)
			}
			converted.Incref()
			argsOut = append(argsOut, converted)
		}

		_, err := callback.Base().CallFunctionObjArgs(argsOut...)
		if err != nil {
			// Panicing with the whole py error object causes a double panic.
			// Suspect this is because python has cleaned up by the time the runtime evals the error
			panic(err.Error())
		}
	}
}
