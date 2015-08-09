package event

import (
	"github.com/qur/gopy/lib"
)

func InitPyModule(parent *py.Module) error {
	/* Create package */
	var err error
	var module *py.Module
	methods := []py.Method{
		{"register_listener", Py_RegisterListener, "register an event listener"},
		{"raise_event", Py_RaiseEvent, "raise an event"},
		{"clear", Py_Clear, "clear all events"},
	}
	if module, err = py.InitModule("gem.event", methods); err != nil {
		return err
	}

	createEventConstants(module)

	if err = parent.AddObject("event", module); err != nil {
		return err
	}

	return nil
}
