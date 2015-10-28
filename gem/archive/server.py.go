package archive

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/pybind"
)

var ServerDef = pybind.Define("Server", (*Server)(nil))
var RegisterServer = pybind.GenerateRegisterFunc(ServerDef)
var NewServer = pybind.GenerateConstructor(ServerDef).(func() *Server)

func (s *Server) Py_start(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(s.Start)
	return fn(args, kwds)
}

func (s *Server) Py_stop(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(s.Stop)
	return fn(args, kwds)
}
