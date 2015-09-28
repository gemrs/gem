package auth

import (
	"github.com/qur/gopy/lib"
)

type registerFunc func(*py.Module) error

var moduleRegisterFuncs = []registerFunc{
	RegisterProviderImpl,
	createAuthConstants,
}

func InitPyModule(parent *py.Module) error {
	/* Create package */
	var err error
	var module *py.Module
	if module, err = py.InitModule("gem.auth", []py.Method{}); err != nil {
		return err
	}

	/* Register modules */
	for _, registerFunc := range moduleRegisterFuncs {
		if err = registerFunc(module); err != nil {
			return err
		}
	}

	if err = parent.AddObject("auth", module); err != nil {
		return err
	}

	return nil

}
