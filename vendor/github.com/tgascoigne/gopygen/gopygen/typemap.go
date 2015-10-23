package gopygen

import (
	"reflect"

	"github.com/qur/gopy/lib"
)

type TypeConvInFunc func(py.Object, string) (interface{}, error)
type TypeConvOutFunc func(interface{}, string) (py.Object, error)

var TypeConvIn = BaseTypeConvIn
var TypeConvOut = BaseTypeConvOut

func BaseTypeConvIn(value py.Object, typ string) (interface{}, error) {
	if typ == "" {
		typ = reflect.TypeOf(value).String()
	}

	switch v := value.(type) {
	case *py.String:
		return v.String(), nil
	case *py.Int:
		return v.Int(), nil
	}
	// don't know how to convert.. return the py.Object and hope the caller can deal with it
	return value, nil
}

func BaseTypeConvOut(value interface{}, typ string) (py.Object, error) {
	if typ == "" {
		typ = reflect.TypeOf(value).String()
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
	// BUG(tom): other types?
}
