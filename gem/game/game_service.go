package game

import (
	"fmt"

	"github.com/sinusoids/gem/gem/auth"
	"github.com/sinusoids/gem/gem/crypto"
	"github.com/sinusoids/gem/gem/event"
	game_event "github.com/sinusoids/gem/gem/game/event"
	"github.com/sinusoids/gem/gem/game/interface/entity"
	"github.com/sinusoids/gem/gem/game/interface/player"
	"github.com/sinusoids/gem/gem/game/packet"
	playerimpl "github.com/sinusoids/gem/gem/game/player"
	"github.com/sinusoids/gem/gem/game/server"
	game_protocol "github.com/sinusoids/gem/gem/protocol/game"
	"github.com/sinusoids/gem/gem/runite"
	"github.com/sinusoids/gem/gem/util/expire"

	"github.com/qur/gopy/lib"
)

// GameService represents the internal state of the game
type GameService struct {
	py.BaseObject
	expire.NonExpirable

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
	svc.NonExpirable = expire.NewNonExpirable()

	game_event.PlayerFinishLogin.Register(event.NewListener(svc, playerFinishLogin))
	game_event.PlayerLogout.Register(event.NewListener(svc, playerCleanup))
	game_event.EntityRegionChange.Register(event.NewListener(svc, svc.EntityUpdate))
	game_event.EntitySectorChange.Register(event.NewListener(svc, svc.EntityUpdate))
	game_event.PlayerAppearanceUpdate.Register(event.NewListener(svc, svc.PlayerUpdate))
	return nil
}

// playerFinishLogin calls player.FinishInit on the PlayerFinishLogin event
func playerFinishLogin(_ *event.Event, args ...interface{}) {
	client := args[0].(player.Player)
	client.FinishInit()
}

// playerCleanup calls player.Cleanup on the PlayerLogout event
func playerCleanup(_ *event.Event, args ...interface{}) {
	client := args[0].(player.Player)
	client.Cleanup()
}

func (svc *GameService) NewClient(conn *server.Connection, service int) server.Client {
	conn.Log().Infof("new game client")
	client := playerimpl.NewPlayer(conn)
	client.SetDecodeFunc(svc.handshake)
	return client
}

func (svc *GameService) EntityUpdate(ev *event.Event, _args ...interface{}) {
	if len(_args) < 1 {
		panic("invalid args length")
	}

	args := _args[0].(map[string]interface{})
	entity := args["entity"].(entity.Entity)
	switch ev {
	case game_event.EntityRegionChange:
		entity.RegionChange()
	case game_event.EntitySectorChange:
		entity.SectorChange()
	}
}

func (svc *GameService) PlayerUpdate(ev *event.Event, _args ...interface{}) {
	if len(_args) < 1 {
		panic("invalid args length")
	}

	args := _args[0].(map[string]interface{})
	player := args["entity"].(player.Player)
	switch ev {
	case game_event.PlayerAppearanceUpdate:
		player.AppearanceChange()
	}
}

// decodePacket decodes from the readBuffer using the ISAAC rand generator
func (svc *GameService) decodePacket(client player.Player) error {
	b := client.Conn().ReadBuffer
	data, err := b.Peek(1)
	if err != nil {
		return err
	}

	idByte := int(data[0])

	rand := client.ISAACIn().Rand()
	realId := uint8(uint32(idByte) - rand)
	packet, err := game_protocol.NewInboundPacket(int(realId))
	if err != nil {
		return fmt.Errorf("%v: packet %v", err, realId)
	}
	err = packet.Decode(b, rand)
	if err != nil {
		return err
	}

	if !client.Conn().IsDisconnecting() {
		client.Conn().Read <- packet
	}
	return nil
}

// packetConsumer is the goroutine which picks packets from the readQueue and does something with them
func (svc *GameService) packetConsumer(client player.Player) {
L:
	for {
		select {
		case <-client.Conn().DisconnectChan:
			break L
		case pkt := <-client.Conn().Read:
			if _, ok := pkt.(*game_protocol.UnknownPacket); ok {
				/* unknown packet; dump to the log */
				client.Log().Debugf("Got unknown packet: %v", pkt)
				continue
			}
			packet.Dispatch(client, pkt)
		}

	}
}
