package pybind

import (
	"reflect"

	"github.com/qur/gopy/lib"
	"github.com/tgascoigne/gopygen/gopygen"
)

func InTypes(fnType reflect.Type) []reflect.Type {
	var types []reflect.Type
	for i := 0; i < fnType.NumIn(); i++ {
		types = append(types, fnType.In(i))
	}
	return types
}

func OutTypes(fnType reflect.Type) []reflect.Type {
	var types []reflect.Type
	for i := 0; i < fnType.NumOut(); i++ {
		types = append(types, fnType.Out(i))
	}
	return types
}

func ConvertIn(types []reflect.Type, args *py.Tuple) ([]reflect.Value, error) {
	convertedArgs := []reflect.Value{}
	argsSlice := args.Slice()

	for i, arg := range argsSlice {
		arg.Incref()
		defer arg.Decref()
		convertedArg, err := gopygen.TypeConvIn(arg, types[i].String())
		if err != nil {
			return nil, err
		}

		convertedArgs = append(convertedArgs, reflect.ValueOf(convertedArg))
	}

	return convertedArgs, nil
}

func ConvertOut(values []reflect.Value) ([]py.Object, error) {
	convertedOuts := []py.Object{}

	for _, ret := range values {
		convertedOut, err := gopygen.TypeConvOut(ret.Interface(), ret.Type().String())
		if err != nil {
			return nil, err
		}
		convertedOuts = append(convertedOuts, convertedOut)
	}

	return convertedOuts, nil
}

func ReflectValues(args ...interface{}) []reflect.Value {
	values := []reflect.Value{}
	for _, a := range args {
		values = append(values, reflect.ValueOf(a))
	}
	return values
}
