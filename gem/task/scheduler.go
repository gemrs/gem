package task

type _Scheduler struct {
	taskQueues map[TaskHook]chan Task
}

var Scheduler _Scheduler

func init() {
	queues := make(map[TaskHook]chan Task)
	for _, hook := range taskHookConstants {
		queues[hook] = make(chan Task, 32)
	}
	Scheduler = _Scheduler{taskQueues: queues}
}

func (scheduler *_Scheduler) Submit(task Task) {
	task.Future(scheduler)
}

func (scheduler *_Scheduler) Tick(hook TaskHook) {
	if queue, ok := scheduler.taskQueues[hook]; ok {
		empty := false
		for !empty {
			select {
			case task := <-queue:
				reschedule := task.Callback(task)
				if reschedule {
					task.Future(scheduler)
				}
			default:
				empty = true
			}
		}
	}
}
