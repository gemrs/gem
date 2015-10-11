// Generated by gopygen; DO NOT EDIT
package player

import (
	"fmt"
	"gem/encoding"

	"github.com/gtank/isaac"

	"github.com/qur/gopy/lib"
	"github.com/tgascoigne/gopygen/gopygen"
)

// Sometimes we might generate code which doesn't use some of the above imports
// Use them here just in case
var _ = fmt.Sprintf("")
var _ = gopygen.Dummy

var SessionDef = py.Class{
	Name:    "Session",
	Pointer: (*Session)(nil),
}

// Registers this type with a python module
func RegisterSession(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = SessionDef.Create(); err != nil {
		return err
	}

	if err = module.AddObject("Session", class); err != nil {
		return err
	}

	return nil
}

// Alloc allocates an object for use in python land.
// Copies the member fields from this object to the newly allocated object
// Usage: obj := GoObject{X:1, Y: 2}.Alloc()
func (obj Session) Alloc() (*Session, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	// Allocate
	alloc_, err := SessionDef.Alloc(0)
	if err != nil {
		return nil, err
	}
	alloc := alloc_.(*Session)
	// Copy fields

	alloc.RandIn = obj.RandIn

	alloc.RandOut = obj.RandOut

	alloc.RandKey = obj.RandKey

	alloc.SecureBlockSize = obj.SecureBlockSize

	alloc.target = obj.target

	return alloc, nil
}

func (obj *Session) PyGet_RandIn() (py.Object, error) {
	return gopygen.TypeConvOut(obj.RandIn, "isaac.ISAAC")
}

func (obj *Session) PySet_RandIn(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "isaac.ISAAC")
	if err != nil {
		return err
	}
	obj.RandIn = val.(isaac.ISAAC)
	return nil
}

func (obj *Session) PyGet_RandOut() (py.Object, error) {
	return gopygen.TypeConvOut(obj.RandOut, "isaac.ISAAC")
}

func (obj *Session) PySet_RandOut(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "isaac.ISAAC")
	if err != nil {
		return err
	}
	obj.RandOut = val.(isaac.ISAAC)
	return nil
}

func (obj *Session) PyGet_RandKey() (py.Object, error) {
	return gopygen.TypeConvOut(obj.RandKey, "[]int32")
}

func (obj *Session) PySet_RandKey(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "[]int32")
	if err != nil {
		return err
	}
	obj.RandKey = val.([]int32)
	return nil
}

func (obj *Session) PyGet_SecureBlockSize() (py.Object, error) {
	return gopygen.TypeConvOut(obj.SecureBlockSize, "int")
}

func (obj *Session) PySet_SecureBlockSize(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "int")
	if err != nil {
		return err
	}
	obj.SecureBlockSize = val.(int)
	return nil
}

func (obj *Session) PyGet_target() (py.Object, error) {
	return gopygen.TypeConvOut(obj.target, "encoding.Writer")
}

func (obj *Session) PySet_target(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "encoding.Writer")
	if err != nil {
		return err
	}
	obj.target = val.(encoding.Writer)
	return nil
}

func (session *Session) Py_SendMessage(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_SendMessage: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "string")
	if err != nil {
		return nil, err
	}

	session.SendMessage(in_0.(string))

	py.None.Incref()
	return py.None, nil

}

func (session *Session) Py_SetTarget(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_SetTarget: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "encoding.Writer")
	if err != nil {
		return nil, err
	}

	session.SetTarget(in_0.(encoding.Writer))

	py.None.Incref()
	return py.None, nil

}