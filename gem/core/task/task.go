package task

import (
	"github.com/gemrs/gem/gem/util/safe"
	"github.com/gemrs/willow/log"
)

type TaskCallback func(*Task) bool

type Cycles int

type Task struct {
	Callback TaskCallback
	When     TaskHook
	Interval Cycles
	User     interface{}
	counter  Cycles
	logger   log.Log
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
	defer safe.Recover(task.logger)

	task.counter = task.counter - 1
	if task.counter == 0 {
		reschedule := task.Callback(task)
		if reschedule {
			task.counter = task.Interval
		}
	}

	return task.counter == 0
}
