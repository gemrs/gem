package game

import (
	"math/rand"

	"gem/auth"
	"gem/crypto"
	"gem/encoding"
	"gem/event"
	"gem/protocol"
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
	session := conn.Session

	loginBlock := protocol.ClientLoginBlock{}
	if err := loginBlock.Decode(b, nil); err != nil {
		return err
	}

	expectedSecureBlockSize := int(loginBlock.LoginLen) - ((9 * 4) + 1 + 1 + 1 + 2)
	if expectedSecureBlockSize != int(loginBlock.SecureBlockSize) {
		conn.Log.Errorf("Secure block size mismatch: got %v expected %v", loginBlock.SecureBlockSize, expectedSecureBlockSize)
		conn.Disconnect()
	}

	session.SecureBlockSize = int(loginBlock.SecureBlockSize)

	conn.Log.Debugf("Login block: %#v", loginBlock)

	conn.decode = svc.decodeSecureBlock
	return nil
}

func (svc *gameService) decodeSecureBlock(ctx *context, b *encoding.Buffer) error {
	conn := ctx.conn
	session := conn.Session

	rsaBlock := encoding.RSABlock{&protocol.ClientSecureLoginBlock{}}
	rsaArgs := encoding.RSADecodeArgs{
		Key:       svc.key,
		BlockSize: session.SecureBlockSize,
	}
	if err := rsaBlock.Decode(b, rsaArgs); err != nil {
		return err
	}
	secureBlock := rsaBlock.Codable.(*protocol.ClientSecureLoginBlock)

	conn.Log.Debugf("Secure login block: %#v", secureBlock)

	profile, responseCode := svc.auth.LookupProfile(string(secureBlock.Username), string(secureBlock.Password))

	conn.Profile = profile

	if responseCode == auth.AuthOkay {
		response := protocol.ServerLoginResponse{
			Response: encoding.Int8(responseCode),
			Rights:   encoding.Int8(conn.Profile.Rights),
			Flagged:  0,
		}
		if err := response.Encode(conn, nil); err != nil {
			return err
		}
		conn.canWrite <- 1
		conn.decode = svc.decodePacket
	} else {
		response := protocol.ServerLoginResponseUnsuccessful{
			Response: encoding.Int8(responseCode),
		}
		if err := response.Encode(conn, nil); err != nil {
			return err
		}
		conn.canWrite <- 1
		conn.Disconnect()
	}

	event.Dispatcher.Raise(event.PlayerLogin, conn)
	return nil
}

func (svc *gameService) decodePacket(ctx *context, b *encoding.Buffer) error {
	//TODO: Parse packets
	return nil
}
