package player

import (
	"gem/encoding"
	"gem/protocol"

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

	target encoding.Writer
}

// SendMessage puts a message to the player's chat window
func (session *Session) SendMessage(message string) {
	session.target.WriteEncodable(&protocol.ServerChatMessage{
		Message: encoding.JString(message),
	})
}

// SetTarget directs where packets should be sent.
func (session *Session) SetTarget(e encoding.Writer) {
	session.target = e
}
