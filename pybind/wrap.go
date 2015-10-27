package pybind

import (
	"reflect"

	"github.com/qur/gopy/lib"
)

type Constructor func(*py.Type, *py.Tuple, *py.Dict) (py.Object, error)
type Wrapper func(*py.Tuple, *py.Dict) (py.Object, error)

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

		/*		val := reflect.ValueOf(fn)
				typ := reflect.TypeOf(fn)

				var obj interface{}
				args, err := convertIn(typ, pyArgs)
				if err != nil {
					return nil, err
				}

				out := val.Call(args)
				obj = out[0].Interface()
				if errIf := out[1].Interface(); errIf != nil {
					return nil, errIf.(error)
				}

				fmt.Println("created obj, setting base")

				if base, ok := obj.(*BaseObject); ok {
					base.SetBase(pyobj)
				}
				fmt.Println("set base, returning obj")
				return obj.(py.Object), nil
		*/
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
