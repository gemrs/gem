package task

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
