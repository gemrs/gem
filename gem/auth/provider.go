//glua:bind module gem.game.auth
package auth

import (
	"github.com/gemrs/gem/gem/game/interface/player"
)

//go:generate stringer -type=AuthResponse
//glua:bind
type AuthResponse int

//glua:bind constructor AuthResponse
func NewAuthResponse(i int) *AuthResponse {
	x := AuthResponse(i)
	return &x
}

//go:generate glua .

//glua:bind
const (
	AuthPending AuthResponse = iota
	AuthDelay
	AuthOkay
	AuthInvalidCredentials
	AuthDisabled
	AuthDuplicateSession
	AuthUpdates
	AuthServerFull
	AuthNoLoginServer
	AuthTooManyConnections
	AuthBadSessionId
	AuthRejected
	AuthMembersWorld
	AuthIncomplete
	AuthUpdating
	AuthUnknown
	AuthAttemptsExceeded
	AuthMembersArea
	_
	_
	AuthInvalidLoginServer
	AuthInvalidTransferring
	AuthEnd
)

// Provider is a provider of authorization.
type Provider interface {
	LookupProfile(name, password string) (player.Profile, AuthResponse)
}
