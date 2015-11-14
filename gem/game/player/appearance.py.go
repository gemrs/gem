package player

import (
	"github.com/gemrs/gem/pybind"
)

var AppearanceDef = pybind.Define("Appearance", (*Appearance)(nil))
var RegisterAppearance = pybind.GenerateRegisterFunc(AppearanceDef)
var NewAppearance = pybind.GenerateConstructor(AppearanceDef).(func() *Appearance)
