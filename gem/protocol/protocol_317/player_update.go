package protocol_317

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/protocol"
)

// +gen define_outbound:"Pkt81,SzVar16"
type PlayerUpdate protocol.PlayerUpdate

func (struc PlayerUpdate) Encode(w io.Writer, flags interface{}) {
	buf := encoding.NewBitBuffer(w)

	updateBlock := encoding.NewBuffer()

	// Update our player
	struc.buildUpdateBlock(updateBlock, struc.Me)
	struc.buildMovementBlock(buf, struc.Me)

	// Update known players
	buf.Write(8, uint32(len(struc.Updating)))
	for _, idx := range struc.Updating {
		if struc.Removing[idx] {
			buf.Write(1, 1)
			buf.Write(2, 3)
		} else {
			other := struc.Others[idx]
			struc.buildMovementBlock(buf, other)
			if other.Flags != 0 {
				struc.buildUpdateBlock(updateBlock, other)
			}
		}
	}

	// Add new players
	for _, idx := range struc.Adding {
		other, ok := struc.Others[idx]
		if !ok {
			panic("missing player being added")
		}
		struc.addPlayer(buf, other)

		// Force appearance update
		other.Flags |= entity.MobFlagIdentityUpdate
		struc.buildUpdateBlock(updateBlock, other)
	}

	updateBlockBytes := updateBlock.Bytes()
	if len(updateBlockBytes) > 0 {
		buf.Write(11, 0x7FF)
		buf.Close()
		w.Write(updateBlockBytes)
	} else {
		buf.Close()
	}
}

func (struc PlayerUpdate) addPlayer(buf *encoding.BitBuffer, other protocol.PlayerUpdateData) {
	buf.Write(11, uint32(other.Index))

	buf.Write(1, 1) // Update required
	buf.Write(1, 1) // Discard walk queue

	deltaX, deltaY, _ := other.Position.Delta(struc.Me.Position)
	buf.Write(5, uint32(deltaY))
	buf.Write(5, uint32(deltaX))
}

func (struc PlayerUpdate) getModifiedUpdateFlags(updatingPlayer protocol.PlayerUpdateData) entity.Flags {
	// Clear some flags that don't apply to self updates
	updatingThisPlayer := struc.Me.Index == updatingPlayer.Index
	flags := updatingPlayer.Flags
	if updatingThisPlayer {
		flags = flags & ^entity.MobFlagChatUpdate
	}
	return flags
}

func (struc PlayerUpdate) buildMovementBlock(buf *encoding.BitBuffer, player protocol.PlayerUpdateData) {
	flags := struc.getModifiedUpdateFlags(player)

	// Anything to do?
	if flags == 0 {
		buf.Write(1, 0) // No updates
		return
	}
	buf.Write(1, 1) // This player has updates

	// Do we have any non-movement updates to perform?
	otherUpdateFlags := (flags & ^entity.MobFlagMovementUpdate)

	updatingThisPlayer := struc.Me.Index == player.Index

	switch {
	// When updating other players, don't send warp movements
	case updatingThisPlayer && (flags&entity.MobFlagRegionUpdate) != 0:
		localPos := player.Position.LocalTo(struc.Me.LoadedRegion)

		buf.Write(2, 3) // update type 3 = warp to location
		buf.Write(2, uint32(localPos.Z()))
		buf.WriteBit(false) // discard walk queue? not sure when/if we need this
		buf.WriteBit(otherUpdateFlags != 0)
		buf.Write(7, uint32(localPos.Y()))
		buf.Write(7, uint32(localPos.X()))

	case (flags & entity.MobFlagRunUpdate) != 0:
		current, last := player.WaypointQueue.WalkDirection()

		buf.Write(2, 2) // update type 2 = running
		buf.Write(3, uint32(last))
		buf.Write(3, uint32(current))
		buf.WriteBit(otherUpdateFlags != 0)

	case (flags & entity.MobFlagWalkUpdate) != 0:
		current, _ := player.WaypointQueue.WalkDirection()

		buf.Write(2, 1) // update type 1 = walking
		buf.Write(3, uint32(current))
		buf.WriteBit(otherUpdateFlags != 0)

	default:
		buf.Write(2, 0) // update type 0 = no movement updates
	}
}

