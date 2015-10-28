package pybind 

import (
	"reflect"

	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/python"
)

func TypeConvIn(value py.Object, typ string) (interface{}, error) {
	if typ == "" {
		typ = reflect.TypeOf(value).String()
	}

	if typ == "[]string" {
		if list, ok := value.(*py.List); ok {
			output := make([]string, list.Size())
			for i, obj := range list.Slice() {
				str, err := TypeConvIn(obj, "string")
				if err != nil {
					return nil, err
				}
				output[i] = str.(string)
			}
			return output, nil
		}
	}

	switch v := value.(type) {
	case *py.String:
		return v.String(), nil
	case *py.Int:
		return v.Int(), nil
	}
	return value, nil
}

func TypeConvOut(value interface{}, typ string) (py.Object, error) {
	if typ == "" {
		typ = reflect.TypeOf(value).String()
	}

	if typ == "error" {
		if value != nil {
			err := value.(error)
			python.Raise(err)
			return nil, nil
		} else {
			return py.None, nil
		}
	}
	return py.BuildValue(PyTupleFormatString(typ), value)
}

func PyTupleFormatString(typ string) string {
	if str, ok := PyTupleTypeMap[typ]; ok {
		return str
	}
	return "O"
}

var PyTupleTypeMap = map[string]string{
	"int":     "i",
	"int8":    "c",
	"int16":   "h",
	"int32":   "l",
	"int64":   "L",
	"uint":    "I",
	"uint8":   "b",
	"uint16":  "H",
	"uint32":  "k",
	"uint64":  "K",
	"string":  "s",
	"float32": "f",
	"float64": "d",
}
