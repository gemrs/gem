package gem

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/log"
	"github.com/sinusoids/gem/gem/python"
)

type registerFunc func(*py.Module) error

var moduleRegisterFuncs = []registerFunc{
	log.RegisterSysLog,
	log.RegisterModule,
}

func init() {
	lock := py.NewLock()
	defer lock.Unlock()

	/* Create package */
	var err error
	var module *py.Module
	if module, err = python.InitModule("gem", []py.Method{}); err != nil {
		panic(err)
	}

	/* Register modules */
	for _, registerFunc := range moduleRegisterFuncs {
		if err = registerFunc(module); err != nil {
			panic(err)
		}
	}

	/* Create our logger object */
	log.InitSysLog()
	if err := module.AddObject("syslog", log.Sys); err != nil {
		panic(err)
	}
}
