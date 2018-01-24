package protocol_317

import (
	"github.com/gemrs/gem/gem/core/crypto"
	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/game/server"
)

func (p protocolImpl) Handshake(conn *server.Connection) int {
	var serviceSelect InboundServiceSelect
	serviceSelect.Decode(conn.NetConn(), nil)

	return int(serviceSelect.Service)
}

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
	client.SetDecodeFunc(handler.handshake)
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

// handshake performs the isaac key exchange
func (handler *LoginHandler) handshake(client server.GameClient) error {
	serverSeed := handler.serverIsaacSeed

	handshake := InboundGameHandshake{}
	handshake.Decode(client.Conn().ReadBuffer, nil)

	client.Conn().Write <- &OutboundGameHandshake{
		ServerISAACSeed: [2]encoding.Uint32{
			encoding.Uint32(serverSeed[0]), encoding.Uint32(serverSeed[1]),
		},
	}

	client.SetDecodeFunc(handler.decodeLoginBlock)
	return nil
}

// decodeLoginBlock handles the unencrypted login block
func (handler *LoginHandler) decodeLoginBlock(client server.GameClient) error {
	loginBlock := InboundLoginBlock{}
	loginBlock.Decode(client.Conn().ReadBuffer, nil)

	expectedSecureBlockSize := int(loginBlock.LoginLen) - ((9 * 4) + 1 + 1 + 1 + 2)
	if expectedSecureBlockSize != int(loginBlock.SecureBlockSize) {
		client.Log().Error("Secure block size mismatch: got %v expected %v", loginBlock.SecureBlockSize, expectedSecureBlockSize)
		client.Conn().Disconnect()
	}

	handler.secureBlockSize = int(loginBlock.SecureBlockSize)

	client.SetDecodeFunc(handler.decodeSecureBlock)
	return nil
}

// decodeSecureBlock handles the secure login block and the login response (via doLogin)
func (handler *LoginHandler) decodeSecureBlock(client server.GameClient) error {
	// Decode the RSA block
	rsaBlock := encoding.RSABlock{
		Codable: &InboundSecureLoginBlock{},
	}
	rsaArgs := encoding.RSADecodeArgs{
		Key:       handler.rsaKeypair,
		BlockSize: handler.secureBlockSize,
	}
	rsaBlock.Decode(client.Conn().ReadBuffer, rsaArgs)

	secureBlock := rsaBlock.Codable.(*InboundSecureLoginBlock)

	// Seed the RNGs
	inSeed := make([]uint32, 4)
	outSeed := make([]uint32, 4)
	for i := range inSeed {
		inSeed[i] = uint32(secureBlock.ISAACSeed[i])
		outSeed[i] = uint32(secureBlock.ISAACSeed[i]) + 50
	}

	// Store results and pass back to the caller
	handler.inSeed = inSeed
	handler.outSeed = outSeed

	handler.username = string(secureBlock.Username)
	handler.password = string(secureBlock.Password)

	return handler.callback(handler)
}
