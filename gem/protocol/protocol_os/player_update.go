package protocol_os

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/game/data"
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/item"
	"github.com/gemrs/gem/gem/protocol"
)

type PlayerUpdate protocol.PlayerUpdate

func (struc PlayerUpdate) attachment() *PlayerData {
	return struc.Me.ProtoData().(*PlayerData)
}

func (struc PlayerUpdate) Encode(w_ io.Writer, flags interface{}) {
	w := &bytes.Buffer{}
	data := struc.attachment()

	translatePlayerFlags := flags.(func(entity.Flags) UpdateFlags)

	maskBuf := encoding.NewBuffer()
	struc.processLocalPlayers(w, maskBuf, 0, translatePlayerFlags)
	struc.processLocalPlayers(w, maskBuf, 1, translatePlayerFlags)
	struc.processExternalPlayers(w, maskBuf, 2, translatePlayerFlags)
	struc.processExternalPlayers(w, maskBuf, 3, translatePlayerFlags)

	maskBytes := maskBuf.Bytes()
	if len(maskBytes) > 0 {
		w.Write(maskBytes)
	}

	w_.Write(w.Bytes())

	// Rebuild the player lists to preserve index ordering
	allPlayers := struc.Me.WorldInstance().GetPlayers()
	data.LocalPlayerCount = 0
	data.ExternalPlayerCount = 0
	for i, _ := range data.skipFlags {
		if i == 0 {
			continue
		}

		data.skipFlags[i].cycle()
		p := allPlayers[i]

		if playerVisible(struc.Me, p) {
			data.LocalPlayers[data.LocalPlayerCount] = i
			data.LocalPlayerCount++
		} else {
			data.ExternalPlayers[data.ExternalPlayerCount] = i
			data.ExternalPlayerCount++
		}
	}
}

func (struc PlayerUpdate) getPlayer(index int) protocol.Player {
	allPlayers := struc.Me.WorldInstance().GetPlayers()
	return allPlayers[index]
}

func playerVisible(me, them protocol.Player) bool {
	if them == nil {
		return false
	}

	mySector := me.Position().Sector()
	theirSector := them.Position().Sector()
	dx, dy, dz := mySector.Delta(theirSector)
	return dx <= protocol.PlayerViewDistance && dy <= protocol.PlayerViewDistance && dz == 0
}

func (struc PlayerUpdate) processLocalPlayers(w io.Writer, maskBuf *encoding.Buffer, iter int, translatePlayerFlags func(entity.Flags) UpdateFlags) {
	buf := encoding.NewBitBuffer(w)
	defer buf.Close()
	data := struc.attachment()

	data.skipCount = 0
	for i := 0; i < data.LocalPlayerCount; i++ {
		index := data.LocalPlayers[i]

		if data.skipFlags[index].shouldUpdate(iter) {
			if data.skipCount > 0 {
				data.skipCount--
				data.skipFlags[index].updateNextIter()
				continue
			}

			player := struc.getPlayer(index)

			updateRequired := false
			if player != nil && struc.getModifiedUpdateFlags(player) != 0 {
				updateRequired = true
			}

			if !playerVisible(struc.Me, player) {
				buf.Write(1, 1)
				struc.removePlayer(buf, player)
			} else if updateRequired {
				buf.Write(1, 1)
				struc.updateLocalPlayers(buf, maskBuf, player, translatePlayerFlags)
			} else {
				buf.Write(1, 0)
				struc.skipLocalPlayers(buf, i, iter)
				data.skipFlags[index].updateNextIter()
			}
		}
	}
}

