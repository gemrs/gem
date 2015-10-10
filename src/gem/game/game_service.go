package game

import (
	"io"

	"gem/auth"
	"gem/crypto"
	"gem/encoding"
	"gem/runite"
)

// gameService represents the internal state of the game
type gameService struct {
	runite *runite.Context
	key    *crypto.Keypair
	auth   auth.Provider
}

// newGameService constructs a new gameService
func newGameService(runite *runite.Context, key *crypto.Keypair, auth auth.Provider) *gameService {
	return &gameService{
		runite: runite,
		key:    key,
		auth:   auth,
	}
}

// encodePacket encodes an encoding.Encodable using the ISAAC rand generator
func (svc *gameService) encodePacket(conn *Connection, b *encoding.Buffer, codable encoding.Encodable) error {
	return codable.Encode(conn.writeBuffer, &conn.Session.RandOut)
}

// decodePacket decodes from the readBuffer using the ISAAC rand generator
func (svc *gameService) decodePacket(conn *Connection, b *encoding.Buffer) error {

	return io.EOF
}
