package player

import (
	"github.com/qur/gopy/lib"
)

//go:generate gopygen -type Appearance -type Animations $GOFILE

type Appearance struct {
	py.BaseObject

	Gender   int
	HeadIcon int

	TorsoModel int
	ArmsModel  int
	LegsModel  int
	HeadModel  int
	HandsModel int
	FeetModel  int
	BeardModel int

	HairColor  int
	TorsoColor int
	LegColor   int
	FeetColor  int
	SkinColor  int
}

func (a *Appearance) Init() error {
	a.Gender = 0
	a.HeadIcon = 0

	a.TorsoModel = 19
	a.ArmsModel = 29
	a.LegsModel = 39
	a.HeadModel = 3
	a.HandsModel = 35
	a.FeetModel = 44
	a.BeardModel = 10

	a.HairColor = 7
	a.TorsoColor = 8
	a.LegColor = 9
	a.FeetColor = 5
	a.SkinColor = 0
	return nil
}

type Animations struct {
	py.BaseObject

	AnimIdle       int
	AnimSpotRotate int
	AnimWalk       int
	AnimRotate180  int
	AnimRotateCCW  int
	AnimRotateCW   int
	AnimRun        int
}

func (a *Animations) Init() error {
	a.AnimIdle = 0x328
	a.AnimSpotRotate = 0x337
	a.AnimWalk = 0x333
	a.AnimRotate180 = 0x334
	a.AnimRotateCCW = 0x335
	a.AnimRotateCW = 0x336
	a.AnimRun = 0x338
	return nil
}
