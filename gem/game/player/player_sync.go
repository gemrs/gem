package player

import (
	"github.com/sinusoids/gem/gem/encoding"
	engine_event "github.com/sinusoids/gem/gem/engine/event"
	"github.com/sinusoids/gem/gem/event"
	"github.com/sinusoids/gem/gem/game/interface/entity"
	"github.com/sinusoids/gem/gem/game/interface/player"
	game_protocol "github.com/sinusoids/gem/gem/protocol/game"
)

// FinishInit is called once the player has finished the low level login sequence
func (client *Player) FinishInit() {
	client.Conn().Write <- &game_protocol.OutboundPlayerInit{
		Membership: encoding.Int8(1),
		Index:      encoding.Int16(client.Index()),
	}
	engine_event.Tick.Register(event.NewListener(client.PlayerUpdate))
	engine_event.PostTick.Register(event.NewListener(client.ClearUpdateFlags))
}

// RegionUpdate is called when the player enters a new region
// It sends the region update packet and sets the correct update flags
func (client *Player) RegionUpdate(_ *event.Event, _ ...interface{}) {
	sector := client.Position().Sector()
	client.Conn().Write <- &game_protocol.OutboundRegionUpdate{
		SectorX: encoding.Int16(sector.X()),
		SectorY: encoding.Int16(sector.Y()),
	}

	session := client.Session().(*Session)
	session.SetFlags(entity.MobFlagRegionUpdate)
}

// AppearanceUpdate is called when the player's appearance changes
// It ensures the player's appearance is synced at next update
func (client *Player) AppearanceUpdate(_ *event.Event, _ ...interface{}) {
	session := client.Session().(*Session)
	session.SetFlags(entity.MobFlagIdentityUpdate)
}

// PlayerUpdate snapshots the player in their current state and syncs the client
func (client *Player) PlayerUpdate(_ *event.Event, _ ...interface{}) {
	client.Conn().Write <- &game_protocol.PlayerUpdate{
		OurPlayer: player.Snapshot(client),
	}
}

// ClearUpdate is called on post-tick, and clears the player's update flags
func (client *Player) ClearUpdateFlags(_ *event.Event, _ ...interface{}) {
	session := client.Session().(*Session)
	session.ClearFlags()
}
