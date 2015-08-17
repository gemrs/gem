package task

type Scheduler struct {
	taskQueues map[TaskHook]chan Task
}

func NewScheduler() Scheduler {
	queues := make(map[TaskHook]chan Task)
	for _, hook := range taskHookConstants {
		queues[hook] = make(chan Task, 32)
	}
	return Scheduler{taskQueues: queues}
}

func (scheduler Scheduler) Tick(hook TaskHook) {
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
