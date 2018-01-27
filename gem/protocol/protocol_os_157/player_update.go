package protocol_os_157

import (
	"bytes"
	"fmt"
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

	//struc.Me.Log.Debug("local player list is %v\n", data.localPlayers)

	maskBuf := encoding.NewBuffer()
	struc.processLocalPlayers(w, maskBuf, 0)
	struc.processLocalPlayers(w, maskBuf, 1)
	struc.processExternalPlayers(w, maskBuf, 2)
	struc.processExternalPlayers(w, maskBuf, 3)

	maskBytes := maskBuf.Bytes()
	if len(maskBytes) > 0 {
		w.Write(maskBytes)
	}

	w_.Write(w.Bytes())

	// Rebuild the player lists to preserve index ordering
	data.localPlayerCount = 0
	data.externalPlayerCount = 0
	for i, _ := range data.skipFlags {
		if i == 0 {
			continue
		}

		data.skipFlags[i].cycle()
		if _, ok := struc.Others[i]; ok {
			data.localPlayers[data.localPlayerCount] = i
			data.localPlayerCount++
		} else {
			data.externalPlayers[data.externalPlayerCount] = i
			data.externalPlayerCount++
		}
	}
}

func (struc PlayerUpdate) processLocalPlayers(w io.Writer, maskBuf *encoding.Buffer, iter int) {
	buf := encoding.NewBitBuffer(w)
	defer buf.Close()
	data := struc.attachment()

	data.skipCount = 0
	for i := 0; i < data.localPlayerCount; i++ {
		index := data.localPlayers[i]

		if data.skipFlags[index].shouldUpdate(iter) {
			if data.skipCount > 0 {
				data.skipCount--
				data.skipFlags[index].updateNextIter()
				continue
			}

			player, ok := struc.Others[index]

			updateRequired := false
			if ok {
				if struc.getModifiedUpdateFlags(player) != 0 {
					updateRequired = true
				}
			}

			if updateRequired {
				buf.Write(1, 1)
				struc.updateLocalPlayers(buf, maskBuf, player)
			} else {
				buf.Write(1, 0)
				struc.skipLocalPlayers(buf, i, iter)
				data.skipFlags[index].updateNextIter()
			}
		}
	}
}

func (struc PlayerUpdate) isAdding(playerIndex int) bool {
	for _, i := range struc.Adding {
		if i == playerIndex {
			return true
		}
	}
	return false
}

func (struc PlayerUpdate) processExternalPlayers(w io.Writer, maskBuf *encoding.Buffer, iter int) {
	buf := encoding.NewBitBuffer(w)
	defer buf.Close()
	data := struc.attachment()

	data.skipCount = 0
	for i := 0; i < data.externalPlayerCount; i++ {
		index := data.externalPlayers[i]

		if data.skipFlags[index].shouldUpdate(iter) {
			if data.skipCount > 0 {
				data.skipCount--
				data.skipFlags[index].updateNextIter()
				continue
			}

			if index != struc.Me.Index && struc.isAdding(index) {
				buf.Write(1, 1)
				player, ok := struc.Others[index]
				if !ok {
					panic(fmt.Errorf("don't have player data for index %v\n", index))
				}

				struc.addPlayer(buf, maskBuf, player)
				data.skipFlags[index].updateNextIter()
			} else {
				buf.Write(1, 0)
				struc.skipExternalPlayers(buf, i, iter)
				data.skipFlags[index].updateNextIter()
			}
		}
	}
}

func (struc PlayerUpdate) skipLocalPlayers(buf *encoding.BitBuffer, i int, iter int) {
	data := struc.attachment()

	for x := i + 1; x < data.localPlayerCount; x++ {
		p, ok := struc.Others[x]
		if !ok {
			panic(fmt.Errorf("don't have player data for index %v\n", x))
		}

		index := data.localPlayers[x]

		if data.skipFlags[index].shouldUpdate(iter) {
			flags := struc.getModifiedUpdateFlags(p)
			if flags != 0 {
				break
			}
			data.skipCount++
		}
	}
	struc.writeSkip(buf, data.skipCount)
}

func (struc PlayerUpdate) skipExternalPlayers(buf *encoding.BitBuffer, i int, iter int) {
	data := struc.attachment()

	for x := i + 1; x < data.externalPlayerCount; x++ {
		index := data.externalPlayers[x]
		if data.skipFlags[index].shouldUpdate(iter) {
			if index != struc.Me.Index && struc.isAdding(index) {
				break
			}
			data.skipCount++
		}
	}
	struc.writeSkip(buf, data.skipCount)
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

func (struc PlayerUpdate) addPlayer(buf *encoding.BitBuffer, maskBuf *encoding.Buffer, player protocol.PlayerUpdateData) {
	// Add player
	buf.Write(2, 0)

	// No region updates
	buf.Write(1, 0)

	// Absolute position
	buf.Write(13, uint32(player.Position.X()))
	buf.Write(13, uint32(player.Position.Y()))

	// Also send flag based updates
	buf.Write(1, 1)

	// Force identity update
	flags := struc.getModifiedUpdateFlags(player) & ^entity.MobFlagMovementUpdate
	flags |= entity.MobFlagIdentityUpdate
	struc.buildBlockUpdates(maskBuf, player, flags)
}

func (struc PlayerUpdate) updateLocalPlayers(buf *encoding.BitBuffer, maskBuf *encoding.Buffer, thisPlayer protocol.PlayerUpdateData) {
	flags := struc.getModifiedUpdateFlags(thisPlayer) & ^entity.MobFlagMovementUpdate
	if flags != 0 {
		buf.Write(1, 1)
	} else {
		buf.Write(1, 0)
	}

	struc.buildMovementBlock(buf, thisPlayer)
	struc.buildBlockUpdates(maskBuf, thisPlayer, flags)
}

func (struc PlayerUpdate) buildBlockUpdates(maskBuf *encoding.Buffer, thisPlayer protocol.PlayerUpdateData, flags entity.Flags) {
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
		struc.Me.Log.Debug("movement update for player %v", player.Index)
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
