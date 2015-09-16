package task

import (
	"github.com/qur/gopy/lib"
)

type TaskCallback func(*Task) bool

type Cycles int

type Task struct {
	Callback TaskCallback
	When     TaskHook
	Interval Cycles
	User     interface{}
	counter  Cycles
}

func NewTask(callback TaskCallback, when TaskHook, interval Cycles, user interface{}) *Task {
	return &Task{
		Callback: callback,
		When:     when,
		Interval: interval,
		User:     user,
		counter:  interval,
	}
}

func (task *Task) Tick() bool {
	task.counter = task.counter - 1
	if task.counter == 0 {
		reschedule := task.Callback(task)
		if reschedule {
			task.counter = task.Interval
		}
	}

	return task.counter == 0
}

func PythonTask(callback py.Object, when TaskHook, interval Cycles, user py.Object) *Task {
	callback.Incref()
	user.Incref()
	cbFunc := func(task *Task) bool {
		argsTuple, err := py.BuildValue("sO", string(when), user)
		if err != nil {
			panic(err)
		}

		reschedule, err := callback.Base().CallObject(argsTuple.(*py.Tuple))
		if err != nil {
			panic(err)
		}

		return reschedule.(*py.Bool).Bool()
	}

	return NewTask(cbFunc, when, interval, user)
}
