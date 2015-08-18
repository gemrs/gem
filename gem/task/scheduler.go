package task

import (
	"container/list"
)

type _Scheduler struct {
	tasks map[TaskHook]*list.List
}

var Scheduler _Scheduler

func init() {
	Scheduler = NewScheduler()
}

func NewScheduler() _Scheduler {
	var s _Scheduler
	s.tasks = make(map[TaskHook]*list.List)
	for _, hook := range taskHookConstants {
		s.tasks[hook] = list.New()
	}
	return s
}

func (scheduler *_Scheduler) Submit(task *Task) {
	scheduler.tasks[task.When].PushBack(task)
}

func (scheduler *_Scheduler) Tick(hook TaskHook) {
	if queue, ok := scheduler.tasks[hook]; ok {
		for e := queue.Front(); e != nil; e = e.Next() {
			task := e.Value.(*Task)
			expire := task.Tick()
			if expire {
				queue.Remove(e)
			}
		}
	}
}
