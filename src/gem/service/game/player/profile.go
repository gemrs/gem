package player

import (
	"fmt"
	"github.com/qur/gopy/lib"
)

//go:generate gopygen -type Profile $GOFILE
// Profile represents the saved state of a user
type Profile struct {
	py.BaseObject

	Username string
	Password string /* todo: hash */
}

func (p *Profile) String() string {
	return fmt.Sprintf("Username: %v", p.Username)
}
