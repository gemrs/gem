package player

import (
	"github.com/sinusoids/gem/pybind"
)

var SessionDef = pybind.Define("Session", (*Session)(nil))
var RegisterSession = pybind.GenerateRegisterFunc(SessionDef)
var NewSession = pybind.GenerateConstructor(SessionDef).(func() *Session)
