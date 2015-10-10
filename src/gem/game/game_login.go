package game

import (
	"math/rand"

	"gem/auth"
	"gem/encoding"
	"gem/event"
	"gem/protocol"
)

// handshake performs the isaac key exchange
func (svc *gameService) handshake(conn *Connection, b *encoding.Buffer) error {
	session := conn.Session

	session.RandKey = make([]int32, 4)
	session.RandKey[2] = rand.Int31()
	session.RandKey[3] = rand.Int31()

	handshake := protocol.InboundGameHandshake{}
	if err := handshake.Decode(b, nil); err != nil {
		return err
	}

	conn.write <- &protocol.OutboundGameHandshake{
		ServerISAACSeed: [2]encoding.Int32{
			encoding.Int32(session.RandKey[2]), encoding.Int32(session.RandKey[3]),
		},
	}

	conn.decode = svc.decodeLoginBlock
	return nil
}

// decodeLoginBlock handles the unencrypted login block
func (svc *gameService) decodeLoginBlock(conn *Connection, b *encoding.Buffer) error {
	session := conn.Session

	loginBlock := protocol.InboundLoginBlock{}
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

// decodeSecureBlock handles the secure login block and the login response (via doLogin)
func (svc *gameService) decodeSecureBlock(conn *Connection, b *encoding.Buffer) error {
	session := conn.Session

	rsaBlock := encoding.RSABlock{&protocol.InboundSecureLoginBlock{}}
	rsaArgs := encoding.RSADecodeArgs{
		Key:       svc.key,
		BlockSize: session.SecureBlockSize,
	}
	if err := rsaBlock.Decode(b, rsaArgs); err != nil {
		return err
	}
	secureBlock := rsaBlock.Codable.(*protocol.InboundSecureLoginBlock)

	// Seed the RNGs
	inSeed := make([]uint32, 4)
	outSeed := make([]uint32, 4)
	for i := range inSeed {
		inSeed[i] = uint32(secureBlock.ISAACSeed[i])
		outSeed[i] = uint32(secureBlock.ISAACSeed[i]) + 50
	}
	session.RandIn.SeedArray(inSeed)
	session.RandOut.SeedArray(outSeed)

	conn.Log.Debugf("Secure login block: %#v", secureBlock)

	username := string(secureBlock.Username)
	password := string(secureBlock.Password)
	password = auth.HashPassword(password)

	return svc.doLogin(conn, username, password)
}

// doLogin authenticates the user, sends the login response, and sets up the client for standard packet processing
func (svc *gameService) doLogin(conn *Connection, username, password string) error {
	profile, responseCode := svc.auth.LookupProfile(username, password)

	conn.Profile = profile

	if responseCode != auth.AuthOkay {
		conn.write <- &protocol.OutboundLoginResponseUnsuccessful{
			Response: encoding.Int8(responseCode),
		}
		conn.Disconnect()
		return nil
	}

	// Successful login, do all the stuff
	conn.write <- &protocol.OutboundLoginResponse{
		Response: encoding.Int8(responseCode),
		Rights:   encoding.Int8(conn.Profile.Rights),
		Flagged:  0,
	}
	conn.decode = svc.decodePacket
	conn.encode = svc.encodePacket

	event.Dispatcher.Raise(event.PlayerLogin, conn)
	go func() {
		<-conn.disconnect
		event.Dispatcher.Raise(event.PlayerLogout, conn)
	}()
	return nil
}
