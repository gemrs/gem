package event

import (
	"github.com/qur/gopy/lib"

	"gem/python"
)

func init() {
	/* Create package */
	var err error
	var module *py.Module
	methods := []py.Method{
		{"register_listener", Py_RegisterListener, "register an event listener"},
		{"raise_event", Py_RaiseEvent, "raise an event"},
		{"clear", Py_Clear, "clear all events"},
	}
	if module, err = python.InitModule("gem.event", methods); err != nil {
		panic(err)
	}

	createEventConstants(module)
}
