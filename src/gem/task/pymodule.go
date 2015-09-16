package task

import (
	"github.com/qur/gopy/lib"
)

type registerFunc func(*py.Module) error

var moduleRegisterFuncs = []registerFunc{}

func InitPyModule(parent *py.Module) error {
	/* Create package */
	var err error
	var module *py.Module
	methods := []py.Method{
		{"submit", Py_Submit, "submit a task to the scheduler"},
	}
	if module, err = py.InitModule("gem.task", methods); err != nil {
		return err
	}

	createTaskHookConstants(module)

	/* Register modules */
	for _, registerFunc := range moduleRegisterFuncs {
		if err = registerFunc(module); err != nil {
			return err
		}
	}

	if err = parent.AddObject("task", module); err != nil {
		return err
	}

	return nil
}
