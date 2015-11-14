package player

import (
	"github.com/gemrs/gem/gem/encoding"
	engine_event "github.com/gemrs/gem/gem/engine/event"
	"github.com/gemrs/gem/gem/event"
	game_event "github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/interface/entity"
	"github.com/gemrs/gem/gem/game/interface/player"
	game_protocol "github.com/gemrs/gem/gem/protocol/game"
)

func (client *Player) LoadProfile() {
	profile := client.Profile().(*Profile)
	client.SetPosition(profile.Position())
	client.SetAppearance(profile.Appearance())

	game_event.PlayerLoadProfile.NotifyObservers(client, client.Profile().(*Profile))
}

// FinishInit is called once the player has finished the low level login sequence
func (client *Player) FinishInit() {
	client.Conn().Write <- &game_protocol.OutboundPlayerInit{
		Membership: encoding.Uint8(1),
		Index:      encoding.Uint16(client.Index()),
	}

	engine_event.PreTick.Register(event.NewListener(client, client.PreTick))
	engine_event.Tick.Register(event.NewListener(client, client.Tick))
	engine_event.PostTick.Register(event.NewListener(client, client.PostTick))
}

// Cleanup is called when the player logs out
func (client *Player) Cleanup() {

}

func (client *Player) SectorChange() {
	client.sector.Remove(client)
	client.sector = client.world.Sector(client.Position().Sector())
	client.sector.Add(client)
}

// RegionUpdate is called when the player enters a new region
// It sends the region update packet and sets the correct update flags
func (client *Player) RegionChange() {
	client.loadedRegion = client.Position().RegionOf()

	sector := client.Position().Sector()
	client.Conn().Write <- &game_protocol.OutboundRegionUpdate{
		SectorX: encoding.Uint16(sector.X()),
		SectorY: encoding.Uint16(sector.Y()),
	}

	client.SetFlags(entity.MobFlagRegionUpdate)
}

func (client *Player) PreTick(_ *event.Event, _ ...interface{}) {
	client.WaypointQueue().Tick(client)
	client.Profile().SetPosition(client.Position())
}

// Tick snapshots the player in their current state and syncs the client
// Also re-syncs the current session with the player's profile
func (client *Player) Tick(_ *event.Event, _ ...interface{}) {
	client.Conn().Write <- &game_protocol.PlayerUpdate{
		OurPlayer: player.Snapshot(client),
	}
}

// ClearUpdate is called on post-tick, and clears the player's update flags
func (client *Player) PostTick(_ *event.Event, _ ...interface{}) {
	client.ClearFlags()
}
