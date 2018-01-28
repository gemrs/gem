//glua:bind module gem.game
package game

import (
	"fmt"

	"github.com/gemrs/gem/gem/auth"
	"github.com/gemrs/gem/gem/core/crypto"
	"github.com/gemrs/gem/gem/core/event"
	engine_event "github.com/gemrs/gem/gem/engine/event"
	"github.com/gemrs/gem/gem/game/entity"
	game_event "github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/impl"
	"github.com/gemrs/gem/gem/game/item"
	"github.com/gemrs/gem/gem/game/packet"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/game/world"
	"github.com/gemrs/gem/gem/protocol"
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

	item.LoadDefinitions(runite)

	engine_event.Tick.Register(event.NewObserver(svc, svc.PlayerTick))
	return svc
}

func (svc *GameService) NewClient(conn *server.Connection, service int) server.GameClient {
	conn.Log().Info("new game client")
	slot := svc.world.FindPlayerSlot()
	client := impl.NewPlayer(slot, conn, svc.world)
	svc.world.SetPlayerSlot(slot, client)

	loginHandler := server.Proto.NewLoginHandler()
	loginHandler.SetServerIsaacSeed(client.ServerIsaacSeed())
	loginHandler.SetRsaKeypair(svc.key)
	loginHandler.SetCompleteCallback(func(loginHandler server.LoginHandler) error {
		username, password := loginHandler.Username(), loginHandler.Password()
		profile, responseCode := svc.auth.LookupProfile(username, password)

		if responseCode != protocol.AuthOkay {
			client.Send(protocol.OutboundLoginResponse{
				Response: responseCode,
			})
			return nil
		}

		client.InitIsaac(loginHandler.IsaacSeeds())
		client.SetProfile(profile)

		// Successful login, do all the stuff
		client.Send(protocol.OutboundLoginResponse{
			Response: responseCode,
			Rights:   int(client.Profile().Rights()),
			Flagged:  false,
			Index:    client.Index(),
		})

		client.SetDecodeFunc(svc.decodePacket)
		go svc.packetConsumer(client)

		client.LoadProfile()
		client.FinishInit()
		game_event.PlayerLogin.NotifyObservers(client)

		go func() {
			client.Conn().WaitForDisconnect()
			worldSector := svc.world.Sector(client.Position().Sector())
			worldSector.Remove(client)
			svc.world.SetPlayerSlot(slot, nil)
			game_event.PlayerLogout.NotifyObservers(client)
		}()
		return nil
	})
	loginHandler.Perform(client)
	return client
}

func (svc *GameService) World() *world.Instance {
	return svc.world
}

func doForAllPlayers(entities []entity.Entity, fn func(protocol.Player)) {
	for _, e := range entities {
		p := e.(protocol.Player)
		fn(p)
	}
}

func (svc *GameService) PlayerTick(ev *event.Event, _args ...interface{}) {
	allPlayers := svc.world.AllEntities(entity.PlayerType)

	// Ordering is important here. We want to run these things in this specific order,
	// and 'concurrently' (albeit in a single thread) for each player. ie. All waypoint
	// queues are updated for all players before syncing entity lists.
	doForAllPlayers(allPlayers, (protocol.Player).UpdateInteractionQueue)
	doForAllPlayers(allPlayers, (protocol.Player).SyncInventories)
	doForAllPlayers(allPlayers, (protocol.Player).SyncEntityList)
	doForAllPlayers(allPlayers, (protocol.Player).SendPlayerSync)
	doForAllPlayers(allPlayers, (protocol.Player).SendGroundItemSync)
	doForAllPlayers(allPlayers, (protocol.Player).UpdateVisibleEntities)
	doForAllPlayers(allPlayers, (protocol.Player).ProcessChatQueue)
	doForAllPlayers(allPlayers, (protocol.Player).ClearFlags)
	svc.world.UpdateEntityCollections()
}

// decodePacket decodes from the readBuffer using the ISAAC rand generator
func (svc *GameService) decodePacket(client server.GameClient) error {
	b := client.Conn().ReadBuffer
	data, err := b.Peek(1)
	if err != nil {
		panic(err)
	}

	idByte := int(data[0])

	rand := client.IsaacIn().Rand()
	realId := uint8(uint32(idByte) - rand)
	realId = uint8(idByte)
	packet, err := server.Proto.NewInboundPacket(int(realId))
	if err != nil {
		return fmt.Errorf("%v: packet %v", err, realId)
	}
	packet.Decode(b, rand)

	if !client.Conn().IsDisconnecting() {
		client.Conn().Read <- packet
	}
	return nil
}

// packetConsumer is the goroutine which picks packets from the readQueue and does something with them
func (svc *GameService) packetConsumer(client protocol.Player) {
L:
	for {
		select {
		case <-client.Conn().DisconnectChan:
			break L
		case pkt := <-client.Conn().Read:
			message := server.Proto.Decode(pkt)
			switch message := message.(type) {
			case *protocol.UnknownPacket:
				/* unknown packet; dump to the log */
				client.Log().Debug("Got unknown packet: %v", pkt)
				continue
			default:
				packet.Dispatch(client, message)
			}

		}

	}
}
