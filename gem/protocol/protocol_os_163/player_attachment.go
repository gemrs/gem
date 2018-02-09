package protocol_os_163

import "github.com/gemrs/gem/gem/protocol"

type playerData struct {
	skipCount              int
	skipFlags              [protocol.MaxPlayers]skipBits
	localPlayers           [protocol.MaxPlayers]int
	externalPlayers        [protocol.MaxPlayers]int
	localPlayerCount       int
	externalPlayerCount    int
	frame                  FrameType
	playerIndexInitialized bool
}

func newPlayerData() *playerData {
	return &playerData{
		frame: FixedFrame,
	}
}

func getPlayerData(d interface{}) *playerData {
	return d.(*playerData)
}

type skipBits int

func (s *skipBits) cycle() {
	*s >>= 1
}

func (s *skipBits) updateNextIter() {
	*s |= 0x2
}

func (s *skipBits) shouldUpdate(iter int) bool {
	var secondIter, invert bool
	var result bool

	switch iter {
	case 0:
	case 1:
		secondIter = true
	case 2:
		invert = true
	case 3:
		invert = true
		secondIter = true
	default:
		panic("invalid update iter")
	}

	if secondIter {
		result = *s&0x1 != 0
	} else {
		result = *s&0x1 == 0
	}

	if invert {
		result = !result
	}

	return result
}
