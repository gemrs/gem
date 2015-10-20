package game

import (
	"gem/encoding"
	"gem/event"
	"gem/game/entity"
	"gem/game/player"
	"gem/protocol"
)

func (client *Player) PlayerInit() {
	client.Conn().Write <- &protocol.OutboundPlayerInit{
		Membership: encoding.Int8(1),
		Index:      encoding.Int16(client.Index()),
	}
}

func (client *Player) RegionUpdate(_ *event.Event, _ ...interface{}) {
	sector := client.Position().Sector()
	client.Conn().Write <- &protocol.OutboundRegionUpdate{
		SectorX: encoding.Int16(sector.X),
		SectorY: encoding.Int16(sector.Y),
	}

	session := client.Session().(*Session)
	session.SetFlags(entity.MobFlagRegionUpdate)
}

func (client *Player) AppearanceUpdate(_ *event.Event, _ ...interface{}) {
	session := client.Session().(*Session)
	session.SetFlags(entity.MobFlagIdentityUpdate)
}

func (client *Player) PlayerUpdate(_ *event.Event, _ ...interface{}) {
	client.Conn().Write <- &protocol.PlayerUpdate{
		OurPlayer: player.Snapshot(client),
	}
}

func (client *Player) ClearUpdateFlags(_ *event.Event, _ ...interface{}) {
	session := client.Session().(*Session)
	session.ClearFlags()
}
