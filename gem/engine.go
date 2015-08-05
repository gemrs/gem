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
}

func (e *Engine) TestRegister(callback py.Object) {
	event.Dispatcher.Register(event.Event("TEST_EVENT"), event.PythonListener(callback))
}

func (e *Engine) TestRaise() {
	event.Dispatcher.Raise(event.Event("TEST_EVENT"))
}
