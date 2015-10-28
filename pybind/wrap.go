package pybind

import (
	"reflect"

	"github.com/qur/gopy/lib"
)

type Constructor func(*py.Type, *py.Tuple, *py.Dict) (py.Object, error)
type Wrapper func(*py.Tuple, *py.Dict) (py.Object, error)

func Define(name string, ptr interface{}, init interface{}) py.Class {
	return py.Class{
		Name:    name,
		Flags:   py.TPFLAGS_BASETYPE,
		Pointer: ptr,
		New:     WrapConstructor(init),
	}
}

func GenerateRegisterFunc(def py.Class) func(*py.Module) error {
	return func(module *py.Module) error {
		return Register(def, module)
	}
}

func Register(def py.Class, module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = def.Create(); err != nil {
		return err
	}

	if err = module.AddObject(def.Name, class); err != nil {
		return err
	}

	return nil
}

func WrapConstructor(fn interface{}) Constructor {
	wrapper := Wrap(fn)
	return func(pyType *py.Type, pyArgs *py.Tuple, pyKwds *py.Dict) (py.Object, error) {
		lock := py.NewLock()
		defer lock.Unlock()

		pyObj, err := pyType.Alloc(0)
		if err != nil {
			return nil, err
		}

		// prepend the object to pyArgs
		argsSlice := pyArgs.Slice()
		argsSlice = append([]py.Object{pyObj}, argsSlice...)
		pyArgs, err = py.PackTuple(argsSlice...)
		if err != nil {
			return nil, err
		}

		_, err = wrapper(pyArgs, pyKwds)
		if err != nil {
			return nil, err
		}

		return pyObj, nil
	}
}

func Wrap(fn interface{}) Wrapper {
	val := reflect.ValueOf(fn)
	typ := reflect.TypeOf(fn)

	var inTypes []reflect.Type
	for i := 0; i < typ.NumIn(); i++ {
		inTypes = append(inTypes, typ.In(i))
	}

	return func(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
		lock := py.NewLock()
		defer lock.Unlock()

		// Convert args to []Value
		convertedArgs, err := ConvertIn(inTypes, args)
		if err != nil {
			return nil, err
		}

		// Call the go function
		outs := val.Call(convertedArgs)

		// Convert outs to []py.Object
		convertedOuts, err := ConvertOut(outs)
		if err != nil {
			return nil, err
		}

		// Return
		if len(convertedOuts) == 1 {
			return convertedOuts[0], nil
		} else {
			return py.PackTuple(convertedOuts...)
		}
	}
}
