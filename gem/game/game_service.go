//glua:bind module gem.game
package game

import (
	"fmt"

	"github.com/gemrs/gem/gem/auth"
	"github.com/gemrs/gem/gem/crypto"
	engine_event "github.com/gemrs/gem/gem/engine/event"
	"github.com/gemrs/gem/gem/event"
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/packet"
	"github.com/gemrs/gem/gem/game/player"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/game/world"
	"github.com/gemrs/gem/gem/protocol/game_protocol"
	"github.com/gemrs/gem/gem/runite"
	"github.com/gemrs/gem/gem/util/expire"
)

//go:generate glua .

// GameService represents the internal state of the game
//glua:bind
type GameService struct {
	expire.NonExpirable

	runite *runite.Context
	key    *crypto.Keypair
	auth   auth.Provider
	world  *world.Instance
}

//glua:bind constructor GameService
func NewGameService(runite *runite.Context, rsaKeyPath string, auth auth.Provider) *GameService {
	var err error
	var key *crypto.Keypair
	key, err = crypto.LoadPrivateKey(rsaKeyPath)
	if err != nil {
		panic(err)
	}

	svc := &GameService{
		runite:       runite,
		key:          key,
		auth:         auth,
		NonExpirable: expire.NewNonExpirable(),
		world:        world.NewInstance(),
	}

	engine_event.Tick.Register(event.NewObserver(svc, svc.PlayerTick))
	return svc
}

func (svc *GameService) NewClient(conn *server.Connection, service int) server.Client {
	conn.Log().Info("new game client")
	client := player.NewPlayer(conn, svc.world)
	client.SetDecodeFunc(svc.handshake)
	return client
}

func (svc *GameService) World() *world.Instance {
	return svc.world
}

func doForAllPlayers(entities []entity.Entity, fn func(*player.Player)) {
	for _, e := range entities {
		p := e.(*player.Player)
		fn(p)
	}
}

func (svc *GameService) PlayerTick(ev *event.Event, _args ...interface{}) {
	allPlayers := svc.world.AllEntities(entity.PlayerType)

	// Ordering is important here. We want to run these things in this specific order,
	// and 'concurrently' (albeit in a single thread) for each player. ie. All waypoint
	// queues are updated for all players before syncing entity lists.
	doForAllPlayers(allPlayers, (*player.Player).SyncInventories)
	doForAllPlayers(allPlayers, (*player.Player).UpdateWaypointQueue)
	doForAllPlayers(allPlayers, (*player.Player).SyncEntityList)
	doForAllPlayers(allPlayers, (*player.Player).SendPlayerSync)
	doForAllPlayers(allPlayers, (*player.Player).SendGroundItemSync)
	doForAllPlayers(allPlayers, (*player.Player).UpdateVisibleEntities)
	doForAllPlayers(allPlayers, (*player.Player).ProcessChatQueue)
	doForAllPlayers(allPlayers, (*player.Player).ClearFlags)
	svc.world.UpdateEntityCollections()
}

// decodePacket decodes from the readBuffer using the ISAAC rand generator
func (svc *GameService) decodePacket(client *player.Player) error {
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
func (svc *GameService) packetConsumer(client *player.Player) {
L:
	for {
		select {
		case <-client.Conn().DisconnectChan:
			break L
		case pkt := <-client.Conn().Read:
			if _, ok := pkt.(*game_protocol.UnknownPacket); ok {
				/* unknown packet; dump to the log */
				client.Log().Debug("Got unknown packet: %v", pkt)
				continue
			}
			packet.Dispatch(client, pkt)
		}

	}
}
