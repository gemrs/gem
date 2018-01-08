package game

import (
	"github.com/gemrs/gem/gem/auth"
	"github.com/gemrs/gem/gem/encoding"
	game_event "github.com/gemrs/gem/gem/game/event"
	playeriface "github.com/gemrs/gem/gem/game/interface/player"
	"github.com/gemrs/gem/gem/game/player"
	"github.com/gemrs/gem/gem/protocol"
	game_protocol "github.com/gemrs/gem/gem/protocol/game"
)

// handshake performs the isaac key exchange
func (svc *GameService) handshake(client *player.Player) error {
	serverSeed := client.ServerISAACSeed()

	handshake := protocol.InboundGameHandshake{}
	if err := handshake.Decode(client.Conn().ReadBuffer, nil); err != nil {
		return err
	}

	client.Conn().Write <- &protocol.OutboundGameHandshake{
		ServerISAACSeed: [2]encoding.Uint32{
			encoding.Uint32(serverSeed[0]), encoding.Uint32(serverSeed[1]),
		},
	}

	client.SetDecodeFunc(svc.decodeLoginBlock)
	return nil
}

// decodeLoginBlock handles the unencrypted login block
func (svc *GameService) decodeLoginBlock(client *player.Player) error {
	loginBlock := game_protocol.InboundLoginBlock{}
	if err := loginBlock.Decode(client.Conn().ReadBuffer, nil); err != nil {
		return err
	}

	expectedSecureBlockSize := int(loginBlock.LoginLen) - ((9 * 4) + 1 + 1 + 1 + 2)
	if expectedSecureBlockSize != int(loginBlock.SecureBlockSize) {
		client.Log().Error("Secure block size mismatch: got %v expected %v", loginBlock.SecureBlockSize, expectedSecureBlockSize)
		client.Conn().Disconnect()
	}

	client.SetSecureBlockSize(int(loginBlock.SecureBlockSize))

	client.SetDecodeFunc(svc.decodeSecureBlock)
	return nil
}

// decodeSecureBlock handles the secure login block and the login response (via doLogin)
func (svc *GameService) decodeSecureBlock(client *player.Player) error {
	rsaBlock := encoding.RSABlock{&game_protocol.InboundSecureLoginBlock{}}
	rsaArgs := encoding.RSADecodeArgs{
		Key:       svc.key,
		BlockSize: client.SecureBlockSize(),
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
	client.InitISAAC(inSeed, outSeed)

	username := string(secureBlock.Username)
	password := string(secureBlock.Password)
	password = auth.HashPassword(password)

	return svc.doLogin(client, username, password)
}

// doLogin authenticates the user, sends the login response, and sets up the client for standard packet processing
func (svc *GameService) doLogin(client *player.Player, username, password string) error {
	profile, responseCode := svc.auth.LookupProfile(username, password)

	if responseCode != auth.AuthOkay {
		client.Conn().Write <- &game_protocol.OutboundLoginResponseUnsuccessful{
			Response: encoding.Uint8(responseCode),
		}
		return nil
	}

	client.SetProfile(profile.(playeriface.Profile))

	// Successful login, do all the stuff
	client.Conn().Write <- &game_protocol.OutboundLoginResponse{
		Response: encoding.Uint8(responseCode),
		Rights:   encoding.Uint8(client.Profile().Rights()),
		Flagged:  0,
	}
	client.SetDecodeFunc(svc.decodePacket)
	go svc.packetConsumer(client)

	client.LoadProfile()
	client.FinishInit()
	game_event.PlayerLogin.NotifyObservers(client)

	go func() {
		client.Conn().WaitForDisconnect()
		game_event.PlayerLogout.NotifyObservers(client)
	}()
	return nil
}
