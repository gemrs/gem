package game

import (
	"gem/encoding"
	"gem/event"
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

	client.flags |= protocol.MobFlagRegionUpdate
}

func (client *Player) PlayerUpdate(_ *event.Event, _ ...interface{}) {
	updateBlock := protocol.PlayerUpdate{
		UpdateFlags: client.flags,
	}

	if (client.flags & protocol.MobFlagMovementUpdate) != 0 {
		switch {
		case (client.flags & protocol.MobFlagRegionUpdate) != 0:
			updateBlock.OurMovementBlock.Warp = protocol.WarpMovement{
				Location:         client.Position().LocalTo(client.region),
				DiscardWalkQueue: true,
			}

		case (client.flags & protocol.MobFlagRunUpdate) != 0:
			updateBlock.OurMovementBlock.Run = protocol.RunMovement{}

		case (client.flags & protocol.MobFlagWalkUpdate) != 0:
			updateBlock.OurMovementBlock.Walk = protocol.WalkMovement{}

		}
	}

	client.Conn().Write <- &updateBlock
}

func (client *Player) ClearUpdateFlags(_ *event.Event, _ ...interface{}) {
	client.flags = 0
}