func (struc PlayerUpdate) processExternalPlayers(w io.Writer, maskBuf *encoding.Buffer, iter int, translatePlayerFlags func(entity.Flags) UpdateFlags) {
	buf := encoding.NewBitBuffer(w)
	defer buf.Close()
	data := struc.attachment()

	data.skipCount = 0
	for i := 0; i < data.ExternalPlayerCount; i++ {
		index := data.ExternalPlayers[i]

		if data.skipFlags[index].shouldUpdate(iter) {
			if data.skipCount > 0 {
				data.skipCount--
				data.skipFlags[index].updateNextIter()
				continue
			}

			if index != struc.Me.Index() && playerVisible(struc.Me, struc.getPlayer(index)) {
				buf.Write(1, 1)
				player := struc.getPlayer(index)
				if player == nil {
					panic(fmt.Errorf("don't have player data for index %v\n", index))
				}

				struc.addPlayer(buf, maskBuf, player, translatePlayerFlags)
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

	for x := i + 1; x < data.LocalPlayerCount; x++ {
		index := data.LocalPlayers[x]
		p := struc.getPlayer(index)
		if p == nil {
			panic(fmt.Errorf("don't have player data for index %v\n", index))
		}

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

	for x := i + 1; x < data.ExternalPlayerCount; x++ {
		index := data.ExternalPlayers[x]
		if data.skipFlags[index].shouldUpdate(iter) {
			if index != struc.Me.Index() && playerVisible(struc.Me, struc.getPlayer(index)) {
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

func (struc PlayerUpdate) addPlayer(buf *encoding.BitBuffer, maskBuf *encoding.Buffer, player protocol.Player, translatePlayerFlags func(entity.Flags) UpdateFlags) {
	// Add player
	buf.Write(2, 0)

	// No region updates
	buf.Write(1, 0)

	// Absolute position
	buf.Write(13, uint32(player.Position().X()))
	buf.Write(13, uint32(player.Position().Y()))

	// Also send flag based updates
	buf.Write(1, 1)

	// Force identity update
	flags := struc.getModifiedUpdateFlags(player) & ^entity.MobFlagMovementUpdate
	flags |= entity.MobFlagIdentityUpdate
	struc.buildBlockUpdates(maskBuf, player, flags, translatePlayerFlags)
}

func (struc PlayerUpdate) removePlayer(buf *encoding.BitBuffer, player protocol.Player) {
	buf.Write(1, 0)

	// No updates once put onto the external player list
	buf.Write(1, 0)
}

func (struc PlayerUpdate) updateLocalPlayers(buf *encoding.BitBuffer, maskBuf *encoding.Buffer, thisPlayer protocol.Player, translatePlayerFlags func(entity.Flags) UpdateFlags) {
	flags := struc.getModifiedUpdateFlags(thisPlayer) & ^entity.MobFlagMovementUpdate
	if flags != 0 {
		buf.Write(1, 1)
	} else {
		buf.Write(1, 0)
	}

	struc.buildMovementBlock(buf, thisPlayer)
	struc.buildBlockUpdates(maskBuf, thisPlayer, flags, translatePlayerFlags)
}

func (struc PlayerUpdate) buildBlockUpdates(maskBuf *encoding.Buffer, thisPlayer protocol.Player, entityFlags entity.Flags, translatePlayerFlags func(entity.Flags) UpdateFlags) {
	flags := translatePlayerFlags(entityFlags)

	if flags > 0 {
		if flags >= 0x100 {
			flags |= 0x2
			maskBuf.PutU8(int(flags & 0xFF))
			maskBuf.PutU8(int(flags >> 8))
		} else {
			maskBuf.PutU8(int(flags & 0xFF))
		}
	}

	if (entityFlags & entity.MobFlagChatUpdate) != 0 {
		struc.buildChatUpdateBlock(maskBuf, thisPlayer)
	}

	if (entityFlags & entity.MobFlagIdentityUpdate) != 0 {
		struc.buildAppearanceUpdateBlock(maskBuf, thisPlayer)
	}

	return
}

func (struc PlayerUpdate) buildChatUpdateBlock(maskBuf *encoding.Buffer, thisPlayer protocol.Player) {
	message := thisPlayer.ChatMessageQueue()[0]
	maskBuf.PutU16(message.Effects<<8 | message.Colour)
	maskBuf.Put8(int(thisPlayer.Profile().Rights()))

	// If 1, doesn't display overhead
	maskBuf.Put8(0, encoding.IntOffset128)

	huffmanBlock := encoding.NewBuffer()
	huffmanBlock.PutU8(len(message.Message), encoding.IntPacked)
	huffmanBlock.PutBytes(message.PackedMessage)

	maskBuf.PutU8(huffmanBlock.Len())
	huffmanBlock.WriteTo(maskBuf)
}

func (struc PlayerUpdate) buildAppearanceUpdateBlock(maskBuf *encoding.Buffer, thisPlayer protocol.Player) {
	appearance := thisPlayer.Profile().Appearance()
	appearanceBuf := encoding.NewBuffer()

	appearanceBuf.PutU8(appearance.Gender())
	appearanceBuf.PutU8(appearance.SkullIcon())
	appearanceBuf.PutU8(appearance.HeadIcon())

	var anims *data.Animations
	equipment := thisPlayer.Profile().Equipment()
	if weapon := equipment.Slot(item.EquipmentWeapon); weapon != nil {
		anims = weapon.Definition().WeaponData().CharAnimations()
	} else {
		anims = data.DefaultAnimations
	}

	var torsoCoversArms, helmCoversFace, helmCoversHair bool
	if torso := equipment.Slot(item.EquipmentTorso); torso != nil {
		torsoCoversArms = torso.Definition().EquipmentData().CoversArms()
	}

	if helm := equipment.Slot(item.EquipmentHelm); helm != nil {
		helmCoversFace = helm.Definition().EquipmentData().CoversFace()
		helmCoversHair = helm.Definition().EquipmentData().CoversHair()
	}

	for _, slot := range []int{
		item.EquipmentHelm,
		item.EquipmentCape,
		item.EquipmentAmulet,
		item.EquipmentWeapon} {
		if e := equipment.Slot(slot); e != nil {
			appearanceBuf.PutU16(0x200 + e.Definition().Id())
		} else {
			appearanceBuf.PutU8(0)
		}
	}

	if e := equipment.Slot(item.EquipmentTorso); e != nil {
		appearanceBuf.PutU16(0x200 + e.Definition().Id())
	} else {
		appearanceBuf.PutU16(0x100 + appearance.Model(protocol.BodyPartTorso))
	}

	if e := equipment.Slot(item.EquipmentShield); e != nil {
		appearanceBuf.PutU16(0x200 + e.Definition().Id())
	} else {
		appearanceBuf.Put8(0)
	}

	if e := equipment.Slot(item.EquipmentTorso); e != nil && !torsoCoversArms {
		appearanceBuf.PutU8(0)
	} else {
		appearanceBuf.PutU16(0x100 + appearance.Model(protocol.BodyPartArms))
	}

	if e := equipment.Slot(item.EquipmentLegs); e != nil {
		appearanceBuf.PutU16(0x200 + e.Definition().Id())
	} else {
		appearanceBuf.PutU16(0x100 + appearance.Model(protocol.BodyPartLegs))
	}

	if e := equipment.Slot(item.EquipmentHelm); e == nil && !helmCoversHair && !helmCoversFace {
		appearanceBuf.PutU16(0x100 + appearance.Model(protocol.BodyPartHead))
	} else {
		appearanceBuf.PutU8(0)
	}

	if e := equipment.Slot(item.EquipmentHands); e != nil {
		appearanceBuf.PutU16(0x200 + e.Definition().Id())
	} else {
		appearanceBuf.PutU16(0x100 + appearance.Model(protocol.BodyPartHands))
	}

	if e := equipment.Slot(item.EquipmentFeet); e != nil {
		appearanceBuf.PutU16(0x200 + e.Definition().Id())
	} else {
		appearanceBuf.PutU16(0x100 + appearance.Model(protocol.BodyPartFeet))
	}

	if e := equipment.Slot(item.EquipmentHelm); (e != nil && helmCoversFace) || appearance.Gender() == 1 {
		appearanceBuf.PutU8(0)
	} else {
		appearanceBuf.PutU16(0x100 + appearance.Model(protocol.BodyPartBeard))
	}

	for _, part := range []protocol.BodyPart{
		protocol.BodyPartHair,
		protocol.BodyPartTorso,
		protocol.BodyPartLegs,
		protocol.BodyPartFeet,
		protocol.BodyPartSkin} {
		appearanceBuf.PutU8(appearance.Color(part))
	}

	for _, anim := range []int{
		anims.Idle,
		anims.SpotRotate,
		anims.Walk,
		anims.Rotate180,
		anims.RotateCCW,
		anims.RotateCW,
		anims.Run} {
		appearanceBuf.PutU16(anim)
	}

	// Name
	appearanceBuf.PutStringZ(thisPlayer.Profile().Username())
	// Combat level
	appearanceBuf.PutU8(thisPlayer.Profile().Skills().CombatLevel())
	// Skill level
	appearanceBuf.PutU16(0)
	// Hidden
	appearanceBuf.PutU8(0)

	srcBlock := appearanceBuf.Bytes()
	block := make([]byte, len(srcBlock))
	for i := range srcBlock {
		block[i] = srcBlock[i] + 128
	}
	maskBuf.PutU8(len(block), encoding.IntOffset128)
	maskBuf.Write(block)
}

func (struc PlayerUpdate) getModifiedUpdateFlags(updatingPlayer protocol.Player) entity.Flags {
	// Clear some flags that don't apply to self updates
	flags := updatingPlayer.Flags()
	flags &= ^entity.MobFlagRegionUpdate

	return flags
}

func (struc PlayerUpdate) buildMovementBlock(buf *encoding.BitBuffer, player protocol.Player) {
	flags := struc.getModifiedUpdateFlags(player)

	switch {
	case (flags & entity.MobFlagRunUpdate) != 0:
		current, _ := player.WaypointQueue().WalkDirection()

		buf.Write(2, 2) // update type 2 = running
		buf.Write(4, uint32(current))

	case (flags & entity.MobFlagWalkUpdate) != 0:
		current, _ := player.WaypointQueue().WalkDirection()

		buf.Write(2, 1) // update type 1 = walking
		buf.Write(3, uint32(current))

	default:
		buf.Write(2, 0) // update type 0 = no movement updates
	}
}
