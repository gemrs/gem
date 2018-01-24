//glua:bind module gem.game.protocol
package protocol

//go:generate glua .

//go:generate stringer -type=AuthResponse
//glua:bind
type AuthResponse int

//glua:bind constructor AuthResponse
func NewAuthResponse(i int) *AuthResponse {
	x := AuthResponse(i)
	return &x
}

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

//glua:bind
type Rights int

func NewRights(i int) *Rights {
	x := Rights(i)
	return &x
}

//glua:bind
const (
	RightsPlayer Rights = iota
	RightsModerator
	RightsAdmin
)
