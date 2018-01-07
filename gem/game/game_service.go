//glua:bind module gem.game
package game

import (
	"fmt"

	"github.com/gemrs/gem/gem/auth"
	"github.com/gemrs/gem/gem/crypto"
	engine_event "github.com/gemrs/gem/gem/engine/event"
	"github.com/gemrs/gem/gem/event"
	game_event "github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/interface/entity"
	"github.com/gemrs/gem/gem/game/interface/player"
	"github.com/gemrs/gem/gem/game/packet"
	playerimpl "github.com/gemrs/gem/gem/game/player"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/game/world"
	game_protocol "github.com/gemrs/gem/gem/protocol/game"
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

	game_event.PlayerFinishLogin.Register(event.NewObserver(svc, playerFinishLogin))
	game_event.PlayerLogout.Register(event.NewObserver(svc, playerCleanup))
	game_event.EntityRegionChange.Register(event.NewObserver(svc, svc.EntityUpdate))
	game_event.EntitySectorChange.Register(event.NewObserver(svc, svc.EntityUpdate))
	game_event.PlayerAppearanceUpdate.Register(event.NewObserver(svc, svc.PlayerUpdate))

	engine_event.Tick.Register(event.NewObserver(svc, svc.PlayerTick))
	return svc
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
	conn.Log().Info("new game client")
	client := playerimpl.NewPlayer(conn, svc.world)
	client.SetDecodeFunc(svc.handshake)
	return client
}

func (svc *GameService) World() *world.Instance {
	return svc.world
}

func doForAllPlayers(entities []entity.Entity, fn func(*playerimpl.Player)) {
	for _, e := range entities {
		p := e.(*playerimpl.Player)
		fn(p)
	}
}

func (svc *GameService) PlayerTick(ev *event.Event, _args ...interface{}) {
	allPlayers := svc.world.AllEntities(entity.PlayerType)

	doForAllPlayers(allPlayers, (*playerimpl.Player).UpdateWaypointQueue)
	doForAllPlayers(allPlayers, (*playerimpl.Player).SyncEntityList)
	doForAllPlayers(allPlayers, (*playerimpl.Player).SendPlayerSync)
	doForAllPlayers(allPlayers, (*playerimpl.Player).ClearFlags)
	doForAllPlayers(allPlayers, (*playerimpl.Player).UpdateVisibleEntities)
	svc.world.UpdateEntityCollections()
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
				client.Log().Debug("Got unknown packet: %v", pkt)
				continue
			}
			packet.Dispatch(client, pkt)
		}

	}
}
