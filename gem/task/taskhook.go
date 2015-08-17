package task

import (
	"github.com/qur/gopy/lib"
)

// A unique, hookable identifier for a certain event
type TaskHook string

// There should be a nice way of doing this. Perhaps a generator or something.
const (
	PreTick  TaskHook = "PreTick"
	Tick     TaskHook = "Tick"
	PostTick TaskHook = "PostTick"
)

var taskHookConstants = []TaskHook{
	PreTick,
	Tick,
	PostTick,
}

func createTaskHookConstants(module *py.Module) {
	for _, event := range taskHookConstants {
		if pyString, err := py.NewString((string)(event)); err != nil {
			panic(err)
		} else if err := module.AddObject((string)(event), pyString); err != nil {
			panic(err)
		}
	}
}
