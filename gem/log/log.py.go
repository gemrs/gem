package log

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/pybind"
)

var SysLogDef = pybind.Define("SysLog", (*SysLog)(nil))
var RegisterSysLog = pybind.GenerateRegisterFunc(SysLogDef)
var NewSysLog = pybind.GenerateConstructor(SysLogDef).(func() *SysLog)

func (log *SysLog) Py_begin_redirect(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(log.BeginRedirect)
	return fn(args, kwds)
}

func (log *SysLog) Py_end_redirect(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(log.EndRedirect)
	return fn(args, kwds)
}

func (log *SysLog) Py_module(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(log.Module)
	return fn(args, kwds)
}

var ModuleDef = pybind.Define("Module", (*Module)(nil))
var RegisterModule = pybind.GenerateRegisterFunc(ModuleDef)
var NewModule = pybind.GenerateConstructor(ModuleDef).(func() *Module)

func (log *Module) Py_submodule(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(log.SubModule)
	return fn(args, kwds)
}

func (log *Module) Py_critical(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(log.Critical)
	return fn(args, kwds)
}

func (log *Module) Py_debug(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(log.Debug)
	return fn(args, kwds)
}

func (log *Module) Py_error(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(log.Error)
	return fn(args, kwds)
}

func (log *Module) Py_fatal(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(log.Fatal)
	return fn(args, kwds)
}

func (log *Module) Py_info(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(log.Info)
	return fn(args, kwds)
}

func (log *Module) Py_notice(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(log.Notice)
	return fn(args, kwds)
}

func (log *Module) Py_warning(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(log.Warning)
	return fn(args, kwds)
}
