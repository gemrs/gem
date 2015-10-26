package player

import (
	"github.com/gtank/isaac"

	"github.com/sinusoids/gem/gem/game/interface/entity"
	"github.com/sinusoids/gem/gem/game/position"
)

type Session interface {
	Flags() entity.Flags
	Region() *position.Region
	WalkDirection() (current int, last int)
	ISAACIn() *isaac.ISAAC
	ISAACOut() *isaac.ISAAC
	ServerISAACSeed() []uint32
	InitISAAC(inSeed, outSeed []uint32)
	SecureBlockSize() int
	SetSecureBlockSize(s int)
}
