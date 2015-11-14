package log

import (
	"github.com/qur/gopy/lib"

	"github.com/gemrs/gem/pybind"
)

var PyModuleDef = pybind.Define("Module", (*PyModule)(nil))
var RegisterPyModule = pybind.GenerateRegisterFunc(PyModuleDef)
var NewPyModule = pybind.GenerateConstructor(PyModuleDef).(func(string, py.Object) *PyModule)

type PyModule struct {
	py.BaseObject

	*Module
}

func (m *PyModule) Init(tag string, ctxObj py.Object) {
	var ctx Context
	if c, ok := ctxObj.(*PyContext); ok {
		ctx = c
	} else {
		ctx = NilContext
	}
	m.Module = New(tag, ctx).(*Module)
}

func (m *PyModule) PyGet_tag() (py.Object, error) {
	fn := pybind.Wrap(func() string {
		return m.Module.tag
	})
	return fn(nil, nil)
}

func (m *PyModule) PyGet_ctx() (py.Object, error) {
	fn := pybind.Wrap(func() py.Object {
		if ctx, ok := m.Module.ctx.(*PyContext); ok {
			return ctx
		}
		py.None.Incref()
		return py.None
	})
	return fn(nil, nil)
}

func (m *PyModule) Py_child(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(func(tag string, ctx py.Object) *PyModule {
		return NewPyModule(m.Module.tag+"/"+tag, ctx)
	})
	return fn(args, kwds)
}

func (m *PyModule) Py_debug(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(func(message string) {
		m.Module.Debug(message)
	})
	return fn(args, kwds)
}

func (m *PyModule) Py_error(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(func(message string) {
		m.Module.Error(message)
	})
	return fn(args, kwds)
}

func (m *PyModule) Py_info(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(func(message string) {
		m.Module.Info(message)
	})
	return fn(args, kwds)
}

func (m *PyModule) Py_notice(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(func(message string) {
		m.Module.Notice(message)
	})
	return fn(args, kwds)
}
