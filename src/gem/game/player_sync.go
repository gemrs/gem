package game

import (
	"gem/encoding"
	"gem/event"
	"gem/protocol"
)

func (client *GameClient) PlayerInit() {
	client.Conn().Write <- &protocol.OutboundPlayerInit{
		Membership: encoding.Int8(1),
		Index:      encoding.Int16(client.Index()),
	}
}

func (client *GameClient) RegionUpdate(_ *event.Event, _ ...interface{}) {
	sector := client.Position().Sector()
	client.Conn().Write <- &protocol.OutboundRegionUpdate{
		SectorX: encoding.Int16(sector.X),
		SectorY: encoding.Int16(sector.Y),
	}
}
