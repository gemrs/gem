package position

import (
	"fmt"
	"math"

	"github.com/qur/gopy/lib"
)

const (
	AreaSize   int = 8
	RegionSize int = 13
)

//go:generate gopygen -type Absolute -type Sector -type Region -type Local $GOFILE

// Positionable is an object which has an absolute position in the world
type Positionable interface {
	Position() *Absolute
	SetPosition(*Absolute)
}

// Absolute is a coordinate mapping to a single tile within the world
type Absolute struct {
	py.BaseObject

	X, Y, Z int
}

func (pos *Absolute) String() string {
	return fmt.Sprintf("Absolute<%v, %v, %v>", pos.X, pos.Y, pos.Z)
}

func (pos *Absolute) Init(x, y, z int) error {
	pos.X = x
	pos.Y = y
	pos.Z = z
	return nil
}

// Sector calculates the sector which contains an Absolute
func (pos *Absolute) Sector() *Sector {
	sector, err := NewSector(
		pos.X/AreaSize,
		pos.Y/AreaSize,
		pos.Z,
	)
	if err != nil {
		panic(err)
	}
	return sector
}

// RegionOf returns a Region centered on this position
func (pos *Absolute) RegionOf() *Region {
	region, err := NewRegion(nil)
	if err != nil {
		panic(err)
	}
	region.Rebase(pos)
	return region
}

// LocalTo calculates the local coordinates of an Absolute relative to a region
func (pos *Absolute) LocalTo(region *Region) *Local {
	local, err := NewLocal(
		pos.X-(AreaSize*region.Origin.X),
		pos.Y-(AreaSize*region.Origin.Y),
		pos.Z,
		region,
	)
	if err != nil {
		panic(err)
	}
	return local
}

// A Sector is an 8x8 tile chunk of the map
type Sector struct {
	py.BaseObject

	X, Y, Z int
}

func (s *Sector) Init(x, y, z int) error {
	s.X = x
	s.Y = y
	s.Z = z
	return nil
}

// A region is a 13x13 sector (104x104 tile) chunk.
// This is the loaded area of the map around the player.
type Region struct {
	py.BaseObject

	// The Origin of the region is the lowest corner sector, NOT the center.
	Origin *Sector
}

func (region *Region) Init(origin *Sector) error {
	if origin == nil {
		sector, err := NewSector(0, 0, 0)
		if err != nil {
			return err
		}
		origin = sector
	}
	region.Origin = origin
	return nil
}

// Rebase adjusts the region such that it's new center is the sector containing the given Absolute
func (region *Region) Rebase(absolute *Absolute) {
	var err error
	region.Origin, err = NewSector(
		absolute.Sector().X-((RegionSize-1)/2),
		absolute.Sector().Y-((RegionSize-1)/2),
		absolute.Sector().Z,
	)
	if err != nil {
		panic(err)
	}
}

// SectorDelta calculates the delta between two regions in terms of sectors
func (region *Region) SectorDelta(other *Region) (x, y, z int) {
	x = int(math.Abs(float64(other.Origin.X - region.Origin.X)))
	y = int(math.Abs(float64(other.Origin.Y - region.Origin.Y)))
	z = int(math.Abs(float64(other.Origin.Z - region.Origin.Z)))

	return x, y, z
}

// Local is a coordinate relative to the base of a Region
type Local struct {
	py.BaseObject

	X, Y, Z int
	// Region is the region which the coordinates are relative to
	Region *Region
}

func (local *Local) Init(x, y, z int, region *Region) error {
	local.X = x
	local.Y = y
	local.Z = z
	local.Region = region
	return nil
}

// Absolute converts a local coordinate into an absolute coordinate
func (local *Local) Absolute() *Absolute {
	abs, err := NewAbsolute(
		local.X+(AreaSize*local.Region.Origin.X),
		local.Y+(AreaSize*local.Region.Origin.Y),
		local.Z,
	)
	if err != nil {
		panic(err)
	}
	return abs
}
