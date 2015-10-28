package player

import (
	"math/rand"

	"github.com/gtank/isaac"
	"github.com/qur/gopy/lib"
)

// Session is the set of non-persistant properties of a player
type Session struct {
	py.BaseObject

	randIn        isaac.ISAAC
	randOut       isaac.ISAAC
	serverRandKey []uint32

	secureBlockSize int
}

func (s *Session) Init() {
	s.serverRandKey = []uint32{
		uint32(rand.Int31()), uint32(rand.Int31()),
	}
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
