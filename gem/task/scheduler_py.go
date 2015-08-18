package task

import (
	"time"

	"github.com/qur/gopy/lib"
	_ "github.com/tgascoigne/gopygen/gopygen"
)

func Py_Submit(args *py.Tuple) (py.Object, error) {
	var taskHook string
	var callback py.Object
	var interval int
	var userdata py.Object
	err := py.ParseTuple(args, "OsiO", &callback, &taskHook, &interval, &userdata)
	if err != nil {
		return nil, err
	}

	task := PythonTask(callback, TaskHook(taskHook), time.Duration(interval)*time.Millisecond, userdata)
	Scheduler.Submit(task)
	py.None.Incref()
	return py.None, nil
}

/*
func (scheduler *Scheduler) Py_Tick(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_Tick: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "TaskHook")
	if err != nil {
		return nil, err
	}

	scheduler.Tick(in_0.(TaskHook))

	py.None.Incref()
	return py.None, nil

}
*/
