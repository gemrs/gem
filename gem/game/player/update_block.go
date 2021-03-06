package player

// this file is in this package rather than game_protocol because it depends
// heavily on the structures in here.. (ugh)

import (
	"bytes"
	"io"

	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/item"
)

type PlayerUpdateBlock struct {
	OurPlayer *Player
	buf       []byte
	err       error
}

func NewPlayerUpdateBlock(player *Player) *PlayerUpdate {
	block := &PlayerUpdateBlock{
		OurPlayer: player,
	}
	block.err = block.build()
	return (*PlayerUpdate)(block)
}

func (block *PlayerUpdateBlock) Encode(w io.Writer, flags interface{}) error {
	w.Write(block.buf)
	return block.err
}

func (struc *PlayerUpdateBlock) build() error {
	var byteBuffer bytes.Buffer
	w := &byteBuffer
	buf := encoding.NewBitBuffer(w)

	updateBlock := encoding.NewBuffer()

	// Update our player
	struc.buildUpdateBlock(updateBlock, struc.OurPlayer)
	err := struc.buildMovementBlock(buf, struc.OurPlayer)
	if err != nil {
		return err
	}

	visibleEntities := struc.OurPlayer.VisibleEntities()
	ourIndex := struc.OurPlayer.Index()

	// Update known players
	updatingPlayers := visibleEntities.Entities().Clone()
	updatingPlayers.RemoveAll(visibleEntities.Adding())
	updatingPlayers = updatingPlayers.Filter(entity.PlayerType)
	updatingPlayers.Remove(struc.OurPlayer)

	buf.Write(8, uint32(updatingPlayers.Size())) // count of other players to update
	for _, other := range updatingPlayers.Slice() {
		other := other.(*Player)
		if visibleEntities.Removing().Contains(other) {
			buf.Write(1, 1)
			buf.Write(2, 3)
		} else {
			struc.buildMovementBlock(buf, other)
			if other.Flags() != 0 {
				struc.buildUpdateBlock(updateBlock, other)
			}
		}
	}

	// Add new players
	for _, other := range visibleEntities.Adding().Filter(entity.PlayerType).Slice() {
		if ourIndex == other.Index() {
			continue
		}

		other := other.(*Player)
		struc.addPlayer(buf, other)

		// Force appearance update
		other.SetFlags(other.Flags() | entity.MobFlagIdentityUpdate)
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

	struc.buf = byteBuffer.Bytes()

	return nil
}

func (struc *PlayerUpdateBlock) addPlayer(buf *encoding.BitBuffer, other *Player) {
	buf.Write(11, uint32(other.Index()))

	buf.Write(1, 1) // Update required
	buf.Write(1, 1) // Discard walk queue

	deltaX, deltaY, _ := other.Position().Delta(struc.OurPlayer.Position())
	buf.Write(5, uint32(deltaY))
	buf.Write(5, uint32(deltaX))
}

func (struc *PlayerUpdateBlock) Decode(buf io.Reader, flags interface{}) (err error) {
	panic("not implemented")
}

func (struc *PlayerUpdateBlock) getModifiedUpdateFlags(updatingPlayer *Player) entity.Flags {
	// Clear some flags that don't apply to self updates
	updatingThisPlayer := struc.OurPlayer.Index() == updatingPlayer.Index()
	flags := updatingPlayer.Flags()
	if updatingThisPlayer {
		flags = flags & ^entity.MobFlagChatUpdate
	}
	return flags
}

func (struc *PlayerUpdateBlock) buildMovementBlock(buf *encoding.BitBuffer, player *Player) error {
	flags := struc.getModifiedUpdateFlags(player)

	// Anything to do?
	if flags == 0 {
		buf.Write(1, 0) // No updates
		return nil
	}
	buf.Write(1, 1) // This player has updates

	// Do we have any non-movement updates to perform?
	otherUpdateFlags := (flags & ^entity.MobFlagMovementUpdate)

	updatingThisPlayer := struc.OurPlayer.Index() == player.Index()

	switch {
	// When updating other players, don't send warp movements
	case updatingThisPlayer && (flags&entity.MobFlagRegionUpdate) != 0:
		localPos := player.Position().LocalTo(struc.OurPlayer.LoadedRegion())

		buf.Write(2, 3) // update type 3 = warp to location
		buf.Write(2, uint32(localPos.Z()))
		buf.WriteBit(false) // discard walk queue? not sure when/if we need this
		buf.WriteBit(otherUpdateFlags != 0)
		buf.Write(7, uint32(localPos.Y()))
		buf.Write(7, uint32(localPos.X()))

	case (flags & entity.MobFlagRunUpdate) != 0:
		current, last := player.WaypointQueue().WalkDirection()

		buf.Write(2, 2) // update type 2 = running
		buf.Write(3, uint32(last))
		buf.Write(3, uint32(current))
		buf.WriteBit(otherUpdateFlags != 0)

	case (flags & entity.MobFlagWalkUpdate) != 0:
		current, _ := player.WaypointQueue().WalkDirection()

		buf.Write(2, 1) // update type 1 = walking
		buf.Write(3, uint32(current))
		buf.WriteBit(otherUpdateFlags != 0)

	default:
		buf.Write(2, 0) // update type 0 = no movement updates
	}
	return nil
}

func (struc *PlayerUpdateBlock) buildChatUpdateBlock(w io.Writer, other *Player) error {
	message := other.ChatMessageQueue()[0]
	chatBlock := &game_protocol.OutboundPlayerChatMessage{
		Effects:       encoding.Uint8(message.Effects),
		Colour:        encoding.Uint8(message.Colour),
		Rights:        encoding.Uint8(other.Profile().Rights()),
		Length:        encoding.Uint8(len(message.PackedMessage)),
		PackedMessage: message.PackedMessage,
	}

	return chatBlock.Encode(w, nil)
}

func (struc *PlayerUpdateBlock) buildUpdateBlock(w io.Writer, thisPlayer *Player) error {
	flags := struc.getModifiedUpdateFlags(thisPlayer) & ^entity.MobFlagMovementUpdate

	if flags == 0 {
		return nil
	}

	if flags >= 256 {
		flags |= 64
		flagsEnc := encoding.Uint16(flags)
		err := flagsEnc.Encode(w, encoding.IntLittleEndian)
		if err != nil {
			return err
		}
	} else {
		flagsEnc := encoding.Uint8(flags)
		err := flagsEnc.Encode(w, encoding.IntNilFlag)
		if err != nil {
			return err
		}
	}

	/* Update chat */
	if (flags & entity.MobFlagChatUpdate) != 0 {
		err := struc.buildChatUpdateBlock(w, thisPlayer)
		if err != nil {
			return err
		}
	}

	/* Update appearance */
	if (flags & entity.MobFlagIdentityUpdate) != 0 {
		err := struc.buildAppearanceUpdateBlock(w, thisPlayer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (struc *PlayerUpdateBlock) buildAppearanceUpdateBlock(w io.Writer, thisPlayer *Player) error {
	buf := encoding.NewBuffer()
	appearance := thisPlayer.Profile().Appearance()
	equipment := thisPlayer.Profile().Equipment()

	var weapon int
	if !equipment.SlotPopulated(item.EquipmentWeapon) {
		weapon = 0
	} else {
		weapon = equipment.Slot(item.EquipmentWeapon).Definition().Id()
	}

	anims := gamedata.Weapons[weapon].CharAnimations()

	equipmentId := func(slot int) (int, bool) {
		if !equipment.SlotPopulated(slot) {
			return 0, false
		}
		return equipment.Slot(slot).Definition().Id(), true
	}

	equipmentDef := func(slot int) (gamedata.EquipmentDef, bool) {
		id, ok := equipmentId(slot)
		if !ok {
			return gamedata.EquipmentDef{}, false
		}

		return gamedata.Equipment[id], true
	}

	var torsoCoversArms, helmCoversFace, helmCoversHair bool
	torso, ok := equipmentDef(item.EquipmentTorso)
	if ok {
		torsoCoversArms = torso.CoversArms()
	}

	helm, ok := equipmentDef(item.EquipmentHelm)
	if ok {
		helmCoversFace = helm.CoversFace()
		helmCoversHair = helm.CoversHair()
	}

	fields := make([]encoding.Encodable, 0)

	fields = append(fields, encoding.Uint8(appearance.Gender()))
	fields = append(fields, encoding.Uint8(appearance.HeadIcon()))

	for _, slot := range []int{item.EquipmentHelm, item.EquipmentCape, item.EquipmentAmulet, item.EquipmentWeapon} {
		if id, ok := equipmentId(slot); ok {
			fields = append(fields, encoding.Uint16(0x200+id))
		} else {
			fields = append(fields, encoding.Uint8(0))
		}
	}

	if id, ok := equipmentId(item.EquipmentTorso); ok {
		fields = append(fields, encoding.Uint16(0x200+id))
	} else {
		fields = append(fields, encoding.Uint16(0x100+appearance.Model(BodyPartTorso)))
	}

	if id, ok := equipmentId(item.EquipmentShield); ok {
		fields = append(fields, encoding.Uint16(0x200+id))
	} else {
		fields = append(fields, encoding.Uint8(0))
	}

	if _, ok := equipmentId(item.EquipmentTorso); ok && !torsoCoversArms {
		fields = append(fields, encoding.Uint8(0))
	} else {
		fields = append(fields, encoding.Uint16(0x100+appearance.Model(BodyPartArms)))
	}

	if id, ok := equipmentId(item.EquipmentLegs); ok {
		fields = append(fields, encoding.Uint16(0x200+id))
	} else {
		fields = append(fields, encoding.Uint16(0x100+appearance.Model(BodyPartLegs)))
	}

	if _, ok := equipmentId(item.EquipmentHelm); !ok && !helmCoversHair && !helmCoversFace {
		fields = append(fields, encoding.Uint16(0x100+appearance.Model(BodyPartHead)))
	} else {
		fields = append(fields, encoding.Uint8(0))
	}

	if id, ok := equipmentId(item.EquipmentHands); ok {
		fields = append(fields, encoding.Uint16(0x200+id))
	} else {
		fields = append(fields, encoding.Uint16(0x100+appearance.Model(BodyPartHands)))
	}

	if id, ok := equipmentId(item.EquipmentFeet); ok {
		fields = append(fields, encoding.Uint16(0x200+id))
	} else {
		fields = append(fields, encoding.Uint16(0x100+appearance.Model(BodyPartFeet)))
	}

	if _, ok := equipmentId(item.EquipmentHelm); (ok && helmCoversFace) || appearance.Gender() == 1 {
		fields = append(fields, encoding.Uint8(0))
	} else {
		fields = append(fields, encoding.Uint16(0x100+appearance.Model(BodyPartBeard)))
	}

	for _, part := range []BodyPart{BodyPartHair, BodyPartTorso, BodyPartLegs, BodyPartFeet, BodyPartSkin} {
		fields = append(fields, encoding.Uint8(appearance.Color(part)))
	}

	for _, anim := range []int{anims.Idle, anims.SpotRotate, anims.Walk,
		anims.Rotate180, anims.RotateCCW, anims.RotateCW, anims.Run} {
		fields = append(fields, encoding.Uint16(anim))
	}

	fields = append(fields, encoding.NameHash(thisPlayer.Profile().Username()))
	fields = append(fields, encoding.Uint8(thisPlayer.Profile().Skills().CombatLevel()))
	fields = append(fields, encoding.Uint16(0))

	for _, field := range fields {
		err := field.Encode(buf, nil)
		if err != nil {
			return err
		}
	}

	block := buf.Bytes()
	blockSize := encoding.Uint8(len(block))
	err := blockSize.Encode(w, encoding.IntNegate)
	if err != nil {
		return err
	}

	_, err = w.Write(block)
	if err != nil {
		return err
	}

	return nil
}

type PlayerUpdate PlayerUpdateBlock

var PlayerUpdateDefinition = encoding.PacketHeader{
	Type:   (*PlayerUpdate)(nil),
	Number: 81,
	Size:   encoding.SzVar16,
}

func (frm *PlayerUpdate) Encode(buf io.Writer, flags interface{}) (err error) {
	struc := (*PlayerUpdateBlock)(frm)
	hdr := encoding.PacketHeader{
		Number: PlayerUpdateDefinition.Number,
		Size:   PlayerUpdateDefinition.Size,
		Object: struc,
	}
	return hdr.Encode(buf, flags)
}

func (frm *PlayerUpdate) Decode(buf io.Reader, flags interface{}) (err error) {
	struc := (*PlayerUpdateBlock)(frm)
	hdr := encoding.PacketHeader{
		Number: PlayerUpdateDefinition.Number,
		Size:   PlayerUpdateDefinition.Size,
		Object: struc,
	}
	return hdr.Decode(buf, flags)
}
