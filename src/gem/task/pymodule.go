package task

import (
	"github.com/qur/gopy/lib"

	"gem/python"
)

type registerFunc func(*py.Module) error

var moduleRegisterFuncs = []registerFunc{}

func init() {
	/* Create package */
	var err error
	var module *py.Module
	methods := []py.Method{
		{"submit", Py_Submit, "submit a task to the scheduler"},
	}
	if module, err = python.InitModule("gem.task", methods); err != nil {
		panic(err)
	}

	createTaskHookConstants(module)

	/* Register modules */
	for _, registerFunc := range moduleRegisterFuncs {
		if err = registerFunc(module); err != nil {
			panic(err)
		}
	}
}
