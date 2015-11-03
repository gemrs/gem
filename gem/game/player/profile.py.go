package player

import (
	"encoding/json"

	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/pybind"
)

var ProfileDef = pybind.Define("Profile", (*Profile)(nil))
var RegisterProfile = pybind.GenerateRegisterFunc(ProfileDef)
var NewProfile = pybind.GenerateConstructor(ProfileDef).(func(string, string) *Profile)

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

func (p *Profile) Py_serialize(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(func() string {
		obj, err := json.Marshal(p)
		if err != nil {
			panic(err)
		}
		return string(obj)
	})
	return fn(args, nil)
}

func (p *Profile) Py_deserialize(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(func(obj string) {
		err := p.UnmarshalJSON([]byte(obj))
		if err != nil {
			panic(err)
		}
	})
	return fn(args, nil)
}

func (p *Profile) PyStr() string {
	return p.String()
}

var SkillsDef = pybind.Define("Skills", (*Skills)(nil))
var RegisterSkills = pybind.GenerateRegisterFunc(SkillsDef)
var NewSkills = pybind.GenerateConstructor(SkillsDef).(func() *Skills)

func (s *Skills) PyGet_combat_level() (py.Object, error) {
	fn := pybind.Wrap(s.CombatLevel)
	return fn(nil, nil)
}
