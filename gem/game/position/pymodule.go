package position

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/python/modules"
)

type registerFunc func(*py.Module) error

var moduleRegisterFuncs = []registerFunc{
	RegisterAbsolute,
	RegisterSector,
	RegisterRegion,
	RegisterLocal,
}

func init() {
	lock := py.NewLock()
	defer lock.Unlock()

	/* Create package */
	var err error
	var module *py.Module
	if module, err = modules.Init("gem.game.position", []py.Method{}); err != nil {
		panic(err)
	}

	/* Register modules */
	for _, registerFunc := range moduleRegisterFuncs {
		if err = registerFunc(module); err != nil {
			panic(err)
		}
	}
}