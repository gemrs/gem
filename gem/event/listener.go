package event

import (
	"github.com/qur/gopy/lib"
)

type Listener func(Event)

type Listeners []Listener

func (l Listeners) Dispatch(event Event) {
	for _, listener := range l {
		listener(event)
	}
}

func PythonListener(callback py.Object) Listener {
	return func(ev Event) {
		eventArg, err := py.NewString(string(ev))
		if err != nil {
			panic(err)
		}

		argTuple, err := py.PackTuple(eventArg)
		if err != nil {
			panic(err)
		}

		_, err = callback.Base().CallObject(argTuple)
		if err != nil {
			panic(err)
		}
	}
}
