package player

import (
	"github.com/sinusoids/gem/gem/game/interface/entity"
	"github.com/sinusoids/gem/gem/game/server"
	"github.com/sinusoids/gem/gem/log"
)

// DecodeFunc is used for parsing the read stream and dealing with the incoming data.
// If io.EOF is returned, it is assumed that we didn't have enough data, and
// the underlying buffer's read pointer is not altered.
type DecodeFunc func(Player) error

// Player is an Entity representing a player
type Player interface {
	entity.Movable
	Profile() Profile
	SetProfile(p Profile)
	Session() Session
	Log() *log.Module
	SetDecodeFunc(d DecodeFunc)
	Conn() *server.Connection
	FinishInit()
	LoadProfile()
}
