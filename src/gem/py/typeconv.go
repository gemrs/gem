package py

import (
	"github.com/qur/gopy/lib"
	"github.com/tgascoigne/gopygen/gopygen"
)

func SetTypeConvFuncs() {
	gopygen.TypeConvIn = TypeConvIn
	gopygen.TypeConvOut = TypeConvOut
}

func TypeConvIn(value py.Object, typ string) (interface{}, error) {
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
	return gopygen.BaseTypeConvIn(value, typ)
}

func TypeConvOut(value interface{}, typ string) (py.Object, error) {
	if typ == "error" {
		if value != nil {
			err := value.(error)
			Raise(err)
			return nil, nil
		} else {
			return py.None, nil
		}
	}
	return gopygen.BaseTypeConvOut(value, typ)
}
