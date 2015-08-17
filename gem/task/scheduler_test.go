package task

import (
	"testing"
	"time"
)

func TestFuture(t *testing.T) {
	scheduler := NewScheduler()

	taskChan := make(chan string, 5)
	callback := func(task Task) bool {
		t.Logf("Got task %v", task.User)
		taskChan <- task.User.(string)
		return false
	}

	tasks := []Task{
		NewTask(callback, PreTick, time.Second*2, "Task1"),
		NewTask(callback, Tick, time.Second*3, "Task2"),
	}
	count := len(tasks)

	for _, t := range tasks {
		t.Future(scheduler)
	}

	done := false
	timeout := time.After(time.Second * 5)
	for !done {
		scheduler.Tick(PreTick)
		scheduler.Tick(Tick)
		scheduler.Tick(PostTick)

		select {
		case _ = <-taskChan:
			count = count - 1
		case <-timeout:
			done = true
		default:
		}

		if count == 0 {
			done = true
		}
	}

	if count > 0 {
		t.Errorf("Didn't schedule all tasks")
	}
}

func TestReschedule(t *testing.T) {
	scheduler := NewScheduler()

	taskChan := make(chan string, 5)
	count := 5
	callback := func(task Task) bool {
		t.Logf("Got task %v", task.User)
		taskChan <- task.User.(string)
		count = count - 1
		return count > 0
	}

	tasks := []Task{
		NewTask(callback, PreTick, time.Second*2, "Task1"),
	}

	for _, t := range tasks {
		t.Future(scheduler)
	}

	done := false
	timeout := time.After(time.Second * 11)
	newCount := count
	for !done {
		scheduler.Tick(PreTick)
		scheduler.Tick(Tick)
		scheduler.Tick(PostTick)

		select {
		case _ = <-taskChan:
			newCount = newCount - 1
		case <-timeout:
			done = true
		default:
		}

		if newCount == 0 {
			done = true
		}
	}

	if newCount > 0 {
		t.Errorf("Didn't schedule all tasks")
	}
}
