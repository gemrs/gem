package task

import (
	"time"
)

type TaskCallback func(Task) bool

type Task struct {
	Callback TaskCallback
	When     TaskHook
	Duration time.Duration
	User     interface{}
}

func NewTask(callback TaskCallback, when TaskHook, duration time.Duration, user interface{}) Task {
	return Task{
		Callback: callback,
		When:     when,
		Duration: duration,
		User:     user,
	}
}

func (task Task) Future(scheduler Scheduler) {
	go func() {
		select {
		case <-time.After(task.Duration):
			scheduler.taskQueues[task.When] <- task
		}
	}()
}