func (struc PlayerUpdate) buildChatUpdateBlock(w io.Writer, other protocol.PlayerUpdateData) {
	message := other.ChatMessageQueue[0]
	encoding.Uint8(message.Effects).Encode(w, encoding.IntegerFlag(encoding.IntNilFlag))
	encoding.Uint8(message.Colour).Encode(w, encoding.IntegerFlag(encoding.IntNilFlag))
	encoding.Uint8(other.Rights).Encode(w, encoding.IntegerFlag(encoding.IntNilFlag))
	encoding.Uint8(len(message.PackedMessage)).Encode(w, encoding.IntegerFlag(encoding.IntNegate))
	encoding.Bytes(message.PackedMessage).Encode(w, 0)
}

func (struc PlayerUpdate) buildUpdateBlock(w io.Writer, thisPlayer protocol.PlayerUpdateData) {
	flags := struc.getModifiedUpdateFlags(thisPlayer) & ^entity.MobFlagMovementUpdate

	if flags == 0 {
		return
	}

	if flags >= 256 {
		flags |= 64
		flagsEnc := encoding.Uint16(flags)
		flagsEnc.Encode(w, encoding.IntLittleEndian)
	} else {
		flagsEnc := encoding.Uint8(flags)
		flagsEnc.Encode(w, encoding.IntNilFlag)
	}

	/* Update chat */
	if (flags & entity.MobFlagChatUpdate) != 0 {
		struc.buildChatUpdateBlock(w, thisPlayer)
	}

	/* Update appearance */
	if (flags & entity.MobFlagIdentityUpdate) != 0 {
		buf := encoding.NewBuffer()
		appearance := thisPlayer.Appearance
		anims := thisPlayer.Animations
		appearanceBlock := OutboundPlayerAppearance{
			Gender:   encoding.Uint8(appearance.Gender()),
			HeadIcon: encoding.Uint8(appearance.HeadIcon()),

			HelmModel:       encoding.Uint8(0),
			CapeModel:       encoding.Uint8(0),
			AmuletModel:     encoding.Uint8(0),
			RightWieldModel: encoding.Uint8(0),
			TorsoModel:      encoding.Uint16(256 + appearance.Model(protocol.BodyPartTorso)),
			LeftWieldModel:  encoding.Uint8(0),
			ArmsModel:       encoding.Uint16(256 + appearance.Model(protocol.BodyPartArms)),
			LegsModel:       encoding.Uint16(256 + appearance.Model(protocol.BodyPartLegs)),
			HeadModel:       encoding.Uint16(256 + appearance.Model(protocol.BodyPartHead)),
			HandsModel:      encoding.Uint16(256 + appearance.Model(protocol.BodyPartHands)),
			FeetModel:       encoding.Uint16(256 + appearance.Model(protocol.BodyPartFeet)),
			BeardModel:      encoding.Uint16(256 + appearance.Model(protocol.BodyPartBeard)),

			HairColor:  encoding.Uint8(appearance.Color(protocol.BodyPartHair)),
			TorsoColor: encoding.Uint8(appearance.Color(protocol.BodyPartTorso)),
			LegColor:   encoding.Uint8(appearance.Color(protocol.BodyPartLegs)),
			FeetColor:  encoding.Uint8(appearance.Color(protocol.BodyPartFeet)),
			SkinColor:  encoding.Uint8(appearance.Color(protocol.BodyPartSkin)),

			AnimIdle:       encoding.Uint16(anims.Animation(protocol.AnimIdle)),
			AnimSpotRotate: encoding.Uint16(anims.Animation(protocol.AnimSpotRotate)),
			AnimWalk:       encoding.Uint16(anims.Animation(protocol.AnimWalk)),
			AnimRotate180:  encoding.Uint16(anims.Animation(protocol.AnimRotate180)),
			AnimRotateCCW:  encoding.Uint16(anims.Animation(protocol.AnimRotateCCW)),
			AnimRotateCW:   encoding.Uint16(anims.Animation(protocol.AnimRotateCW)),
			AnimRun:        encoding.Uint16(anims.Animation(protocol.AnimRun)),

			Name:        encoding.NameHash(thisPlayer.Username),
			CombatLevel: encoding.Uint8(thisPlayer.Skills.CombatLevel()),
		}

		appearanceBlock.Encode(buf, nil)

		block := buf.Bytes()
		blockSize := encoding.Uint8(len(block))
		blockSize.Encode(w, encoding.IntOffset128)

		_, err := w.Write(block)
		if err != nil {
			panic(err)
		}
	}
	return
}
