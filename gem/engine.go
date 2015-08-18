package gem

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/event"
	"github.com/sinusoids/gem/gem/task"

	"time"
)

var logger *LogModule

//go:generate gopygen $GOFILE Engine
type Engine struct {
	py.BaseObject
}

var EngineTick = 600 * time.Millisecond

func (e *Engine) Start() {
	logger = Logger.Module("engine")
	logger.Info("Starting engine")
	event.Raise(event.Startup)

	// TODO(tom): this should be run in a seperate, interruptible goroutine
	e.run()
}

func (e *Engine) Stop() {
	event.Raise(event.Shutdown)
}

func (e *Engine) run() {
	preTask := task.NewTask(func(*task.Task) bool {
		event.Raise(event.PreTick)
		return true
	}, task.PreTick, 1, nil)

	duringTask := task.NewTask(func(*task.Task) bool {
		event.Raise(event.Tick)
		return true
	}, task.Tick, 1, nil)

	postTask := task.NewTask(func(*task.Task) bool {
		event.Raise(event.PostTick)
		return true
	}, task.PostTick, 1, nil)

	task.Scheduler.Submit(preTask)
	task.Scheduler.Submit(duringTask)
	task.Scheduler.Submit(postTask)

	c := time.Tick(EngineTick)
	for _ = range c {
		task.Scheduler.Tick(task.PreTick)
		task.Scheduler.Tick(task.Tick)
		task.Scheduler.Tick(task.PostTick)
	}
}
