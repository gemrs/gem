package player

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/pybind"
)

var ProfileDef = pybind.Define("Profile", (*Profile)(nil))
var RegisterProfile = pybind.GenerateRegisterFunc(ProfileDef)
var NewProfile = pybind.GenerateConstructor(ProfileDef).(func(string, string) *Profile)

var SkillsDef = pybind.Define("Skills", (*Skills)(nil))
var RegisterSkills = pybind.GenerateRegisterFunc(SkillsDef)
var NewSkills = pybind.GenerateConstructor(SkillsDef).(func() *Skills)

func (p *Profile) Py_Position(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(p.Position)
	return fn(args, kwds)
}

func (p *Profile) Py_SetPosition(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(p.SetPosition)
	return fn(args, kwds)
}

func (p *Profile) Py_Appearance(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(p.Appearance)
	return fn(args, kwds)
}

func (p *Profile) Py_SetAppearance(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(p.SetAppearance)
	return fn(args, kwds)
}

func (p *Profile) PyStr() string {
	return p.String()
}
