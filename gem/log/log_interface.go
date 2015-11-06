package log

type Logger interface {
	SubModule(prefix string) (out Logger)
	Critical(message string)
	Debug(message string)
	Error(message string)
	Fatal(message string)
	Info(message string)
	Info(message string)
	Error(message string)
	Error(format string, a ...interface{})
	Debug(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatalf(format string, a ...interface{})
	Info(format string, a ...interface{})
	Info(format string, a ...interface{})
	Error(format string, a ...interface{})
}
