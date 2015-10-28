package task

import (
	"github.com/qur/gopy/lib"
)

func Py_submit(args *py.Tuple) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var taskHook string
	var callback py.Object
	var interval int
	var userdata py.Object
	err := py.ParseTuple(args, "OsiO", &callback, &taskHook, &interval, &userdata)
	if err != nil {
		return nil, err
	}

	task := PythonTask(callback, TaskHook(taskHook), Cycles(interval), userdata)
	Scheduler.Submit(task)
	py.None.Incref()
	return py.None, nil
}
