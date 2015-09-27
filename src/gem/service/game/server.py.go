// Generated by gopygen; DO NOT EDIT
package game

import (
	"fmt"
	"gem/runite"
	"net"

	"gopkg.in/tomb.v2"

	"github.com/qur/gopy/lib"
	"github.com/tgascoigne/gopygen/gopygen"
)

// Sometimes we might generate code which doesn't use some of the above imports
// Use them here just in case
var _ = fmt.Sprintf("")
var _ = gopygen.Dummy

var ServerDef = py.Class{
	Name:    "Server",
	Pointer: (*Server)(nil),
}

// Registers this type with a python module
func RegisterServer(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = ServerDef.Create(); err != nil {
		return err
	}

	if err = module.AddObject("Server", class); err != nil {
		return err
	}

	return nil
}

// Alloc allocates an object for use in python land.
// Copies the member fields from this object to the newly allocated object
// Usage: obj := GoObject{X:1, Y: 2}.Alloc()
func (obj Server) Alloc() (*Server, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	// Allocate
	alloc_, err := ServerDef.Alloc(0)
	if err != nil {
		return nil, err
	}
	alloc := alloc_.(*Server)
	// Copy fields

	alloc.laddr = obj.laddr

	alloc.ln = obj.ln

	alloc.update = obj.update

	alloc.runite = obj.runite

	alloc.clients = obj.clients

	alloc.nextIndex = obj.nextIndex

	alloc.t = obj.t

	return alloc, nil
}

func (obj *Server) PyGet_laddr() (py.Object, error) {
	return gopygen.TypeConvOut(obj.laddr, "string")
}

func (obj *Server) PySet_laddr(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "string")
	if err != nil {
		return err
	}
	obj.laddr = val.(string)
	return nil
}

func (obj *Server) PyGet_ln() (py.Object, error) {
	return gopygen.TypeConvOut(obj.ln, "net.Listener")
}

func (obj *Server) PySet_ln(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "net.Listener")
	if err != nil {
		return err
	}
	obj.ln = val.(net.Listener)
	return nil
}

func (obj *Server) PyGet_update() (py.Object, error) {
	return gopygen.TypeConvOut(obj.update, "*updateService")
}

func (obj *Server) PySet_update(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "*updateService")
	if err != nil {
		return err
	}
	obj.update = val.(*updateService)
	return nil
}

func (obj *Server) PyGet_runite() (py.Object, error) {
	return gopygen.TypeConvOut(obj.runite, "*runite.Context")
}

func (obj *Server) PySet_runite(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "*runite.Context")
	if err != nil {
		return err
	}
	obj.runite = val.(*runite.Context)
	return nil
}

func (obj *Server) PyGet_clients() (py.Object, error) {
	return gopygen.TypeConvOut(obj.clients, "map[Index]*GameConnection")
}

func (obj *Server) PySet_clients(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "map[Index]*GameConnection")
	if err != nil {
		return err
	}
	obj.clients = val.(map[Index]*GameConnection)
	return nil
}

func (obj *Server) PyGet_nextIndex() (py.Object, error) {
	return gopygen.TypeConvOut(obj.nextIndex, "Index")
}

func (obj *Server) PySet_nextIndex(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "Index")
	if err != nil {
		return err
	}
	obj.nextIndex = val.(Index)
	return nil
}

func (obj *Server) PyGet_t() (py.Object, error) {
	return gopygen.TypeConvOut(obj.t, "tomb.Tomb")
}

func (obj *Server) PySet_t(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "tomb.Tomb")
	if err != nil {
		return err
	}
	obj.t = val.(tomb.Tomb)
	return nil
}

func (s *Server) Py_Start(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 2 {
		return nil, fmt.Errorf("Py_Start: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "string")
	if err != nil {
		return nil, err
	}

	in_1, err := gopygen.TypeConvIn(args[1], "*runite.Context")
	if err != nil {
		return nil, err
	}

	res0 := s.Start(in_0.(string), in_1.(*runite.Context))

	out_0, err := gopygen.TypeConvOut(res0, "error")
	if err != nil {
		return nil, err
	}

	return out_0, nil

}

func (s *Server) Py_Stop(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_Stop: parameter length mismatch")
	}

	res0 := s.Stop()

	out_0, err := gopygen.TypeConvOut(res0, "error")
	if err != nil {
		return nil, err
	}

	return out_0, nil

}