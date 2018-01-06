package auth

import (
	"github.com/gemrs/gem/gem/game/interface/player"
)

//go:generate stringer -type=AuthResponse
type AuthResponse int

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
