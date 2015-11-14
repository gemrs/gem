package game

import (
	"github.com/gemrs/gem/gem/runite"
	"github.com/gemrs/gem/pybind"
)

var UpdateServiceDef = pybind.Define("UpdateService", (*UpdateService)(nil))
var RegisterUpdateService = pybind.GenerateRegisterFunc(UpdateServiceDef)
var NewUpdateService = pybind.GenerateConstructor(UpdateServiceDef).(func(*runite.Context) *UpdateService)
