// Generated by gopygen; DO NOT EDIT
package player

import (
	"fmt"
	"gem/game/position"

	"github.com/qur/gopy/lib"

	"github.com/tgascoigne/gopygen/gopygen"
)

// Sometimes we might generate code which doesn't use some of the above imports
// Use them here just in case
var _ = fmt.Sprintf("")
var _ = gopygen.Dummy

var ProfileDef = py.Class{
	Name:    "Profile",
	Flags:   py.TPFLAGS_BASETYPE,
	Pointer: (*Profile)(nil),
}

// Registers this type with a python module
func RegisterProfile(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = ProfileDef.Create(); err != nil {
		return err
	}

	if err = module.AddObject("Profile", class); err != nil {
		return err
	}

	return nil
}

func (obj *Profile) PyGet_Username() (py.Object, error) {
	return gopygen.TypeConvOut(obj.Username, "string")
}

func (obj *Profile) PySet_Username(arg py.Object) error {
	arg.Incref()
	val, err := gopygen.TypeConvIn(arg, "string")
	if err != nil {
		return err
	}

	if _, ok := val.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		val.(py.Object).Incref()
	}
	arg.Decref()

	var tmp interface{}
	tmp = &obj.Username
	obj.Username = val.(string)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

func (obj *Profile) PyGet_Password() (py.Object, error) {
	return gopygen.TypeConvOut(obj.Password, "string")
}

func (obj *Profile) PySet_Password(arg py.Object) error {
	arg.Incref()
	val, err := gopygen.TypeConvIn(arg, "string")
	if err != nil {
		return err
	}

	if _, ok := val.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		val.(py.Object).Incref()
	}
	arg.Decref()

	var tmp interface{}
	tmp = &obj.Password
	obj.Password = val.(string)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

func (obj *Profile) PyGet_Rights() (py.Object, error) {
	return gopygen.TypeConvOut(obj.Rights, "Rights")
}

func (obj *Profile) PySet_Rights(arg py.Object) error {
	arg.Incref()
	val, err := gopygen.TypeConvIn(arg, "Rights")
	if err != nil {
		return err
	}

	if _, ok := val.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		val.(py.Object).Incref()
	}
	arg.Decref()

	var tmp interface{}
	tmp = &obj.Rights
	obj.Rights = val.(Rights)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

func (obj *Profile) PyGet_Pos() (py.Object, error) {
	return gopygen.TypeConvOut(obj.Pos, "*position.Absolute")
}

func (obj *Profile) PySet_Pos(arg py.Object) error {
	arg.Incref()
	val, err := gopygen.TypeConvIn(arg, "*position.Absolute")
	if err != nil {
		return err
	}

	if _, ok := val.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		val.(py.Object).Incref()
	}
	arg.Decref()

	var tmp interface{}
	tmp = &obj.Pos
	obj.Pos = val.(*position.Absolute)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

func (p *Profile) Py_String(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_String: parameter length mismatch")
	}
	// Convert parameters

	// Make the function call

	res0 := p.String()

	// Remove local references

	out_0, err := gopygen.TypeConvOut(res0, "string")
	if err != nil {
		return nil, err
	}

	return out_0, nil

}
