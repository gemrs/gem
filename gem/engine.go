package gem

import (
	"github.com/qur/gopy/lib"
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
