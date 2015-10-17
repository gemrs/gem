// Generated by gopygen; DO NOT EDIT
package position

import (
	"fmt"

	"github.com/qur/gopy/lib"

	"github.com/tgascoigne/gopygen/gopygen"
)

// Sometimes we might generate code which doesn't use some of the above imports
// Use them here just in case
var _ = fmt.Sprintf("")
var _ = gopygen.Dummy

var AbsoluteDef = py.Class{
	Name:    "Absolute",
	Flags:   py.TPFLAGS_BASETYPE,
	Pointer: (*Absolute)(nil),
}

// Registers this type with a python module
func RegisterAbsolute(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = AbsoluteDef.Create(); err != nil {
		return err
	}

	if err = module.AddObject("Absolute", class); err != nil {
		return err
	}

	return nil
}

// Alloc allocates an object for use in python land.
// Copies the member fields from this object to the newly allocated object
// Usage: obj := GoObject{X:1, Y: 2}.Alloc()
func NewAbsolute(arg_0 int, arg_1 int, arg_2 int) (*Absolute, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	// Allocate
	alloc_, err := AbsoluteDef.Alloc(0)
	if err != nil {
		return nil, err
	}
	alloc := alloc_.(*Absolute)
	err = alloc.Init(arg_0, arg_1, arg_2)
	return alloc, err
}

func (obj *Absolute) PyInit(_args *py.Tuple, kwds *py.Dict) error {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 3 {
		return fmt.Errorf("(Absolute) PyInit: parameter length mismatch")
	}

	args[0].Incref()
	in_0, err := gopygen.TypeConvIn(args[0], "int")
	if err != nil {
		return err
	}

	args[1].Incref()
	in_1, err := gopygen.TypeConvIn(args[1], "int")
	if err != nil {
		return err
	}

	args[2].Incref()
	in_2, err := gopygen.TypeConvIn(args[2], "int")
	if err != nil {
		return err
	}

	err = obj.Init(in_0.(int), in_1.(int), in_2.(int))

	args[0].Decref()

	args[1].Decref()

	args[2].Decref()

	return err
}

func (obj *Absolute) PyGet_X() (py.Object, error) {
	return gopygen.TypeConvOut(obj.X, "int")
}

