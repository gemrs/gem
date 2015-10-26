package player

import (
	"math/rand"

	"github.com/gtank/isaac"
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/gem/game/interface/entity"
	"github.com/sinusoids/gem/gem/game/position"
)

//go:generate gopygen -type Session $GOFILE
// Session is the set of non-persistant properties of a player
type Session struct {
	py.BaseObject

	randIn        isaac.ISAAC
	randOut       isaac.ISAAC
	serverRandKey []uint32

	secureBlockSize int

	region         *position.Region
	flags          entity.Flags
	currentWalkDir int
	lastWalkDir    int
}

func (s *Session) Init() error {
	s.serverRandKey = []uint32{
		uint32(rand.Int31()), uint32(rand.Int31()),
	}
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

func (s *Session) ServerISAACSeed() []uint32 {
	return s.serverRandKey
}

func (s *Session) ISAACIn() *isaac.ISAAC {
	return &s.randIn
}

func (s *Session) ISAACOut() *isaac.ISAAC {
	return &s.randOut
}

func (s *Session) InitISAAC(inSeed, outSeed []uint32) {
	s.randIn.SeedArray(inSeed)
	s.randOut.SeedArray(outSeed)
}

func (s *Session) SecureBlockSize() int {
	return s.secureBlockSize
}

func (s *Session) SetSecureBlockSize(size int) {
	s.secureBlockSize = size
}
