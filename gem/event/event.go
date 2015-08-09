package event

import (
	"github.com/qur/gopy/lib"
)

// A unique, hookable identifier for a certain event
type Event string

// There should be a nice way of doing this. Perhaps a generator or something.
const (
	TestEvent1 Event = "TestEvent1"
	TestEvent2 Event = "TestEvent2"
	TestEvent3 Event = "TestEvent3"
	TestEvent4 Event = "TestEvent4"
	Startup    Event = "Startup"
	Shutdown   Event = "Shutdown"
)

var pyEventConstants = []Event{
	TestEvent1,
	TestEvent2,
	TestEvent3,
	TestEvent4,
	Startup,
	Shutdown,
}

func createEventConstants(module *py.Module) {
	for _, event := range pyEventConstants {
		if pyString, err := py.NewString((string)(event)); err != nil {
			panic(err)
		} else if err := module.AddObject((string)(event), pyString); err != nil {
			panic(err)
		}
	}
}
