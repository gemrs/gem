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
	scheduler task.Scheduler
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
	e.scheduler = task.NewScheduler()

	preTask := task.NewTask(func(task.Task) bool {
		event.Raise(event.PreTick)
		return true
	}, task.PreTick, EngineTick, nil)

	duringTask := task.NewTask(func(task.Task) bool {
		event.Raise(event.Tick)
		return true
	}, task.Tick, EngineTick, nil)

	postTask := task.NewTask(func(task.Task) bool {
		event.Raise(event.PostTick)
		return true
	}, task.PostTick, EngineTick, nil)

	preTask.Future(e.scheduler)
	duringTask.Future(e.scheduler)
	postTask.Future(e.scheduler)

	c := time.Tick(EngineTick)
	for _ = range c {
		e.scheduler.Tick(task.PreTick)
		e.scheduler.Tick(task.Tick)
		e.scheduler.Tick(task.PostTick)
	}
}
