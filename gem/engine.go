package gem

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/event"
)

var logger *LogModule

//go:generate gopygen $GOFILE Engine
type Engine struct {
	py.BaseObject
}

func (e *Engine) Start() {
	logger = Logger.Module("engine")
	logger.Info("Starting engine")
	event.Raise(event.Startup)
}

func (e *Engine) Stop() {
	event.Raise(event.Shutdown)
}
