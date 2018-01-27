//glua:bind module gem.game.auth
package auth

import (
	"github.com/gemrs/gem/gem/protocol"
)

// Provider is a provider of authorization.
type Provider interface {
	LookupProfile(name, password string) (protocol.Profile, protocol.AuthResponse)
}
