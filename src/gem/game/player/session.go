package player

import (
	"github.com/gtank/isaac"
	"github.com/qur/gopy/lib"
)

//go:generate gopygen -type Session $GOFILE
// Session is the set of non-persistant properties of a player
type Session struct {
	py.BaseObject

	RandIn  isaac.ISAAC
	RandOut isaac.ISAAC
	RandKey []int32

	SecureBlockSize int
}

func (s *Session) Init() error {
	return nil
}
