package game

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/auth"
	"github.com/sinusoids/gem/gem/runite"
	"github.com/sinusoids/gem/pybind"
)

var GameServiceDef = pybind.Define("GameService", (*GameService)(nil))
var RegisterGameService = pybind.GenerateRegisterFunc(GameServiceDef)
var NewGameService = pybind.GenerateConstructor(GameServiceDef).(func(*runite.Context, string, auth.Provider) (*GameService, error))

func (svc *GameService) PyGet_world() (py.Object, error) {
	fn := pybind.Wrap(svc.World)
	return fn(nil, nil)
}
