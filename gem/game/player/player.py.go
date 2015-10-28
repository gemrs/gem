package player

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/game/server"
	"github.com/sinusoids/gem/pybind"
)

var PlayerDef = pybind.Define("Player", (*Player)(nil))
var RegisterPlayer = pybind.GenerateRegisterFunc(PlayerDef)
var NewPlayer = pybind.GenerateConstructor(PlayerDef).(func(*server.Connection) *Player)

func (client *Player) Py_Session(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(client.Session)
	return fn(args, kwds)
}

func (client *Player) Py_Profile(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(client.Profile)
	return fn(args, kwds)
}

func (client *Player) Py_SendMessage(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(client.SendMessage)
	return fn(args, kwds)
}

func (client *Player) Py_SetPosition(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(client.SetPosition)
	return fn(args, kwds)
}

func (client *Player) Py_SetAppearance(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(client.SetAppearance)
	return fn(args, kwds)
}

func (client *Player) Py_EntityType(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(client.EntityType)
	return fn(args, kwds)
}
