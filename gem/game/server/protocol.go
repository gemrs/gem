package server

import (
	"github.com/gemrs/gem/fork/github.com/gtank/isaac"
	"github.com/gemrs/gem/gem/core/crypto"
	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/runite"
	"github.com/gemrs/willow/log"
)

var Proto Protocol

func SetProtocolImpl(p Protocol) {
	Proto = p
}

type Protocol interface {
	Encode(message Message) encoding.Encodable
	Decode(message encoding.Decodable) Message
	NewLoginHandler() LoginHandler
	Handshake(conn *Connection) int
	NewInboundPacket(id int) (encoding.Decodable, error)
	NewUpdateService(runite *runite.Context) Service
	GameServiceId() int
	UpdateServiceId() int
}

type Message interface{}

type LoginHandler interface {
	SetRsaKeypair(keypair *crypto.Keypair)
	SetServerIsaacSeed(seed []uint32)
	SetCompleteCallback(cb func(LoginHandler) error)
	Perform(GameClient)

	/* results */
	Username() string
	Password() string
	IsaacSeeds() (in []uint32, out []uint32)
}

// DecodeFunc is used for parsing the read stream and dealing with the incoming data.
// If io.EOF is returned, it is assumed that we didn't have enough data, and
// the underlying buffer's read pointer is not altered.
type DecodeFunc func(GameClient) error

type GameClient interface {
	Conn() *Connection
	Decode() error
	Send(encoding.Encodable) error
	Disconnect()

	SetProtoData(d interface{})
	SetDecodeFunc(d DecodeFunc)
	Log() log.Log
	IsaacIn() *isaac.ISAAC
}
