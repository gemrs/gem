package protocol_os_157

import (
	"github.com/gemrs/gem/gem/core/crypto"
	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol"
)

func (p protocolImpl) NewLoginHandler() server.LoginHandler {
	return &LoginHandler{}
}

type LoginHandler struct {
	serverIsaacSeed []uint32
	callback        func(server.LoginHandler) error
	secureBlockSize int
	rsaKeypair      *crypto.Keypair
	username        string
	password        string
	inSeed, outSeed []uint32
}

func (handler *LoginHandler) Perform(client server.GameClient) {
	serverSeed := handler.serverIsaacSeed

	client.Send(&OutboundGameHandshake{
		ServerISAACSeed: [2]encoding.Uint32{
			encoding.Uint32(serverSeed[0]), encoding.Uint32(serverSeed[1]),
		},
	})

	player := client.(protocol.Player)
	playerData := newPlayerData()
	player.SetProtoData(playerData)
	player.SetCurrentFrame(playerData.frame)
	player.SetDecodeFunc(handler.decodeLoginBlock)
}

func (handler *LoginHandler) SetRsaKeypair(keypair *crypto.Keypair) {
	handler.rsaKeypair = keypair
}

func (handler *LoginHandler) SetServerIsaacSeed(seed []uint32) {
	handler.serverIsaacSeed = seed
}

func (handler *LoginHandler) IsaacSeeds() (in []uint32, out []uint32) {
	return handler.inSeed, handler.outSeed
}

func (handler *LoginHandler) SetCompleteCallback(cb func(server.LoginHandler) error) {
	handler.callback = cb
}

func (handler *LoginHandler) Username() string {
	return handler.username
}

func (handler *LoginHandler) Password() string {
	return handler.password
}

// decodeLoginBlock handles the unencrypted login block
func (handler *LoginHandler) decodeLoginBlock(client server.GameClient) error {
	loginBlock := InboundLoginBlock{}
	loginBlock.Decode(client.Conn().ReadBuffer, nil)
	/*
		expectedSecureBlockSize := int(loginBlock.LoginLen) - ((9 * 4) + 1 + 1 + 1 + 2)
		if expectedSecureBlockSize != int(loginBlock.SecureBlockSize) {
			client.Log().Error("Secure block size mismatch: got %v expected %v", loginBlock.SecureBlockSize, expectedSecureBlockSize)
			client.Conn().Disconnect()
		}
	*/
	handler.secureBlockSize = int(loginBlock.SecureBlockSize)

	client.SetDecodeFunc(handler.decodeSecureBlock)
	return nil
}

// decodeSecureBlock handles the secure login block and the login response (via doLogin)
func (handler *LoginHandler) decodeSecureBlock(client server.GameClient) error {
	// Decode the RSA block
	rsaBlock := encoding.RSABlock{
		Codable: &InboundRsaLoginBlock{},
	}
	rsaArgs := encoding.RSADecodeArgs{
		Key:       handler.rsaKeypair,
		BlockSize: handler.secureBlockSize,
	}
	rsaBlock.Decode(client.Conn().ReadBuffer, rsaArgs)

	secureBlock1 := rsaBlock.Codable.(*InboundRsaLoginBlock)

	// Seed the RNGs
	inSeed := make([]uint32, 4)
	outSeed := make([]uint32, 4)
	for i := range inSeed {
		inSeed[i] = uint32(secureBlock1.ISAACSeed[i])
		outSeed[i] = uint32(secureBlock1.ISAACSeed[i]) + 50
	}

	// Decode XTEA block
	xteaBlock := encoding.XTEABlock{&InboundXteaLoginBlock{}}
	xteaArgs := encoding.XTEADecodeArgs{}
	copy(xteaArgs.Key[0:4], inSeed)
	xteaBlock.Decode(client.Conn().ReadBuffer, xteaArgs)
	secureBlock2 := xteaBlock.Codable.(*InboundXteaLoginBlock)

	// Store results and pass back to the caller
	handler.inSeed = inSeed
	handler.outSeed = outSeed

	handler.username = string(secureBlock2.Username)
	handler.password = string(secureBlock1.Password)

	return handler.callback(handler)
}
