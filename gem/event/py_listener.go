package event

import (
	"github.com/sinusoids/gem/gem/log"
	"github.com/sinusoids/gem/gem/util/safe"

	"github.com/qur/gopy/lib"
	"github.com/tgascoigne/gopygen/gopygen"
)

//go:generate gopygen -type PyListener -excfunc "Notify" $GOFILE

type PyListener struct {
	py.BaseObject

	id     int
	fn     py.Object
	logger *log.Module
}

func (l *PyListener) Init(fn py.Object) error {
	l.id = <-nextId
	fn.Incref()
	l.fn = fn
	l.logger = log.New("python_listener")
	return nil
}

func (l *PyListener) Id() int {
	return l.id
}

func (l *PyListener) Notify(e *Event, args ...interface{}) {
	lock := py.NewLock()
	defer lock.Unlock()

	argsIn := []interface{}{e}
	argsIn = append(argsIn, args...)

	defer safe.Recover(l.logger)

	argsOut := []py.Object{}
	for _, a := range argsIn {
		converted, err := gopygen.TypeConvOut(a, "")
		if err != nil {
			panic(err)
		}
		converted.Incref()
		argsOut = append(argsOut, converted)
	}

	_, err := l.fn.Base().CallFunctionObjArgs(argsOut...)
	if err != nil {
		// Panicing with the whole py error object causes a double panic.
		// Suspect this is because python has cleaned up by the time the runtime evals the error
		panic(err.Error())
	}
}
