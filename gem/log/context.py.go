package log

import (
	"fmt"

	"github.com/qur/gopy/lib"

	"github.com/gemrs/gem/pybind"
)

var PyContextDef = pybind.Define("LogContext", (*PyContext)(nil))
var RegisterPyContext = pybind.GenerateRegisterFunc(PyContextDef)
var NewPyContext = pybind.GenerateConstructor(PyContextDef).(func() *PyContext)

type PyContext struct {
	py.BaseObject
}

func (ctx *PyContext) Init() {}

func (ctx *PyContext) ContextMap() map[string]interface{} {
	lock := py.NewLock()
	defer lock.Unlock()

	obj, err := ctx.CallMethod("log_context", "()")
	if err != nil {
		panic(fmt.Sprintf("cannot call log_context: %v", err))
	}

	var ctxObject py.Object

	err = py.ParseTuple(obj.(*py.Tuple), "O", &ctxObject)
	if err != nil {
		return nil
	}

	var ctxIface interface{}
	ctxIface = ctxObject.(interface{})

	if p, ok := ctxIface.(*py.Dict); ok {
		ctxMap, err := p.MapString()
		if err != nil {
			return nil
		}

		// ctxMap is a map[string]py.Object
		ctxMap2 := make(map[string]interface{})
		for k, v := range ctxMap {
			ctxMap2[k] = v.(interface{})
		}

		return ctxMap2
	}
	panic("invalid ctx returned: not a *py.Dict")
}
