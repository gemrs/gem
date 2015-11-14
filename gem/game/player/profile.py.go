package player

import (
	"encoding/json"

	"github.com/qur/gopy/lib"

	"github.com/gemrs/gem/pybind"
)

var ProfileDef = pybind.Define("Profile", (*Profile)(nil))
var RegisterProfile = pybind.GenerateRegisterFunc(ProfileDef)
var NewProfile = pybind.GenerateConstructor(ProfileDef).(func(string, string) *Profile)

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
