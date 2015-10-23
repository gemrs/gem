package gem

import (
	"github.com/qur/gopy/lib"

	"gem/log"
	"gem/python"
)

type registerFunc func(*py.Module) error

var moduleRegisterFuncs = []registerFunc{
	log.RegisterSysLog,
	log.RegisterModule,
}

func init() {
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
