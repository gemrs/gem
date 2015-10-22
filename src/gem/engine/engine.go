package gem

import (
	"time"

	engine_event "gem/engine/event"
	"gem/log"
	"gem/task"

	"github.com/qur/gopy/lib"
	tomb "gopkg.in/tomb.v2"
)

var logger *log.Module

//go:generate gopygen -type Engine -excfunc "^[a-z].+" $GOFILE
type Engine struct {
	py.BaseObject

	t tomb.Tomb
}

var EngineTick = 600 * time.Millisecond

func (e *Engine) Init() error {
	return nil
}

func (e *Engine) Start() {
	logger = log.New("engine")
	logger.Info("Starting engine")
	engine_event.Startup.NotifyObservers()

	e.t.Go(e.run)
}

func (e *Engine) Join() bool {
	e.t.Wait()
	return true
}

func (e *Engine) Stop() {
	engine_event.Shutdown.NotifyObservers()
	e.t.Kill(nil)
	e.t.Wait()
}

func (e *Engine) run() error {
	// Start the engine ticking...
	preTask := task.NewTask(func(*task.Task) bool {
		engine_event.PreTick.NotifyObservers()
		return true
	}, task.PreTick, 1, nil)

	duringTask := task.NewTask(func(*task.Task) bool {
		engine_event.Tick.NotifyObservers()
		return true
	}, task.Tick, 1, nil)

	postTask := task.NewTask(func(*task.Task) bool {
		engine_event.PostTick.NotifyObservers()
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
