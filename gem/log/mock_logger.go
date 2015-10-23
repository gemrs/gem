package log

import (
	"fmt"
	"strings"
)

type Mock struct {
	Output []string
	parent *Mock
	prefix string
}

func NewMock(prefix string) *Mock {
	return &Mock{
		Output: make([]string, 0),
		prefix: prefix,
		parent: nil,
	}
}

func (log *Mock) HasLogged(message string) bool {
	for _, s := range log.Output {
		if strings.Contains(s, message) {
			return true
		}
	}
	return false
}

const (
	Critical string = "CRITICAL"
	Debug           = "DEBUG"
	Error           = "ERROR"
	Fatal           = "FATAL"
	Info            = "INFO"
	Notice          = "NOTICE"
	Warning         = "WARNING"
)

func (log *Mock) SubModule(prefix string) Logger {
	out := NewMock(prefix)
	out.parent = log
	return out
}

func (log *Mock) doLog(tag, message string) {
	log.Output = append(log.Output, fmt.Sprintf("%v %v", tag, message))
}

func (log *Mock) Critical(message string) {
	if log.parent != nil {
		log.parent.Critical(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.doLog(Critical, message)
}

func (log *Mock) Debug(message string) {
	if log.parent != nil {
		log.parent.Debug(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.doLog(Debug, message)
}

func (log *Mock) Error(message string) {
	if log.parent != nil {
		log.parent.Critical(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.doLog(Error, message)
}

func (log *Mock) Fatal(message string) {
	if log.parent != nil {
		log.parent.Critical(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.doLog(Fatal, message)
}

func (log *Mock) Info(message string) {
	if log.parent != nil {
		log.parent.Info(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.doLog(Info, message)
}

func (log *Mock) Notice(message string) {
	if log.parent != nil {
		log.parent.Notice(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.doLog(Notice, message)
}

func (log *Mock) Warning(message string) {
	if log.parent != nil {
		log.parent.Warning(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.doLog(Warning, message)
}

func (log *Mock) Criticalf(format string, a ...interface{}) {
	log.Critical(fmt.Sprintf(format, a...))
}

func (log *Mock) Debugf(format string, a ...interface{}) {
	log.Debug(fmt.Sprintf(format, a...))
}

func (log *Mock) Errorf(format string, a ...interface{}) {
	log.Error(fmt.Sprintf(format, a...))
}

func (log *Mock) Fatalf(format string, a ...interface{}) {
	log.Fatal(fmt.Sprintf(format, a...))
}

func (log *Mock) Infof(format string, a ...interface{}) {
	log.Info(fmt.Sprintf(format, a...))
}

func (log *Mock) Noticef(format string, a ...interface{}) {
	log.Notice(fmt.Sprintf(format, a...))
}

func (log *Mock) Warningf(format string, a ...interface{}) {
	log.Warning(fmt.Sprintf(format, a...))
}
