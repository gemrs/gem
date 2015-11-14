package server

import (
	"net"

	"github.com/qur/gopy/lib"

	"github.com/gemrs/willow/log"
	"github.com/gemrs/gem/pybind"
)

var ConnectionDef = pybind.Define("Connection", (*Connection)(nil))
var RegisterConnection = pybind.GenerateRegisterFunc(ConnectionDef)
var NewConnection = pybind.GenerateConstructor(ConnectionDef).(func(net.Conn, log.Log) *Connection)

func (c *Connection) PyGet_log() (py.Object, error) {
	fn := pybind.Wrap(c.Log)
	return fn(nil, nil)
}

func (c *Connection) Py_disconnect(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(c.Disconnect)
	return fn(args, kwds)
}
