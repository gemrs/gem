package game

import (
	"github.com/gtank/isaac"
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/game/interface/entity"
	"github.com/sinusoids/gem/gem/game/position"
)

//go:generate gopygen -type Session $GOFILE
// Session is the set of non-persistant properties of a player
type Session struct {
	py.BaseObject

	RandIn  isaac.ISAAC
	RandOut isaac.ISAAC
	RandKey []int32

	SecureBlockSize int

	region         *position.Region
	flags          entity.Flags
	currentWalkDir int
	lastWalkDir    int
}

func (s *Session) Init() error {
	return nil
}

func (s *Session) Flags() entity.Flags {
	return s.flags
}

// SetFlags ORs the given flags with the player's current update flags
func (s *Session) SetFlags(f entity.Flags) {
	s.flags |= f
}

func (s *Session) ClearFlags() {
	s.flags = 0
}

func (s *Session) Region() *position.Region {
	return s.region
}

func (s *Session) SetRegion(r *position.Region) {
	s.region = r
}

func (s *Session) WalkDirection() (current int, last int) {
	return s.currentWalkDir, s.lastWalkDir
}
