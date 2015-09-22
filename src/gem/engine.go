package gem

import (
	"time"

	"gem/log"
	"gem/event"
	"gem/task"

	tomb "gopkg.in/tomb.v2"
	"github.com/qur/gopy/lib"
)

var logger *log.Module

//go:generate gopygen -type Engine -exclude "^[a-z].+" $GOFILE
type Engine struct {
	py.BaseObject

	t tomb.Tomb
}

var EngineTick = 600 * time.Millisecond

func (e *Engine) Start() {
	logger = log.New("engine")
	logger.Info("Starting engine")
	event.Raise(event.Startup)

	e.t.Go(e.run)
}

func (e *Engine) Join() {
	lock := py.NewLock()
	defer lock.Unlock()
	lock.UnblockThreads()
	e.t.Wait()
}

func (e *Engine) Stop() {
	event.Raise(event.Shutdown)
	e.t.Kill(nil)
	e.t.Wait()
}

func (e *Engine) run() error {
	// Start the engine ticking...
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

	// Main engine loop
	c := time.Tick(EngineTick)
	for _ = range c {
		if !e.t.Alive() {
			break
		}

		task.Scheduler.Tick(task.PreTick)
		task.Scheduler.Tick(task.Tick)
		task.Scheduler.Tick(task.PostTick)
	}
	return nil
}
