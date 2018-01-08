//glua:bind module gem.game.player
package player

import (
	"encoding/json"
	"fmt"

	game_event "github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/position"
)

func (player *Player) LoadProfile() {
	profile := player.Profile()
	profile.setPlayer(player)
	player.SetPosition(profile.Position())

	game_event.PlayerLoadProfile.NotifyObservers(player, player.Profile())
}

//go:generate glua .

type Rights int

const (
	RightsPlayer Rights = iota
	RightsModerator
	RightsAdmin
)

// Profile represents the saved state of a user
//glua:bind
type Profile struct {
	username string
	password string
	rights   Rights
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

//glua:bind
func (p *Profile) Password() string {
	return p.password
}

//glua:bind
func (p *Profile) Rights() Rights {
	return p.rights
}

func (p *Profile) setRights(rights Rights) {
	p.rights = rights
}

//glua:bind accessor
func (p *Profile) Position() *position.Absolute {
	return p.position
}

func (p *Profile) SetPosition(pos *position.Absolute) {
	p.position = pos
}

func (p *Profile) Skills() *Skills {
	return p.skills
}

func (p *Profile) Appearance() *Appearance {
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
