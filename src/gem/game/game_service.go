package game

import (
	"fmt"

	"gem/auth"
	"gem/crypto"
	"gem/event"
	"gem/game/packet"
	"gem/game/server"
	"gem/protocol"
	"gem/runite"

	"github.com/qur/gopy/lib"
)

//go:generate gopygen -type GameService -excfield "^[a-z].*" $GOFILE

type GameServiceIface interface {
	GetRunite() *runite.Context
}

// gameService represents the internal state of the game
type GameService struct {
	py.BaseObject

	runite *runite.Context
	key    *crypto.Keypair
	auth   auth.Provider
}

func (svc *GameService) Init(runite *runite.Context, rsaKeyPath string, auth auth.Provider) error {
	var err error
	var key *crypto.Keypair
	key, err = crypto.LoadPrivateKey(rsaKeyPath)
	if err != nil {
		return err
	}

	svc.runite = runite
	svc.key = key
	svc.auth = auth

	PlayerFinishLoginEvent.Register(event.NewListener(finishLogin))
	return nil
}

func (svc *GameService) NewClient(conn *server.Connection, service int) server.Client {
	conn.Log().Infof("new game client")
	client, err := NewPlayer(conn, svc)
	if err != nil {
		panic(err)
	}
	return client
}

// decodePacket decodes from the readBuffer using the ISAAC rand generator
func (svc *GameService) decodePacket(client *Player) error {
	b := client.Conn().ReadBuffer
	data, err := b.Peek(1)
	if err != nil {
		return err
	}

	idByte := int(data[0])

	rand := client.Session().RandIn.Rand()
	realId := uint8(uint32(idByte) - rand)
	packet, err := protocol.NewInboundPacket(int(realId))
	if err != nil {
		return fmt.Errorf("%v: packet %v", err, realId)
	}
	err = packet.Decode(b, rand)
	if err != nil {
		return err
	}

	if !client.IsDisconnecting() {
		client.Conn().Read <- packet
	}
	return nil
}

// packetConsumer is the goroutine which picks packets from the readQueue and does something with them
func (svc *GameService) packetConsumer(client *Player) {
L:
	for {
		select {
		case <-client.Conn().DisconnectChan:
			break L
		case pkt := <-client.Conn().Read:
			if _, ok := pkt.(*protocol.UnknownPacket); ok {
				/* unknown packet; dump to the log */
				client.Log().Debugf("Got unknown packet: %v", pkt)
				continue
			}
			packet.Dispatch(client, pkt)
		}

	}
}
