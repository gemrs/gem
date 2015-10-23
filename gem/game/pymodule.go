package game

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/python"
)

type registerFunc func(*py.Module) error

var moduleRegisterFuncs = []registerFunc{
	RegisterGameService,
	RegisterPlayer,
	RegisterUpdateService,
	RegisterSession,
	RegisterProfile,
	RegisterSkills,
	RegisterAppearance,
	RegisterAnimations,
}

func init() {
	/* Create package */
	var err error
	var module *py.Module
	if module, err = python.InitModule("gem.game", []py.Method{}); err != nil {
		panic(err)
	}

	/* Register modules */
	for _, registerFunc := range moduleRegisterFuncs {
		if err = registerFunc(module); err != nil {
			panic(err)
		}
	}
}
