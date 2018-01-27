package protocol_os_157

import (
	"bytes"
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/protocol"
)

// +gen define_outbound:"Pkt82,SzVar16"
type PlayerUpdate protocol.PlayerUpdate

func (struc PlayerUpdate) attachment() *playerData {
	return struc.Me.ProtoData.(*playerData)
}

func (struc PlayerUpdate) Encode(w_ io.Writer, flags interface{}) {
	w := &bytes.Buffer{}
	data := struc.attachment()

	maskBuf := encoding.NewBuffer()
	struc.processPlayersInViewport(w, maskBuf, true)
	struc.processPlayersInViewport(w, maskBuf, false)
	struc.processPlayersOutsideViewport(w, maskBuf, true)
	struc.processPlayersOutsideViewport(w, maskBuf, false)

	for i, _ := range data.skipFlags {
		data.skipFlags[i] >>= 1
	}

	maskBytes := maskBuf.Bytes()
	if len(maskBytes) > 0 {
		w.Write(maskBytes)
	}

	w_.Write(w.Bytes())
}

var skip int

func (struc PlayerUpdate) processPlayersInViewport(w io.Writer, maskBuf *encoding.Buffer, nsn bool) {
	buf := encoding.NewBitBuffer(w)
	defer buf.Close()
	data := struc.attachment()

	skip = 0
	for i, p := range struc.Visible {
		update := data.skipFlags[p.Index]&0x1 != 0
		if nsn {
			update = data.skipFlags[p.Index]&0x1 == 0
		}

		if update {
			if skip > 0 {
				skip--
				data.skipFlags[p.Index] |= 0x2
				continue
			}

			flags := struc.getModifiedUpdateFlags(p)

			if flags != 0 {
				buf.Write(1, 1)
				struc.updatePlayerInViewport(buf, maskBuf, p)
			} else {
				buf.Write(1, 0)
				struc.skipPlayersInViewport(buf, i, nsn)
				data.skipFlags[p.Index] |= 0x2
			}
		}
	}
}

func (struc PlayerUpdate) processPlayersOutsideViewport(w io.Writer, maskBuf *encoding.Buffer, nsn bool) {
	buf := encoding.NewBitBuffer(w)
	defer buf.Close()

	if nsn {
		buf.Write(1, 0)
		struc.writeSkip(buf, 2045)
	}
}

func (struc PlayerUpdate) skipPlayersInViewport(buf *encoding.BitBuffer, i int, nsn bool) {
	data := struc.attachment()

	for x := i + 1; x < len(struc.Visible); x++ {
		p := struc.Visible[x]
		update := data.skipFlags[p.Index]&0x1 != 0
		if nsn {
			update = data.skipFlags[p.Index]&0x1 == 0
		}

		if update {
			flags := struc.getModifiedUpdateFlags(p)
			if flags != 0 {
				break
			}
			skip++
		}
	}
	struc.writeSkip(buf, skip)
}

func (struc PlayerUpdate) writeSkip(buf *encoding.BitBuffer, skip int) {
	switch {
	case skip == 0:
		buf.Write(2, 0)
	case skip < 32:
		buf.Write(2, 1)
		buf.Write(5, uint32(skip))

	case skip < 256:
		buf.Write(2, 2)
		buf.Write(8, uint32(skip))

	case skip < 2048:
		buf.Write(2, 3)
		buf.Write(11, uint32(skip))

	default:
		panic("invalid skip size")
	}
}

func (struc PlayerUpdate) updatePlayerInViewport(buf *encoding.BitBuffer, maskBuf *encoding.Buffer, thisPlayer protocol.PlayerUpdateData) {
	flags := struc.getModifiedUpdateFlags(thisPlayer) & ^entity.MobFlagMovementUpdate
	if flags != 0 {
		buf.Write(1, 1)
	} else {
		buf.Write(1, 0)
	}

	struc.buildMovementBlock(buf, thisPlayer)

	if flags > 0 {
		if flags >= 0x100 {
			flags |= 0x4
			encoding.Uint8(flags&0xFF).Encode(maskBuf, encoding.IntNilFlag)
			encoding.Uint8(flags>>8).Encode(maskBuf, encoding.IntNilFlag)
		} else {
			encoding.Uint8(flags&0xFF).Encode(maskBuf, encoding.IntNilFlag)
		}
	}

	/* Update appearance */
	if (flags & entity.MobFlagIdentityUpdate) != 0 {
		appearanceBuf := encoding.NewBuffer()
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

			Name:        encoding.String(thisPlayer.Username),
			CombatLevel: encoding.Uint8(thisPlayer.Skills.CombatLevel()),
		}

		appearanceBlock.Encode(appearanceBuf, nil)

		block := appearanceBuf.Bytes()
		blockSize := encoding.Uint8(len(block))
		blockSize.Encode(maskBuf, encoding.IntOffset128)

		for i := len(block) - 1; i >= 0; i-- {
			err := maskBuf.WriteByte(block[i])
			if err != nil {
				panic(err)
			}
		}
	}

	return
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

	switch {
	case (flags & entity.MobFlagRunUpdate) != 0:
		current, last := player.WaypointQueue.WalkDirection()

		buf.Write(2, 2) // update type 2 = running
		buf.Write(3, uint32(last))
		buf.Write(3, uint32(current))

	case (flags & entity.MobFlagWalkUpdate) != 0:
		current, _ := player.WaypointQueue.WalkDirection()

		buf.Write(2, 1) // update type 1 = walking
		buf.Write(3, uint32(current))
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
