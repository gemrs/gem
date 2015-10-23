// Generated by gopygen; DO NOT EDIT
package event

import (
	"fmt"
	"github.com/sinusoids/gem/gem/log"

	"github.com/qur/gopy/lib"
	"github.com/tgascoigne/gopygen/gopygen"
)

// Sometimes we might generate code which doesn't use some of the above imports
// Use them here just in case
var _ = fmt.Sprintf("")
var _ = gopygen.Dummy

var PyListenerDef = py.Class{
	Name:    "PyListener",
	Flags:   py.TPFLAGS_BASETYPE,
	Pointer: (*PyListener)(nil),
}

// Registers this type with a python module
func RegisterPyListener(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = PyListenerDef.Create(); err != nil {
		return err
	}

	if err = module.AddObject("PyListener", class); err != nil {
		return err
	}

	return nil
}

// Alloc allocates an object for use in python land.
// Copies the member fields from this object to the newly allocated object
// Usage: obj := GoObject{X:1, Y: 2}.Alloc()
func NewPyListener(arg_0 py.Object) (*PyListener, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	// Allocate
	alloc_, err := PyListenerDef.Alloc(0)
	if err != nil {
		return nil, err
	}
	alloc := alloc_.(*PyListener)
	err = alloc.Init(arg_0)
	return alloc, err
}

func (obj *PyListener) PyInit(_args *py.Tuple, kwds *py.Dict) error {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return fmt.Errorf("(PyListener) PyInit: parameter length mismatch")
	}

	args[0].Incref()
	in_0, err := gopygen.TypeConvIn(args[0], "py.Object")
	if err != nil {
		return err
	}

	err = obj.Init(in_0.(py.Object))

	args[0].Decref()

	return err
}

func (obj *PyListener) PyGet_id() (py.Object, error) {
	return gopygen.TypeConvOut(obj.id, "int")
}

func (obj *PyListener) PySet_id(arg py.Object) error {
	arg.Incref()
	val, err := gopygen.TypeConvIn(arg, "int")
	if err != nil {
		return err
	}

	if _, ok := val.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		val.(py.Object).Incref()
	}
	arg.Decref()

	var tmp interface{}
	tmp = &obj.id
	obj.id = val.(int)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

func (obj *PyListener) PyGet_fn() (py.Object, error) {
	return gopygen.TypeConvOut(obj.fn, "py.Object")
}

func (obj *PyListener) PySet_fn(arg py.Object) error {
	arg.Incref()
	val, err := gopygen.TypeConvIn(arg, "py.Object")
	if err != nil {
		return err
	}

	if _, ok := val.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		val.(py.Object).Incref()
	}
	arg.Decref()

	var tmp interface{}
	tmp = &obj.fn
	obj.fn = val.(py.Object)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

func (obj *PyListener) PyGet_logger() (py.Object, error) {
	return gopygen.TypeConvOut(obj.logger, "*log.Module")
}

func (obj *PyListener) PySet_logger(arg py.Object) error {
	arg.Incref()
	val, err := gopygen.TypeConvIn(arg, "*log.Module")
	if err != nil {
		return err
	}

	if _, ok := val.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		val.(py.Object).Incref()
	}
	arg.Decref()

	var tmp interface{}
	tmp = &obj.logger
	obj.logger = val.(*log.Module)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

func (l *PyListener) Py_Id(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_Id: parameter length mismatch")
	}
	// Convert parameters

	// Make the function call

	res0 := l.Id()

	// Remove local references

	out_0, err := gopygen.TypeConvOut(res0, "int")
	if err != nil {
		return nil, err
	}
	out_0.Incref()

	return out_0, nil

}
