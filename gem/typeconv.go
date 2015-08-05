package gem

import (
	"github.com/qur/gopy/lib"
	"github.com/tgascoigne/gopygen/gopygen"
)

func init() {
	gopygen.TypeConvIn = TypeConvIn
	gopygen.TypeConvOut = TypeConvOut
}

func TypeConvIn(typ string, value py.Object) (interface{}, error) {
	switch typ {
	case "py.Object":
		return value, nil
	}

	return gopygen.BaseTypeConvIn(typ, value)
}

func TypeConvOut(typ string, value interface{}) (py.Object, error) {
	return gopygen.BaseTypeConvOut(typ, value)
}
