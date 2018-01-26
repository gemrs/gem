// Code generated by bbc; DO NOT EDIT.
package protocol_317

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
)

type OutboundPlayerAppearance struct {
	Gender          encoding.Uint8
	HeadIcon        encoding.Uint8
	HelmModel       encoding.Uint8
	CapeModel       encoding.Uint8
	AmuletModel     encoding.Uint8
	RightWieldModel encoding.Uint8
	TorsoModel      encoding.Uint16
	LeftWieldModel  encoding.Uint8
	ArmsModel       encoding.Uint16
	LegsModel       encoding.Uint16
	HeadModel       encoding.Uint16
	HandsModel      encoding.Uint16
	FeetModel       encoding.Uint16
	BeardModel      encoding.Uint16
	HairColor       encoding.Uint8
	TorsoColor      encoding.Uint8
	LegColor        encoding.Uint8
	FeetColor       encoding.Uint8
	SkinColor       encoding.Uint8
	AnimIdle        encoding.Uint16
	AnimSpotRotate  encoding.Uint16
	AnimWalk        encoding.Uint16
	AnimRotate180   encoding.Uint16
	AnimRotateCCW   encoding.Uint16
	AnimRotateCW    encoding.Uint16
	AnimRun         encoding.Uint16
	NameHash        encoding.Uint64
	CombatLevel     encoding.Uint8
	SkillLevel      encoding.Uint16
}

func (struc *OutboundPlayerAppearance) Encode(buf io.Writer, flags interface{}) {
	struc.Gender.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.HeadIcon.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.HelmModel.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.CapeModel.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.AmuletModel.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.RightWieldModel.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.TorsoModel.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.LeftWieldModel.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.ArmsModel.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.LegsModel.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.HeadModel.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.HandsModel.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.FeetModel.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.BeardModel.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.HairColor.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.TorsoColor.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.LegColor.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.FeetColor.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.SkinColor.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.AnimIdle.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.AnimSpotRotate.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.AnimWalk.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.AnimRotate180.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.AnimRotateCCW.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.AnimRotateCW.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.AnimRun.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.NameHash.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.CombatLevel.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.SkillLevel.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

}

func (struc *OutboundPlayerAppearance) Decode(buf io.Reader, flags interface{}) {
	struc.Gender.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.HeadIcon.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.HelmModel.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.CapeModel.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.AmuletModel.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.RightWieldModel.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.TorsoModel.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.LeftWieldModel.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.ArmsModel.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.LegsModel.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.HeadModel.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.HandsModel.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.FeetModel.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.BeardModel.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.HairColor.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.TorsoColor.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.LegColor.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.FeetColor.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.SkinColor.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.AnimIdle.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.AnimSpotRotate.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.AnimWalk.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.AnimRotate180.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.AnimRotateCCW.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.AnimRotateCW.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.AnimRun.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.NameHash.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.CombatLevel.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	struc.SkillLevel.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

}