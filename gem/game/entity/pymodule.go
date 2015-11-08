package entity

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/game/interface/entity"
	"github.com/sinusoids/gem/gem/python/modules"
)

type registerFunc func(*py.Module) error

var moduleRegisterFuncs = []registerFunc{
	RegisterGenericMob,
	entity.RegisterCollection,
	entity.RegisterSlice,
	entity.RegisterList,
}

func init() {
	lock := py.NewLock()
	defer lock.Unlock()

	/* Create package */
	var err error
	var module *py.Module
	if module, err = modules.Init("gem.game.entity", []py.Method{}); err != nil {
		panic(err)
	}

	/* Register modules */
	for _, registerFunc := range moduleRegisterFuncs {
		if err = registerFunc(module); err != nil {
			panic(err)
		}
	}

	createEntityTypeConstants(module)
}

var entityTypeConstants = map[string]entity.EntityType{
	"IncompleteType": entity.IncompleteType,
	"PlayerType":     entity.PlayerType,
}

func createEntityTypeConstants(module *py.Module) {
	for identifier, typ := range entityTypeConstants {
		if pyString, err := py.NewString((string)(typ)); err != nil {
			panic(err)
		} else if err := module.AddObject(identifier, pyString); err != nil {
			panic(err)
		}
	}
}
