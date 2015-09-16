package gem

import (
	"os"

	"github.com/op/go-logging"
	"github.com/qur/gopy/lib"
)

//go:generate gopygen $GOFILE SysLog LogModule

var Logger *SysLog

type SysLog struct {
	py.BaseObject
}

type LogModule struct {
	py.BaseObject
	logger *logging.Logger
}

var format = logging.MustStringFormatter("%{color}[%{level:-8s}] %{module:-10s}%{color:reset}: %{message}")

func InitSysLog() (*SysLog, error) {
	b := logging.NewLogBackend(os.Stderr, "", 0)
	f := logging.NewBackendFormatter(b, format)
	logging.SetBackend(f)
	var err error
	Logger, err = SysLog{}.Alloc()
	return Logger, err
}

func (log *SysLog) Module(module string) *LogModule {
	logModule, err := LogModule{logger: logging.MustGetLogger(module)}.Alloc()
	if err != nil {
		panic(err)
	}

	return logModule
}

func (log *LogModule) Critical(message string) {
	log.logger.Critical(message)
}

func (log *LogModule) Debug(message string) {
	log.logger.Debug(message)
}

func (log *LogModule) Error(message string) {
	log.logger.Error(message)
}

func (log *LogModule) Fatal(message string) {
	log.logger.Fatal(message)
}

func (log *LogModule) Info(message string) {
	log.logger.Info(message)
}

func (log *LogModule) Notice(message string) {
	log.logger.Notice(message)
}

func (log *LogModule) Warning(message string) {
	log.logger.Warning(message)
}
