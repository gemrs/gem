package player

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/game/server"
	"github.com/sinusoids/gem/pybind"
)

var PlayerDef = pybind.Define("Player", (*Player)(nil))
var RegisterPlayer = pybind.GenerateRegisterFunc(PlayerDef)
var NewPlayer = pybind.GenerateConstructor(PlayerDef).(func(*server.Connection) *Player)

func (client *Player) PyGet_session() (py.Object, error) {
	fn := pybind.Wrap(client.Session)
	return fn(nil, nil)
}

func (client *Player) PyGet_profile() (py.Object, error) {
	fn := pybind.Wrap(client.Profile)
	return fn(nil, nil)
}

func (client *Player) Py_send_message(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(client.SendMessage)
	return fn(args, kwds)
}

func (client *Player) Py_warp(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(client.SetPosition)
	return fn(args, kwds)
}

func (client *Player) PyGet_appearance() (py.Object, error) {
	fn := pybind.Wrap(client.Appearance)
	return fn(nil, nil)
}

func (client *Player) PySet_appearance(value py.Object) error {
	fn := pybind.Wrap(client.SetAppearance)
	args, err := py.PackTuple(value)
	if err != nil {
		return err
	}
	_, err = fn(args, nil)
	return err
}

func (client *Player) PyGet_entity_type() (py.Object, error) {
	fn := pybind.Wrap(client.EntityType)
	return fn(nil, nil)
}
