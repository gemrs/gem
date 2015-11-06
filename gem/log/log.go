package log

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/op/go-logging"
	"github.com/qur/gopy/lib"
)

var Sys *SysLog

type SysLog struct {
	py.BaseObject
	redirectBuffer *bytes.Buffer
	modules        map[string]*Module
}

func (s *SysLog) Init() {}

type Module struct {
	py.BaseObject
	logger *logging.Logger
	parent *Module
	prefix string
}

func (m *Module) Init() {}

var format = logging.MustStringFormatter("%{color}[%{level:-8s}] %{module:-10s}%{color:reset}: %{message}")

func New(module string) *Module {
	return Sys.Module(module)
}

func InitSysLog() {
	SetBackend(os.Stdout)
	Sys = NewSysLog()
}

func SetBackend(out io.Writer) {
	b := logging.NewLogBackend(out, "", 0)
	f := logging.NewBackendFormatter(b, format)
	logging.SetBackend(f)
}

func (log *SysLog) BeginRedirect() {
	log.redirectBuffer = bytes.NewBuffer([]byte{})
	SetBackend(log.redirectBuffer)
}

func (log *SysLog) EndRedirect() {
	SetBackend(os.Stdout)
	_, _ = log.redirectBuffer.WriteTo(os.Stdout)
}

func (log *SysLog) Module(module string) *Module {
	if log.modules == nil {
		log.modules = make(map[string]*Module)
	}

	if logModule, ok := log.modules[module]; ok {
		return logModule
	}

	logModule := NewModule()
	logModule.logger = logging.MustGetLogger(module)

	log.modules[module] = logModule

	return logModule
}

func (log *Module) SubModule(prefix string) (out Logger) {
	logModule := NewModule()

	logModule.parent = log
	logModule.prefix = prefix

	return logModule
}

func (log *Module) Critical(message string) {
	if log.parent != nil {
		log.parent.Critical(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.logger.Critical(message)
}

func (log *Module) Debug(message string) {
	if log.parent != nil {
		log.parent.Debug(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.logger.Debug(message)
}

func (log *Module) Error(message string) {
	if log.parent != nil {
		log.parent.Error(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.logger.Error(message)
}

func (log *Module) Fatal(message string) {
	if log.parent != nil {
		log.parent.Fatal(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.logger.Fatal(message)
}

func (log *Module) Info(message string) {
	if log.parent != nil {
		log.parent.Info(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.logger.Info(message)
}

func (log *Module) Info(message string) {
	if log.parent != nil {
		log.parent.Info(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.logger.Info(message)
}

func (log *Module) Error(message string) {
	if log.parent != nil {
		log.parent.Error(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.logger.Error(message)
}

func (log *Module) Error(format string, a ...interface{}) {
	log.Critical(fmt.Sprintf(format, a...))
}

func (log *Module) Debug(format string, a ...interface{}) {
	log.Debug(fmt.Sprintf(format, a...))
}

func (log *Module) Error(format string, a ...interface{}) {
	log.Error(fmt.Sprintf(format, a...))
}

func (log *Module) Fatalf(format string, a ...interface{}) {
	log.Fatal(fmt.Sprintf(format, a...))
}

func (log *Module) Info(format string, a ...interface{}) {
	log.Info(fmt.Sprintf(format, a...))
}

func (log *Module) Info(format string, a ...interface{}) {
	log.Info(fmt.Sprintf(format, a...))
}

func (log *Module) Error(format string, a ...interface{}) {
	log.Error(fmt.Sprintf(format, a...))
}
