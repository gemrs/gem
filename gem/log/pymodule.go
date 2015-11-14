package log

import (
	"os"

	"github.com/qur/gopy/lib"

	"github.com/gemrs/gem/gem/python/modules"
	"github.com/gemrs/willow/log"
)

type registerFunc func(*py.Module) error

var moduleRegisterFuncs = []registerFunc{
	RegisterPyContext,
	RegisterPyModule,
}

func init() {
	lock := py.NewLock()
	defer lock.Unlock()

	/* Setup default targets */
	stdoutTarget := log.NewTextTarget(os.Stdout)
	bufferingTarget := log.NewBufferingTarget(stdoutTarget)
	log.Targets["stdout"] = bufferingTarget

	/* Create package */
	var err error
	var module *py.Module
	if module, err = modules.Init("gem.log", []py.Method{
		{"begin_redirect", Py_begin_redirect, ""},
		{"end_redirect", Py_end_redirect, ""},
	}); err != nil {
		panic(err)
	}

	/* Register modules */
	for _, registerFunc := range moduleRegisterFuncs {
		if err = registerFunc(module); err != nil {
			panic(err)
		}
	}
}

func Py_begin_redirect() (py.Object, error) {
	if bufferingTarget, ok := log.Targets["stdout"].(*log.BufferingTarget); ok {
		bufferingTarget.Redirect()
	}
	py.None.Incref()
	return py.None, nil
}

func Py_end_redirect() (py.Object, error) {
	if bufferingTarget, ok := log.Targets["stdout"].(*log.BufferingTarget); ok {
		bufferingTarget.Flush()
	}
	py.None.Incref()
	return py.None, nil
}
