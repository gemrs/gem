//glua:bind module gem.engine
package engine

import (
	"time"

	"github.com/gemrs/gem/gem/core/task"
	engine_event "github.com/gemrs/gem/gem/engine/event"
	"github.com/gemrs/gem/gem/util/safe"
	"github.com/gemrs/willow/log"

	tomb "gopkg.in/tomb.v2"
)

var logger = log.New("engine", log.NilContext)

//go:generate glua .

//glua:bind
type Engine struct {
	t tomb.Tomb
}

var EngineTick = 600 * time.Millisecond

//glua:bind constructor Engine
func NewEngine() *Engine {
	return &Engine{}
}

//glua:bind
func (e *Engine) Start() {
	logger.Info("Starting engine")
	engine_event.Startup.NotifyObservers()

	e.t.Go(e.run)
}

//glua:bind
func (e *Engine) Join() {
	e.t.Wait()
}

//glua:bind
func (e *Engine) Stop() {
	engine_event.Shutdown.NotifyObservers()
	e.t.Kill(nil)
}

func (e *Engine) run() error {
	// Start the engine ticking...
	duringTask := task.NewTask(func(*task.Task) bool {
		engine_event.Tick.NotifyObservers()
		return true
	}, task.Tick, 1, nil)

	task.Scheduler.Submit(duringTask)

	// Main engine loop
	c := time.Tick(EngineTick)
	for _ = range c {
		if !e.t.Alive() {
			break
		}

		func() {
			defer safe.Recover(logger)

			task.Scheduler.Tick(task.Tick)
		}()
	}
	return nil
}
