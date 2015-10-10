package game

import (
	"io"

	"gem/auth"
	"gem/crypto"
	"gem/encoding"
	"gem/runite"
)

type gameService struct {
	runite *runite.Context
	key    *crypto.Keypair
	auth   auth.Provider
}

func newGameService(runite *runite.Context, key *crypto.Keypair, auth auth.Provider) *gameService {
	return &gameService{
		runite: runite,
		key:    key,
		auth:   auth,
	}
}

func (svc *gameService) encodePacket(ctx *context, b *encoding.Buffer, codable encoding.Encodable) error {
	conn := ctx.conn
	return codable.Encode(conn.writeBuffer, &conn.Session.RandOut)
}

func (svc *gameService) decodePacket(ctx *context, b *encoding.Buffer) error {
	//TODO: Parse packets
	return io.EOF
}
