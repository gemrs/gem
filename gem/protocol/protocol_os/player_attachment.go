package protocol_os

import "github.com/gemrs/gem/gem/protocol"

type PlayerData struct {
	skipCount              int
	skipFlags              [protocol.MaxPlayers]skipBits
	LocalPlayers           [protocol.MaxPlayers]int
	ExternalPlayers        [protocol.MaxPlayers]int
	LocalPlayerCount       int
	ExternalPlayerCount    int
	Frame                  FrameType
	PlayerIndexInitialized bool
}

func NewPlayerData() *PlayerData {
	return &PlayerData{
		Frame: FixedFrame,
	}
}

func GetPlayerData(d interface{}) *PlayerData {
	return d.(*PlayerData)
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
