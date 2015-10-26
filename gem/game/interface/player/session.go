package player

import (
	"github.com/gtank/isaac"
)

type Session interface {
	ISAACIn() *isaac.ISAAC
	ISAACOut() *isaac.ISAAC
	ServerISAACSeed() []uint32
	InitISAAC(inSeed, outSeed []uint32)
	SecureBlockSize() int
	SetSecureBlockSize(s int)
}