func (obj *Absolute) PySet_X(arg py.Object) error {
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
	tmp = &obj.X
	obj.X = val.(int)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

func (obj *Absolute) PyGet_Y() (py.Object, error) {
	return gopygen.TypeConvOut(obj.Y, "int")
}

func (obj *Absolute) PySet_Y(arg py.Object) error {
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
	tmp = &obj.Y
	obj.Y = val.(int)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

func (obj *Absolute) PyGet_Z() (py.Object, error) {
	return gopygen.TypeConvOut(obj.Z, "int")
}

func (obj *Absolute) PySet_Z(arg py.Object) error {
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
	tmp = &obj.Z
	obj.Z = val.(int)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

var SectorDef = py.Class{
	Name:    "Sector",
	Flags:   py.TPFLAGS_BASETYPE,
	Pointer: (*Sector)(nil),
}

// Registers this type with a python module
func RegisterSector(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = SectorDef.Create(); err != nil {
		return err
	}

	if err = module.AddObject("Sector", class); err != nil {
		return err
	}

	return nil
}

// Alloc allocates an object for use in python land.
// Copies the member fields from this object to the newly allocated object
// Usage: obj := GoObject{X:1, Y: 2}.Alloc()
func NewSector(arg_0 int, arg_1 int, arg_2 int) (*Sector, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	// Allocate
	alloc_, err := SectorDef.Alloc(0)
	if err != nil {
		return nil, err
	}
	alloc := alloc_.(*Sector)
	err = alloc.Init(arg_0, arg_1, arg_2)
	return alloc, err
}

func (obj *Sector) PyInit(_args *py.Tuple, kwds *py.Dict) error {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 3 {
		return fmt.Errorf("(Sector) PyInit: parameter length mismatch")
	}

	args[0].Incref()
	in_0, err := gopygen.TypeConvIn(args[0], "int")
	if err != nil {
		return err
	}

	args[1].Incref()
	in_1, err := gopygen.TypeConvIn(args[1], "int")
	if err != nil {
		return err
	}

	args[2].Incref()
	in_2, err := gopygen.TypeConvIn(args[2], "int")
	if err != nil {
		return err
	}

	err = obj.Init(in_0.(int), in_1.(int), in_2.(int))

	args[0].Decref()

	args[1].Decref()

	args[2].Decref()

	return err
}

func (obj *Sector) PyGet_X() (py.Object, error) {
	return gopygen.TypeConvOut(obj.X, "int")
}

func (obj *Sector) PySet_X(arg py.Object) error {
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
	tmp = &obj.X
	obj.X = val.(int)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

func (obj *Sector) PyGet_Y() (py.Object, error) {
	return gopygen.TypeConvOut(obj.Y, "int")
}

func (obj *Sector) PySet_Y(arg py.Object) error {
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
	tmp = &obj.Y
	obj.Y = val.(int)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

func (obj *Sector) PyGet_Z() (py.Object, error) {
	return gopygen.TypeConvOut(obj.Z, "int")
}

func (obj *Sector) PySet_Z(arg py.Object) error {
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
	tmp = &obj.Z
	obj.Z = val.(int)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

var RegionDef = py.Class{
	Name:    "Region",
	Flags:   py.TPFLAGS_BASETYPE,
	Pointer: (*Region)(nil),
}

// Registers this type with a python module
func RegisterRegion(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = RegionDef.Create(); err != nil {
		return err
	}

	if err = module.AddObject("Region", class); err != nil {
		return err
	}

	return nil
}

// Alloc allocates an object for use in python land.
// Copies the member fields from this object to the newly allocated object
// Usage: obj := GoObject{X:1, Y: 2}.Alloc()
func NewRegion(arg_0 *Sector) (*Region, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	// Allocate
	alloc_, err := RegionDef.Alloc(0)
	if err != nil {
		return nil, err
	}
	alloc := alloc_.(*Region)
	err = alloc.Init(arg_0)
	return alloc, err
}

func (obj *Region) PyInit(_args *py.Tuple, kwds *py.Dict) error {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return fmt.Errorf("(Region) PyInit: parameter length mismatch")
	}

	args[0].Incref()
	in_0, err := gopygen.TypeConvIn(args[0], "*Sector")
	if err != nil {
		return err
	}

	err = obj.Init(in_0.(*Sector))

	args[0].Decref()

	return err
}

func (obj *Region) PyGet_Origin() (py.Object, error) {
	return gopygen.TypeConvOut(obj.Origin, "*Sector")
}

func (obj *Region) PySet_Origin(arg py.Object) error {
	arg.Incref()
	val, err := gopygen.TypeConvIn(arg, "*Sector")
	if err != nil {
		return err
	}

	if _, ok := val.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		val.(py.Object).Incref()
	}
	arg.Decref()

	var tmp interface{}
	tmp = &obj.Origin
	obj.Origin = val.(*Sector)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

var LocalDef = py.Class{
	Name:    "Local",
	Flags:   py.TPFLAGS_BASETYPE,
	Pointer: (*Local)(nil),
}

// Registers this type with a python module
func RegisterLocal(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = LocalDef.Create(); err != nil {
		return err
	}

	if err = module.AddObject("Local", class); err != nil {
		return err
	}

	return nil
}

// Alloc allocates an object for use in python land.
// Copies the member fields from this object to the newly allocated object
// Usage: obj := GoObject{X:1, Y: 2}.Alloc()
func NewLocal(arg_0 int, arg_1 int, arg_2 int, arg_3 *Region) (*Local, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	// Allocate
	alloc_, err := LocalDef.Alloc(0)
	if err != nil {
		return nil, err
	}
	alloc := alloc_.(*Local)
	err = alloc.Init(arg_0, arg_1, arg_2, arg_3)
	return alloc, err
}

func (obj *Local) PyInit(_args *py.Tuple, kwds *py.Dict) error {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 4 {
		return fmt.Errorf("(Local) PyInit: parameter length mismatch")
	}

	args[0].Incref()
	in_0, err := gopygen.TypeConvIn(args[0], "int")
	if err != nil {
		return err
	}

	args[1].Incref()
	in_1, err := gopygen.TypeConvIn(args[1], "int")
	if err != nil {
		return err
	}

	args[2].Incref()
	in_2, err := gopygen.TypeConvIn(args[2], "int")
	if err != nil {
		return err
	}

	args[3].Incref()
	in_3, err := gopygen.TypeConvIn(args[3], "*Region")
	if err != nil {
		return err
	}

	err = obj.Init(in_0.(int), in_1.(int), in_2.(int), in_3.(*Region))

	args[0].Decref()

	args[1].Decref()

	args[2].Decref()

	args[3].Decref()

	return err
}

func (obj *Local) PyGet_X() (py.Object, error) {
	return gopygen.TypeConvOut(obj.X, "int")
}

func (obj *Local) PySet_X(arg py.Object) error {
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
	tmp = &obj.X
	obj.X = val.(int)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

func (obj *Local) PyGet_Y() (py.Object, error) {
	return gopygen.TypeConvOut(obj.Y, "int")
}

func (obj *Local) PySet_Y(arg py.Object) error {
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
	tmp = &obj.Y
	obj.Y = val.(int)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

func (obj *Local) PyGet_Z() (py.Object, error) {
	return gopygen.TypeConvOut(obj.Z, "int")
}

func (obj *Local) PySet_Z(arg py.Object) error {
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
	tmp = &obj.Z
	obj.Z = val.(int)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

func (obj *Local) PyGet_Region() (py.Object, error) {
	return gopygen.TypeConvOut(obj.Region, "*Region")
}

func (obj *Local) PySet_Region(arg py.Object) error {
	arg.Incref()
	val, err := gopygen.TypeConvIn(arg, "*Region")
	if err != nil {
		return err
	}

	if _, ok := val.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		val.(py.Object).Incref()
	}
	arg.Decref()

	var tmp interface{}
	tmp = &obj.Region
	obj.Region = val.(*Region)

	if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
		oldObj.Decref()
	}
	return nil
}

func (pos *Absolute) Py_String(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
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

	res0 := pos.String()

	// Remove local references

	out_0, err := gopygen.TypeConvOut(res0, "string")
	if err != nil {
		return nil, err
	}

	return out_0, nil

}

func (pos *Absolute) Py_Sector(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_Sector: parameter length mismatch")
	}
	// Convert parameters

	// Make the function call

	res0 := pos.Sector()

	// Remove local references

	out_0, err := gopygen.TypeConvOut(res0, "*Sector")
	if err != nil {
		return nil, err
	}

	return out_0, nil

}

func (pos *Absolute) Py_RegionOf(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_RegionOf: parameter length mismatch")
	}
	// Convert parameters

	// Make the function call

	res0 := pos.RegionOf()

	// Remove local references

	out_0, err := gopygen.TypeConvOut(res0, "*Region")
	if err != nil {
		return nil, err
	}

	return out_0, nil

}

func (pos *Absolute) Py_LocalTo(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_LocalTo: parameter length mismatch")
	}
	// Convert parameters

	args[0].Incref()
	in_0, err := gopygen.TypeConvIn(args[0], "*Region")
	if err != nil {
		return nil, err
	}

	// Make the function call

	res0 := pos.LocalTo(in_0.(*Region))

	// Remove local references

	args[0].Decref()

	out_0, err := gopygen.TypeConvOut(res0, "*Local")
	if err != nil {
		return nil, err
	}

	return out_0, nil

}

func (region *Region) Py_Rebase(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_Rebase: parameter length mismatch")
	}
	// Convert parameters

	args[0].Incref()
	in_0, err := gopygen.TypeConvIn(args[0], "*Absolute")
	if err != nil {
		return nil, err
	}

	// Make the function call

	region.Rebase(in_0.(*Absolute))

	// Remove local references

	args[0].Decref()

	py.None.Incref()
	return py.None, nil

}

func (region *Region) Py_SectorDelta(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_SectorDelta: parameter length mismatch")
	}
	// Convert parameters

	args[0].Incref()
	in_0, err := gopygen.TypeConvIn(args[0], "*Region")
	if err != nil {
		return nil, err
	}

	// Make the function call

	res0, res1, res2 := region.SectorDelta(in_0.(*Region))

	// Remove local references

	args[0].Decref()

	out_0, err := gopygen.TypeConvOut(res0, "int")
	if err != nil {
		return nil, err
	}

	out_1, err := gopygen.TypeConvOut(res1, "int")
	if err != nil {
		return nil, err
	}

	out_2, err := gopygen.TypeConvOut(res2, "int")
	if err != nil {
		return nil, err
	}

	return py.PackTuple(out_0, out_1, out_2)

}

func (local *Local) Py_Absolute(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_Absolute: parameter length mismatch")
	}
	// Convert parameters

	// Make the function call

	res0 := local.Absolute()

	// Remove local references

	out_0, err := gopygen.TypeConvOut(res0, "*Absolute")
	if err != nil {
		return nil, err
	}

	return out_0, nil

}
