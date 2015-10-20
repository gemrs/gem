package player

import (
	"fmt"

	"github.com/qur/gopy/lib"

	"gem/game/position"
)

type Rights int

const (
	RightsPlayer Rights = iota
	RightsModerator
	RightsAdmin
)

//go:generate gopygen -type Profile -type Skills $GOFILE

// Profile represents the saved state of a user
type Profile struct {
	py.BaseObject

	Username string
	Password string
	Rights   Rights
	Pos      *position.Absolute

	Skills     *Skills
	Appearance *Appearance
	Animations *Animations
}

func (p *Profile) Init() (err error) {
	p.Skills, err = NewSkills()
	if err != nil {
		return err
	}

	p.Appearance, err = NewAppearance()
	if err != nil {
		return err
	}

	p.Animations, err = NewAnimations()
	if err != nil {
		return err
	}

	return nil
}

type Skills struct {
	py.BaseObject

	CombatLevel int
}

func (s *Skills) Init() error {
	return nil
}

func (p *Profile) String() string {
	return fmt.Sprintf("Username: %v", p.Username)
}
