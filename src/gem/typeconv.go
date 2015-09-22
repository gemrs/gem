package gem

import (
	"github.com/qur/gopy/lib"
	"github.com/tgascoigne/gopygen/gopygen"
)

func init() {
	gopygen.TypeConvIn = TypeConvIn
	gopygen.TypeConvOut = TypeConvOut
}

func TypeConvIn(value py.Object, typ string) (interface{}, error) {
	return gopygen.BaseTypeConvIn(value, typ)
}

func TypeConvOut(value interface{}, typ string) (py.Object, error) {
	if typ == "error" {
		if value != nil {
			err := value.(error)
			py.Raise(err)
			return nil, nil
		} else {
			return py.None, nil
		}
	}
	return gopygen.BaseTypeConvOut(value, typ)
}
