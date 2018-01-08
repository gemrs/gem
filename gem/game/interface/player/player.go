package player

import (
	"github.com/gemrs/gem/fork/github.com/gtank/isaac"

	"github.com/gemrs/gem/gem/game/interface/entity"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/willow/log"
)

// DecodeFunc is used for parsing the read stream and dealing with the incoming data.
// If io.EOF is returned, it is assumed that we didn't have enough data, and
// the underlying buffer's read pointer is not altered.
type DecodeFunc func(Player) error

// Player is an Entity representing a player
type Player interface {
	entity.Movable
	Profile() Profile
	Animations() Animations
	SetProfile(p Profile)
	Log() log.Log
	FinishInit()
	LoadProfile()
	Cleanup()
	LoadedRegion() *position.Region
	VisibleEntities() *entity.Collection
	ClientConfig() ClientConfig

	ChatMessageQueue() []*ChatMessage
	AppendChatMessage(*ChatMessage)

	SetDecodeFunc(d DecodeFunc)
	Conn() *server.Connection

	InitISAAC(inSeed, outSeed []uint32)
	ISAACIn() *isaac.ISAAC
	ISAACOut() *isaac.ISAAC
	ServerISAACSeed() []uint32
	SecureBlockSize() int
	SetSecureBlockSize(s int)
}

type ChatMessage struct {
	Effects       uint8
	Colour        uint8
	Message       string
	PackedMessage []byte
}
