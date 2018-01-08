//glua:bind module gem.game.player
package player

import (
	"encoding/json"
	"fmt"

	"github.com/gemrs/gem/gem/game/interface/player"
	"github.com/gemrs/gem/gem/game/position"
)

//go:generate glua .

// Profile represents the saved state of a user
//glua:bind
type Profile struct {
	username string
	password string
	rights   player.Rights
	position *position.Absolute

	skills     *Skills
	appearance *Appearance
}

//glua:bind constructor Profile
func NewProfile(username, password string) *Profile {
	return &Profile{
		username:   username,
		password:   password,
		skills:     NewSkills(),
		appearance: NewAppearance(),
	}
}

func (p *Profile) setPlayer(player *Player) {
	p.skills.setPlayer(player)
	p.appearance.setPlayer(player)
}

//glua:bind
func (p *Profile) Username() string {
	return p.username
}

func (p *Profile) setUsername(username string) {
	p.username = username
}

//glua:bind
func (p *Profile) Password() string {
	return p.password
}

func (p *Profile) setPassword(password string) {
	p.password = password
}

//glua:bind
func (p *Profile) Rights() player.Rights {
	return p.rights
}

func (p *Profile) setRights(rights player.Rights) {
	p.rights = rights
}

//glua:bind accessor
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
