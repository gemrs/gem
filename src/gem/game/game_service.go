package game

import (
	"fmt"

	"gem/auth"
	"gem/crypto"
	"gem/encoding"
	"gem/protocol"
	"gem/runite"
)

// gameService represents the internal state of the game
type gameService struct {
	runite *runite.Context
	key    *crypto.Keypair
	auth   auth.Provider
}

// newGameService constructs a new gameService
func newGameService(runite *runite.Context, key *crypto.Keypair, auth auth.Provider) *gameService {
	return &gameService{
		runite: runite,
		key:    key,
		auth:   auth,
	}
}

// encodePacket encodes an encoding.Encodable using the ISAAC rand generator
func (svc *gameService) encodePacket(conn *Connection, b *encoding.Buffer, codable encoding.Encodable) error {
	return codable.Encode(conn.writeBuffer, &conn.Session.RandOut)
}

// decodePacket decodes from the readBuffer using the ISAAC rand generator
func (svc *gameService) decodePacket(conn *Connection, b *encoding.Buffer) error {
	data, err := b.Peek(1)
	if err != nil {
		return err
	}

	idByte := int(data[0])

	rand := conn.Session.RandIn.Rand()
	realId := uint8(uint32(idByte) - rand)
	packet, err := protocol.NewInboundPacket(int(realId))
	if err != nil {
		return fmt.Errorf("%v: packet %v", err, realId)
	}
	err = packet.Decode(b, rand)
	if err != nil {
		return err
	}

	conn.read <- packet
	return nil
}

// packetConsumer is the goroutine which picks packets from the readQueue and does something with them
func (svc *gameService) packetConsumer(conn *Connection) {
L:
	for {
		select {
		case <-conn.disconnect:
			break L
		case packet := <-conn.read:
			if _, ok := packet.(*protocol.UnknownPacket); ok {
				/* unknown packet; dump to the log */
				conn.Log.Debugf("Got unknown packet: %v", packet)
				continue
			}
			// TODO: route known packets to a handler
			conn.Log.Debugf("Got known packet %T", packet)
		}

	}
}
