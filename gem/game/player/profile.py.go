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

func (p *Profile) PyGet_position() (py.Object, error) {
	fn := pybind.Wrap(p.Position)
	return fn(nil, nil)
}

func (p *Profile) PySet_position(value py.Object) error {
	fn := pybind.Wrap(p.SetPosition)
	args, err := py.PackTuple(value)
	if err != nil {
		return err
	}
	_, err = fn(args, nil)
	return err
}

func (p *Profile) PyGet_appearance() (py.Object, error) {
	fn := pybind.Wrap(p.Appearance)
	return fn(nil, nil)
}

func (p *Profile) PyGet_skills() (py.Object, error) {
	fn := pybind.Wrap(p.Skills)
	return fn(nil, nil)
}

func (p *Profile) PySet_appearance(value py.Object) error {
	fn := pybind.Wrap(p.SetAppearance)
	args, err := py.PackTuple(value)
	if err != nil {
		return err
	}
	_, err = fn(args, nil)
	return err
}

func (p *Profile) PyStr() string {
	return p.String()
}
