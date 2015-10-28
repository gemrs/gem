package player

import (
	"github.com/sinusoids/gem/pybind"
)

var AppearanceDef = pybind.Define("Appearance", (*Appearance)(nil))
var RegisterAppearance = pybind.GenerateRegisterFunc(AppearanceDef)
var NewAppearance = pybind.GenerateConstructor(AppearanceDef).(func() *Appearance)

var AnimationsDef = pybind.Define("Animations", (*Animations)(nil))
var RegisterAnimations = pybind.GenerateRegisterFunc(AnimationsDef)
var NewAnimations = pybind.GenerateConstructor(AnimationsDef).(func() *Animations)
