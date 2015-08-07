package gem

import (
	"github.com/sinusoids/gem/gem/event"

	"github.com/qur/gopy/lib"
)

type registerFunc func(*py.Module) error

var moduleRegisterFuncs = []registerFunc{
	RegisterSysLog,
	RegisterEngine,
	RegisterLogModule,

	event.InitPyModule,
}

func InitPyModule() error {
	/* Create package */
	var err error
	var module *py.Module
	if module, err = py.InitModule("gem", []py.Method{}); err != nil {
		return err
	}

	/* Register modules */
	for _, registerFunc := range moduleRegisterFuncs {
		if err = registerFunc(module); err != nil {
			return err
		}
	}

	/* Create our logger object */
	if syslog, err := InitSysLog(); err != nil {
		return err
	} else if err = module.AddObject("syslog", syslog); err != nil {
		return err
	}

	return nil
}
