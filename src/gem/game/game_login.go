package game

import (
	"math/rand"

	"gem/auth"
	"gem/encoding"
	game_event "gem/game/event"
	"gem/protocol"
	game_protocol "gem/protocol/game"
)

// handshake performs the isaac key exchange
func (svc *GameService) handshake(client *Player) error {
	session := client.Session().(*Session)

	session.RandKey = make([]int32, 4)
	session.RandKey[2] = rand.Int31()
	session.RandKey[3] = rand.Int31()

	handshake := protocol.InboundGameHandshake{}
	if err := handshake.Decode(client.Conn().ReadBuffer, nil); err != nil {
		return err
	}

	client.Conn().Write <- &protocol.OutboundGameHandshake{
		ServerISAACSeed: [2]encoding.Int32{
			encoding.Int32(session.RandKey[2]), encoding.Int32(session.RandKey[3]),
		},
	}

	client.decode = svc.decodeLoginBlock
	return nil
}

// decodeLoginBlock handles the unencrypted login block
func (svc *GameService) decodeLoginBlock(client *Player) error {
	session := client.Session().(*Session)

	loginBlock := game_protocol.InboundLoginBlock{}
	if err := loginBlock.Decode(client.Conn().ReadBuffer, nil); err != nil {
		return err
	}

	expectedSecureBlockSize := int(loginBlock.LoginLen) - ((9 * 4) + 1 + 1 + 1 + 2)
	if expectedSecureBlockSize != int(loginBlock.SecureBlockSize) {
		client.Log().Errorf("Secure block size mismatch: got %v expected %v", loginBlock.SecureBlockSize, expectedSecureBlockSize)
		client.Disconnect()
	}

	session.SecureBlockSize = int(loginBlock.SecureBlockSize)

	client.Log().Debugf("Login block: %#v", loginBlock)

	client.decode = svc.decodeSecureBlock
	return nil
}

// decodeSecureBlock handles the secure login block and the login response (via doLogin)
func (svc *GameService) decodeSecureBlock(client *Player) error {
	session := client.Session().(*Session)

	rsaBlock := encoding.RSABlock{&game_protocol.InboundSecureLoginBlock{}}
	rsaArgs := encoding.RSADecodeArgs{
		Key:       svc.key,
		BlockSize: session.SecureBlockSize,
	}
	if err := rsaBlock.Decode(client.Conn().ReadBuffer, rsaArgs); err != nil {
		return err
	}
	secureBlock := rsaBlock.Codable.(*game_protocol.InboundSecureLoginBlock)

	// Seed the RNGs
	inSeed := make([]uint32, 4)
	outSeed := make([]uint32, 4)
	for i := range inSeed {
		inSeed[i] = uint32(secureBlock.ISAACSeed[i])
		outSeed[i] = uint32(secureBlock.ISAACSeed[i]) + 50
	}
	session.RandIn.SeedArray(inSeed)
	session.RandOut.SeedArray(outSeed)

	client.Log().Debugf("Secure login block: %#v", secureBlock)

	username := string(secureBlock.Username)
	password := string(secureBlock.Password)
	password = auth.HashPassword(password)

	return svc.doLogin(client, username, password)
}

// doLogin authenticates the user, sends the login response, and sets up the client for standard packet processing
func (svc *GameService) doLogin(client *Player, username, password string) error {
	profile, responseCode := svc.auth.LookupProfile(username, password)

	if responseCode != auth.AuthOkay {
		client.Conn().Write <- &game_protocol.OutboundLoginResponseUnsuccessful{
			Response: encoding.Int8(responseCode),
		}
		client.Disconnect()
		return nil
	}

	client.profile = profile.(*Profile)

	// Successful login, do all the stuff
	client.Conn().Write <- &game_protocol.OutboundLoginResponse{
		Response: encoding.Int8(responseCode),
		Rights:   encoding.Int8(client.Profile().Rights()),
		Flagged:  0,
	}
	client.decode = svc.decodePacket
	go svc.packetConsumer(client)

	game_event.PlayerLogin.NotifyObservers(client)
	game_event.PlayerLoadProfile.NotifyObservers(client)
	game_event.PlayerFinishLogin.NotifyObservers(client)

	go func() {
		client.WaitForDisconnect()
		game_event.PlayerLogout.NotifyObservers(client)
	}()
	return nil
}
