package server

import (
	"net"

	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/log"
	"github.com/sinusoids/gem/pybind"
)

var ConnectionDef = pybind.Define("Connection", (*Connection)(nil))
var RegisterConnection = pybind.GenerateRegisterFunc(ConnectionDef)
var NewConnection = pybind.GenerateConstructor(ConnectionDef).(func(net.Conn, *log.Module) *Connection)

func (c *Connection) Py_Log(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(c.Log)
	return fn(args, kwds)
}

func (c *Connection) Py_Disconnect(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(c.Disconnect)
	return fn(args, kwds)
}

func (c *Connection) Py_Index(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(c.Index)
	return fn(args, kwds)
}
