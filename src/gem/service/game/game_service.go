package game

import (
	"math/rand"

	"gem/auth"
	"gem/crypto"
	"gem/encoding"
	"gem/protocol"
	"gem/runite"
)

type gameService struct {
	runite *runite.Context
	key    *crypto.Keypair
}

func newGameService(runite *runite.Context, key *crypto.Keypair) *gameService {
	return &gameService{
		runite: runite,
		key:    key,
	}
}

func (svc *gameService) handshake(ctx *context, b *encoding.Buffer) error {
	conn := ctx.conn
	session := conn.Session

	session.RandKey = make([]int32, 4)
	session.RandKey[2] = rand.Int31()
	session.RandKey[3] = rand.Int31()

	handshake := protocol.GameHandshake{}
	if err := handshake.Decode(b, nil); err != nil {
		return err
	}

	response := &protocol.GameHandshakeResponse{
		ServerISAACSeed: [2]encoding.Int32{
			encoding.Int32(session.RandKey[2]), encoding.Int32(session.RandKey[3]),
		},
	}

	if err := response.Encode(conn, nil); err != nil {
		return err
	}
	conn.canWrite <- 1
	conn.decode = svc.decodeLoginBlock
	return nil
}

func (svc *gameService) decodeLoginBlock(ctx *context, b *encoding.Buffer) error {
	conn := ctx.conn

	loginBlock := protocol.ClientLoginBlock{}
	if err := loginBlock.Decode(b, nil); err != nil {
		return err
	}

	conn.Log.Debugf("Login block: %#v", loginBlock)

	//TODO: Parse encrypted block

	response := protocol.ServerLoginResponse{
		Response: encoding.Int8(auth.AuthOkay),
		Rights:   0,
		Flagged:  0,
	}
	if err := response.Encode(conn, nil); err != nil {
		return err
	}
	conn.canWrite <- 1
	conn.decode = svc.decodePacket
	return nil
}

func (svc *gameService) decodePacket(ctx *context, b *encoding.Buffer) error {
	//TODO: Parse packets
	return nil
}
