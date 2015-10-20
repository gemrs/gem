package game

import (
	"gem/encoding"
	"gem/event"
	"gem/game/entity"
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

	client.flags |= entity.MobFlagRegionUpdate
	client.Log().Debugf("warp flags %v", client.Flags())
}

func (client *Player) AppearanceUpdate(_ *event.Event, _ ...interface{}) {
	client.flags |= entity.MobFlagIdentityUpdate
}

func (client *Player) PlayerUpdate(_ *event.Event, _ ...interface{}) {
	client.Log().Debugf("doing update %v", client.Flags())
	client.Conn().Write <- &protocol.PlayerUpdate{
		OurPlayer: client,
	}
}

func (client *Player) ClearUpdateFlags(_ *event.Event, _ ...interface{}) {
	client.Log().Debugf("clearing flags %v", client.Flags())

	client.flags = 0
}
