package player

import (
	"encoding/json"
	"fmt"

	"github.com/gemrs/gem/gem/game/interface/player"
	"github.com/gemrs/gem/gem/game/position"
)

// Profile represents the saved state of a user
type Profile struct {
	username string
	password string
	rights   player.Rights
	position *position.Absolute

	skills     *Skills
	appearance *Appearance
}

func (p *Profile) Init(username, password string) {
	p.username = username
	p.password = password

	p.skills = NewSkills()
	p.appearance = NewAppearance()
}

func (p *Profile) Username() string {
	return p.username
}

func (p *Profile) setUsername(username string) {
	p.username = username
}

func (p *Profile) Password() string {
	return p.password
}

func (p *Profile) setPassword(password string) {
	p.password = password
}

func (p *Profile) Rights() player.Rights {
	return p.rights
}

func (p *Profile) setRights(rights player.Rights) {
	p.rights = rights
}

func (p *Profile) Position() *position.Absolute {
	return p.position
}

func (p *Profile) SetPosition(pos *position.Absolute) {
	p.position = pos
}

func (p *Profile) Skills() player.Skills {
	return p.skills
}

func (p *Profile) Appearance() player.Appearance {
	return p.appearance
}

func (p *Profile) SetAppearance(appearance player.Appearance) {
	p.appearance = appearance.(*Appearance)
}

func (p *Profile) String() string {
	return fmt.Sprintf("Username: %v", p.username)
}

func (p *Profile) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonObjForProfile(p))
}

func (p *Profile) UnmarshalJSON(objJSON []byte) error {
	var deserialized jsonProfile
	err := json.Unmarshal(objJSON, &deserialized)
	if err != nil {
		return err
	}

	jsonObjToProfile(p, deserialized)
	return nil
}

type Skills struct {
	combatLevel int
}

func (s *Skills) Init() {}

func (s *Skills) CombatLevel() int {
	return s.combatLevel
}

func (s *Skills) setCombatLevel(combatLevel int) {
	s.combatLevel = combatLevel
}
