package log

type Logger interface {
	SubModule(prefix string) (out Logger)
	Critical(message string)
	Debug(message string)
	Error(message string)
	Fatal(message string)
	Info(message string)
	Notice(message string)
	Warning(message string)
	Criticalf(format string, a ...interface{})
	Debugf(format string, a ...interface{})
	Errorf(format string, a ...interface{})
	Fatalf(format string, a ...interface{})
	Infof(format string, a ...interface{})
	Noticef(format string, a ...interface{})
	Warningf(format string, a ...interface{})
}
