//glua:bind module gem.engine.event
package engine_event

import (
	"github.com/gemrs/gem/gem/core/event"
)

//go:generate glua .

//glua:bind
var (
	Startup  = event.NewEvent("Startup")
	Shutdown = event.NewEvent("Shutdown")
	PreTick  = event.NewEvent("PreTick")
	Tick     = event.NewEvent("Tick")
	PostTick = event.NewEvent("PostTick")
)
