package player

import (
	"github.com/qur/gopy/lib"
)

type registerFunc func(*py.Module) error

var moduleRegisterFuncs = []registerFunc{
	RegisterSession,
	RegisterProfile,
}

func InitPyModule(parent *py.Module) error {
	/* Create package */
	var err error
	var module *py.Module
	if module, err = py.InitModule("gem.game.player", []py.Method{}); err != nil {
		return err
	}

	/* Register modules */
	for _, registerFunc := range moduleRegisterFuncs {
		if err = registerFunc(module); err != nil {
			return err
		}
	}

	if err = parent.AddObject("player", module); err != nil {
		return err
	}

	return nil
}
