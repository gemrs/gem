package game

import (
	"github.com/qur/gopy/lib"

	"gem/game/player"
)

type registerFunc func(*py.Module) error

var moduleRegisterFuncs = []registerFunc{
	RegisterConnection,
	RegisterServer,
	player.InitPyModule,
}

func InitPyModule(parent *py.Module) error {
	/* Create package */
	var err error
	var module *py.Module
	if module, err = py.InitModule("gem.game", []py.Method{}); err != nil {
		return err
	}

	/* Register modules */
	for _, registerFunc := range moduleRegisterFuncs {
		if err = registerFunc(module); err != nil {
			return err
		}
	}

	if err = parent.AddObject("game", module); err != nil {
		return err
	}

	return nil
}
