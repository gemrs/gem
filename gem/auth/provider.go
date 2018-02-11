//glua:bind module gem.game.auth
package auth

import (
	"github.com/gemrs/gem/gem/protocol"
)

// Provider is a provider of authorization.
type Provider interface {
	LoadProfile(name, password string) (protocol.Profile, protocol.AuthResponse)
	SaveProfile(profile protocol.Profile)
}
