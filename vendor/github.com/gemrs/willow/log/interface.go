package log

// Level is a string identifier for a logging level
type Level string

const (
	LvlInfo   Level = "INFO"
	LvlError        = "ERROR"
	LvlDebug        = "DEBUG"
	LvlNotice       = "NOTICE"
)

// Context is a provider of context for a Record
type Context interface {
	ContextMap() map[string]interface{}
}

// A Record is a single entry in the log
type Record interface {
	Level() Level
	Tag() string
	Message() string
	Context() Context
}

// Handler is a target for Records
type Handler interface {
	Handle(Record)
}

// Logger provides convenience functions for dispatching messages at standard log levels
type Logger interface {
	Info(string, ...interface{})
	Error(string, ...interface{})
	Debug(string, ...interface{})
	Notice(string, ...interface{})
}

// A Dispatcher constructs records and pushes them to one or more Handlers
type Dispatcher interface {
	Dispatch(lvl Level, message string)
}

// A Log is both a Dispatcher and a Logger
type Log interface {
	Logger
	Dispatcher

	Child(tag string, ctx Context) Log
}
