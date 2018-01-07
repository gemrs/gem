// NOT Generated by bbc (but should be in future)
package game

import (
	"bytes"
	"io"

	"github.com/gemrs/gem/gem/encoding"
	"github.com/gemrs/gem/gem/game/interface/entity"
	"github.com/gemrs/gem/gem/game/interface/player"
)

type PlayerUpdateBlock struct {
	OurPlayer player.Player
	buf       []byte
	err       error
}

func NewPlayerUpdateBlock(player player.Player) *PlayerUpdate {
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
		other := other.(player.Player)
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

		other := other.(player.Player)
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

func (struc *PlayerUpdateBlock) addPlayer(buf *encoding.BitBuffer, other player.Player) {
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

func (struc *PlayerUpdateBlock) buildMovementBlock(buf *encoding.BitBuffer, player player.Player) error {
	flags := player.Flags()

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

func (struc *PlayerUpdateBlock) buildUpdateBlock(w io.Writer, thisPlayer player.Player) error {
	flags := thisPlayer.Flags() & ^entity.MobFlagMovementUpdate
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

	/* Update appearance */
	if (flags & entity.MobFlagIdentityUpdate) != 0 {
		buf := encoding.NewBuffer()
		appearance := thisPlayer.Profile().Appearance()
		anims := thisPlayer.Animations()
		appearanceBlock := OutboundPlayerAppearance{
			Gender:   encoding.Uint8(appearance.Gender()),
			HeadIcon: encoding.Uint8(appearance.HeadIcon()),

			HelmModel:       encoding.Uint8(0),
			CapeModel:       encoding.Uint8(0),
			AmuletModel:     encoding.Uint8(0),
			RightWieldModel: encoding.Uint8(0),
			TorsoModel:      encoding.Uint16(256 + appearance.Model(player.Torso)),
			LeftWieldModel:  encoding.Uint8(0),
			ArmsModel:       encoding.Uint16(256 + appearance.Model(player.Arms)),
			LegsModel:       encoding.Uint16(256 + appearance.Model(player.Legs)),
			HeadModel:       encoding.Uint16(256 + appearance.Model(player.Head)),
			HandsModel:      encoding.Uint16(256 + appearance.Model(player.Hands)),
			FeetModel:       encoding.Uint16(256 + appearance.Model(player.Feet)),
			BeardModel:      encoding.Uint16(256 + appearance.Model(player.Beard)),

			HairColor:  encoding.Uint8(appearance.Color(player.Hair)),
			TorsoColor: encoding.Uint8(appearance.Color(player.Torso)),
			LegColor:   encoding.Uint8(appearance.Color(player.Legs)),
			FeetColor:  encoding.Uint8(appearance.Color(player.Feet)),
			SkinColor:  encoding.Uint8(appearance.Color(player.Skin)),

			AnimIdle:       encoding.Uint16(anims.Animation(player.AnimIdle)),
			AnimSpotRotate: encoding.Uint16(anims.Animation(player.AnimSpotRotate)),
			AnimWalk:       encoding.Uint16(anims.Animation(player.AnimWalk)),
			AnimRotate180:  encoding.Uint16(anims.Animation(player.AnimRotate180)),
			AnimRotateCCW:  encoding.Uint16(anims.Animation(player.AnimRotateCCW)),
			AnimRotateCW:   encoding.Uint16(anims.Animation(player.AnimRotateCW)),
			AnimRun:        encoding.Uint16(anims.Animation(player.AnimRun)),
		}

		err := appearanceBlock.Encode(buf, nil)
		if err != nil {
			return err
		}

		block := buf.Bytes()
		blockSize := encoding.Uint8(len(block))
		err = blockSize.Encode(w, encoding.IntNegate)
		if err != nil {
			return err
		}

		_, err = w.Write(block)
		if err != nil {
			return err
		}
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
