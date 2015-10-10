package game

import (
	"math/rand"

	"gem/auth"
	"gem/encoding"
	"gem/event"
	"gem/protocol"
)

func (svc *gameService) handshake(conn *Connection, b *encoding.Buffer) error {
	session := conn.Session

	session.RandKey = make([]int32, 4)
	session.RandKey[2] = rand.Int31()
	session.RandKey[3] = rand.Int31()

	handshake := protocol.GameHandshake{}
	if err := handshake.Decode(b, nil); err != nil {
		return err
	}

	conn.write <- &protocol.GameHandshakeResponse{
		ServerISAACSeed: [2]encoding.Int32{
			encoding.Int32(session.RandKey[2]), encoding.Int32(session.RandKey[3]),
		},
	}

	conn.decode = svc.decodeLoginBlock
	return nil
}

func (svc *gameService) decodeLoginBlock(conn *Connection, b *encoding.Buffer) error {
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

func (svc *gameService) decodeSecureBlock(conn *Connection, b *encoding.Buffer) error {
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

	profile, responseCode := svc.auth.LookupProfile(username, password)

	conn.Profile = profile

	if responseCode != auth.AuthOkay {
		conn.write <- &protocol.ServerLoginResponseUnsuccessful{
			Response: encoding.Int8(responseCode),
		}
		conn.Disconnect()
		return nil
	}

	// Successful login, do all the stuff
	conn.write <- &protocol.ServerLoginResponse{
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
