package log

import (
	"os"
	"fmt"

	"github.com/op/go-logging"
	"github.com/qur/gopy/lib"
)

//go:generate gopygen -type SysLog -type Module -exclude "(Critical|Debug|Error|Fatal|Info|Notice|Warning)f" $GOFILE

var Sys *SysLog

type SysLog struct {
	py.BaseObject
}

type Module struct {
	py.BaseObject
	logger *logging.Logger
}

var format = logging.MustStringFormatter("%{color}[%{level:-8s}] %{module:-10s}%{color:reset}: %{message}")

func New(module string) *Module {
	return Sys.Module(module)
}

func InitSysLog() error {
	b := logging.NewLogBackend(os.Stderr, "", 0)
	f := logging.NewBackendFormatter(b, format)
	logging.SetBackend(f)
	var err error
	Sys, err = SysLog{}.Alloc()
	return err
}

func (log *SysLog) Module(module string) *Module {
	logModule, err := Module{logger: logging.MustGetLogger(module)}.Alloc()
	if err != nil {
		panic(err)
	}

	return logModule
}

func (log *Module) Critical(message string) {
	log.logger.Critical(message)
}

func (log *Module) Criticalf(format string, a ...interface{}) {
	log.logger.Critical(fmt.Sprintf(format, a...))
}

func (log *Module) Debug(message string) {
	log.logger.Debug(message)
}

func (log *Module) Debugf(format string, a ...interface{}) {
	log.logger.Debug(fmt.Sprintf(format, a...))
}

func (log *Module) Error(message string) {
	log.logger.Error(message)
}

func (log *Module) Errorf(format string, a ...interface{}) {
	log.logger.Error(fmt.Sprintf(format, a...))
}

func (log *Module) Fatal(message string) {
	log.logger.Fatal(message)
}

func (log *Module) Fatalf(format string, a ...interface{}) {
	log.logger.Fatal(fmt.Sprintf(format, a...))
}

func (log *Module) Info(message string) {
	log.logger.Info(message)
}

func (log *Module) Infof(format string, a ...interface{}) {
	log.logger.Info(fmt.Sprintf(format, a...))
}

func (log *Module) Notice(message string) {
	log.logger.Notice(message)
}

func (log *Module) Noticef(format string, a ...interface{}) {
	log.logger.Notice(fmt.Sprintf(format, a...))
}

func (log *Module) Warning(message string) {
	log.logger.Warning(message)
}

func (log *Module) Warningf(format string, a ...interface{}) {
	log.logger.Warning(fmt.Sprintf(format, a...))
}
