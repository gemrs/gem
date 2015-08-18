package task

import (
	"time"

	"github.com/qur/gopy/lib"
)

type TaskCallback func(Task) bool

type Task struct {
	Callback TaskCallback
	When     TaskHook
	Interval time.Duration
	User     interface{}
}

func NewTask(callback TaskCallback, when TaskHook, interval time.Duration, user interface{}) Task {
	return Task{
		Callback: callback,
		When:     when,
		Interval: interval,
		User:     user,
	}
}

func (task Task) Future(scheduler *_Scheduler) {
	go func() {
		select {
		case <-time.After(task.Interval):
			scheduler.taskQueues[task.When] <- task
		}
	}()
}

func PythonTask(callback py.Object, when TaskHook, interval time.Duration, user py.Object) Task {
	callback.Incref()
	user.Incref()
	cbFunc := func(task Task) bool {
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
