package log

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/op/go-logging"
	"github.com/qur/gopy/lib"
)

//go:generate gopygen -type SysLog -type Module -exclude "(Critical|Debug|Error|Fatal|Info|Notice|Warning)f" $GOFILE

var Sys *SysLog

type SysLog struct {
	py.BaseObject
	redirectBuffer *bytes.Buffer
}

type Module struct {
	py.BaseObject
	logger *logging.Logger
	parent *Module
	prefix string
}

var format = logging.MustStringFormatter("%{color}[%{level:-8s}] %{module:-10s}%{color:reset}: %{message}")

func New(module string) *Module {
	return Sys.Module(module)
}

func InitSysLog() error {
	var err error
	SetBackend(os.Stdout)
	Sys, err = SysLog{}.Alloc()
	return err
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
	logModule, err := Module{logger: logging.MustGetLogger(module)}.Alloc()
	if err != nil {
		panic(err)
	}

	return logModule
}

func (log *Module) SubModule(prefix string) *Module {
	logModule, err := Module{
		parent: log,
		prefix: prefix,
	}.Alloc()
	if err != nil {
		panic(err)
	}

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
		log.parent.Critical(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.logger.Error(message)
}

func (log *Module) Fatal(message string) {
	if log.parent != nil {
		log.parent.Critical(fmt.Sprintf("[%v] %v", log.prefix, message))
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

func (log *Module) Notice(message string) {
	if log.parent != nil {
		log.parent.Notice(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.logger.Notice(message)
}

func (log *Module) Warning(message string) {
	if log.parent != nil {
		log.parent.Warning(fmt.Sprintf("[%v] %v", log.prefix, message))
		return
	}
	log.logger.Warning(message)
}

func (log *Module) Criticalf(format string, a ...interface{}) {
	log.Critical(fmt.Sprintf(format, a...))
}

func (log *Module) Debugf(format string, a ...interface{}) {
	log.Debug(fmt.Sprintf(format, a...))
}

func (log *Module) Errorf(format string, a ...interface{}) {
	log.Error(fmt.Sprintf(format, a...))
}

func (log *Module) Fatalf(format string, a ...interface{}) {
	log.Fatal(fmt.Sprintf(format, a...))
}

func (log *Module) Infof(format string, a ...interface{}) {
	log.Info(fmt.Sprintf(format, a...))
}

func (log *Module) Noticef(format string, a ...interface{}) {
	log.Notice(fmt.Sprintf(format, a...))
}

func (log *Module) Warningf(format string, a ...interface{}) {
	log.Warning(fmt.Sprintf(format, a...))
}
