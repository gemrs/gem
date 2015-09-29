package py

// #include <python2.7/Python.h>
// #cgo LDFLAGS: -lpython2.7
import "C"

func Raise(err error) {
	C.PyErr_SetString(C.PyExc_RuntimeError, C.CString(err.Error()))
}
