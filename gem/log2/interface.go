package log

type Level string

const (
	LvlInfo  Level = "INFO"
	LvlError       = "ERROR"
	LvlDebug       = "DEBUG"
)

type Context interface {
	ContextMap() map[string]interface{}
}

type Record interface {
	Level() Level
	Tag() string
	Message() string
	Context() Context
}

type Handler interface {
	Handle(Record)
}

type Logger interface {
	Info(string, ...interface{})
	Error(string, ...interface{})
	Debug(string, ...interface{})
}

type Dispatcher interface {
	Dispatch(lvl Level, message string)
}

type Log interface {
	Logger
	Dispatcher
}
