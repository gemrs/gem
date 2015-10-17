// Generated by gopygen; DO NOT EDIT
package log

import (
	"bytes"
	"fmt"

	"github.com/op/go-logging"
	"github.com/qur/gopy/lib"

	"github.com/tgascoigne/gopygen/gopygen"
)

// Sometimes we might generate code which doesn't use some of the above imports
// Use them here just in case
var _ = fmt.Sprintf("")
var _ = gopygen.Dummy

var SysLogDef = py.Class{
	Name:    "SysLog",
	Flags:   py.TPFLAGS_BASETYPE,
	Pointer: (*SysLog)(nil),
}

// Registers this type with a python module
func RegisterSysLog(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = SysLogDef.Create(); err != nil {
		return err
	}

	if err = module.AddObject("SysLog", class); err != nil {
		return err
	}

	return nil
}

// Alloc allocates an object for use in python land.
// Copies the member fields from this object to the newly allocated object
// Usage: obj := GoObject{X:1, Y: 2}.Alloc()
func NewSysLog() (*SysLog, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	// Allocate
	alloc_, err := SysLogDef.Alloc(0)
	if err != nil {
		return nil, err
	}
	alloc := alloc_.(*SysLog)
	err = alloc.Init()
	return alloc, err
}

func (obj *SysLog) PyInit(_args *py.Tuple, kwds *py.Dict) error {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return fmt.Errorf("(SysLog) PyInit: parameter length mismatch")
	}

	return obj.Init()
}

func (obj *SysLog) PyGet_redirectBuffer() (py.Object, error) {
	return gopygen.TypeConvOut(obj.redirectBuffer, "*bytes.Buffer")
}

func (obj *SysLog) PySet_redirectBuffer(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "*bytes.Buffer")
	if err != nil {
		return err
	}
	obj.redirectBuffer = val.(*bytes.Buffer)
	return nil
}

func (obj *SysLog) PyGet_modules() (py.Object, error) {
	return gopygen.TypeConvOut(obj.modules, "map[string]*Module")
}

func (obj *SysLog) PySet_modules(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "map[string]*Module")
	if err != nil {
		return err
	}
	obj.modules = val.(map[string]*Module)
	return nil
}

var ModuleDef = py.Class{
	Name:    "Module",
	Flags:   py.TPFLAGS_BASETYPE,
	Pointer: (*Module)(nil),
}

// Registers this type with a python module
func RegisterModule(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = ModuleDef.Create(); err != nil {
		return err
	}

	if err = module.AddObject("Module", class); err != nil {
		return err
	}

	return nil
}

// Alloc allocates an object for use in python land.
// Copies the member fields from this object to the newly allocated object
// Usage: obj := GoObject{X:1, Y: 2}.Alloc()
func NewModule() (*Module, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	// Allocate
	alloc_, err := ModuleDef.Alloc(0)
	if err != nil {
		return nil, err
	}
	alloc := alloc_.(*Module)
	err = alloc.Init()
	return alloc, err
}

func (obj *Module) PyInit(_args *py.Tuple, kwds *py.Dict) error {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return fmt.Errorf("(Module) PyInit: parameter length mismatch")
	}

	return obj.Init()
}

func (obj *Module) PyGet_logger() (py.Object, error) {
	return gopygen.TypeConvOut(obj.logger, "*logging.Logger")
}

func (obj *Module) PySet_logger(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "*logging.Logger")
	if err != nil {
		return err
	}
	obj.logger = val.(*logging.Logger)
	return nil
}

func (obj *Module) PyGet_parent() (py.Object, error) {
	return gopygen.TypeConvOut(obj.parent, "*Module")
}

func (obj *Module) PySet_parent(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "*Module")
	if err != nil {
		return err
	}
	obj.parent = val.(*Module)
	return nil
}

func (obj *Module) PyGet_prefix() (py.Object, error) {
	return gopygen.TypeConvOut(obj.prefix, "string")
}

func (obj *Module) PySet_prefix(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "string")
	if err != nil {
		return err
	}
	obj.prefix = val.(string)
	return nil
}

func (log *SysLog) Py_BeginRedirect(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_BeginRedirect: parameter length mismatch")
	}

	log.BeginRedirect()

	py.None.Incref()
	return py.None, nil

}

func (log *SysLog) Py_EndRedirect(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_EndRedirect: parameter length mismatch")
	}

	log.EndRedirect()

	py.None.Incref()
	return py.None, nil

}

func (log *SysLog) Py_Module(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_Module: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "string")
	if err != nil {
		return nil, err
	}

	res0 := log.Module(in_0.(string))

	out_0, err := gopygen.TypeConvOut(res0, "*Module")
	if err != nil {
		return nil, err
	}

	return out_0, nil

}

func (log *Module) Py_SubModule(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_SubModule: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "string")
	if err != nil {
		return nil, err
	}

	res0 := log.SubModule(in_0.(string))

	out_0, err := gopygen.TypeConvOut(res0, "*Module")
	if err != nil {
		return nil, err
	}

	return out_0, nil

}

func (log *Module) Py_Critical(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_Critical: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "string")
	if err != nil {
		return nil, err
	}

	log.Critical(in_0.(string))

	py.None.Incref()
	return py.None, nil

}

func (log *Module) Py_Debug(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_Debug: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "string")
	if err != nil {
		return nil, err
	}

	log.Debug(in_0.(string))

	py.None.Incref()
	return py.None, nil

}

func (log *Module) Py_Error(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_Error: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "string")
	if err != nil {
		return nil, err
	}

	log.Error(in_0.(string))

	py.None.Incref()
	return py.None, nil

}

func (log *Module) Py_Fatal(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_Fatal: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "string")
	if err != nil {
		return nil, err
	}

	log.Fatal(in_0.(string))

	py.None.Incref()
	return py.None, nil

}

func (log *Module) Py_Info(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_Info: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "string")
	if err != nil {
		return nil, err
	}

	log.Info(in_0.(string))

	py.None.Incref()
	return py.None, nil

}

func (log *Module) Py_Notice(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_Notice: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "string")
	if err != nil {
		return nil, err
	}

	log.Notice(in_0.(string))

	py.None.Incref()
	return py.None, nil

}

func (log *Module) Py_Warning(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_Warning: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "string")
	if err != nil {
		return nil, err
	}

	log.Warning(in_0.(string))

	py.None.Incref()
	return py.None, nil

}
