package game

import (
	"github.com/sinusoids/gem/gem/runite"
	"github.com/sinusoids/gem/pybind"
)

var UpdateServiceDef = pybind.Define("UpdateService", (*UpdateService)(nil))
var RegisterUpdateService = pybind.GenerateRegisterFunc(UpdateServiceDef)
var NewUpdateService = pybind.GenerateConstructor(UpdateServiceDef).(func(*runite.Context) *UpdateService)
