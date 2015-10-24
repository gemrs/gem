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

//go:generate gopygen -type Absolute -type Sector -type Region -type Local -excfield ".*" $GOFILE

// Positionable is an object which has an absolute position in the world
type Positionable interface {
	Position() *Absolute
	SetPosition(*Absolute)
}

// Absolute is a coordinate mapping to a single tile within the world
type Absolute struct {
	py.BaseObject

	x, y, z int
}

func (pos *Absolute) Init(x, y, z int) error {
	pos.x = x
	pos.y = y
	pos.z = z
	return nil
}

func (pos *Absolute) Compare(other *Absolute) bool {
	return pos.x == other.x &&
		pos.y == other.y &&
		pos.z == other.z
}

func (pos *Absolute) X() int {
	return pos.x
}

func (pos *Absolute) Y() int {
	return pos.y
}

func (pos *Absolute) Z() int {
	return pos.z
}

func (pos *Absolute) String() string {
	return fmt.Sprintf("Absolute<%v, %v, %v>", pos.x, pos.y, pos.z)
}

// Sector calculates the sector which contains an Absolute
func (pos *Absolute) Sector() *Sector {
	sector, err := NewSector(
		pos.x/AreaSize,
		pos.y/AreaSize,
		pos.z,
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
		pos.x-(AreaSize*region.origin.x),
		pos.y-(AreaSize*region.origin.y),
		pos.z,
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

	x, y, z int
}

func (s *Sector) Init(x, y, z int) error {
	s.x = x
	s.y = y
	s.z = z
	return nil
}

func (s *Sector) X() int {
	return s.x
}

func (s *Sector) Y() int {
	return s.y
}

func (s *Sector) Z() int {
	return s.z
}

func (s *Sector) Compare(other *Sector) bool {
	return s.x == other.x &&
		s.y == other.y &&
		s.z == other.z
}

// A region is a 13x13 sector (104x104 tile) chunk.
// This is the loaded area of the map around the player.
type Region struct {
	py.BaseObject

	// The origin of the region is the lowest corner sector, NOT the center.
	origin *Sector
}

func (region *Region) Init(origin *Sector) error {
	if origin == nil {
		sector, err := NewSector(0, 0, 0)
		if err != nil {
			return err
		}
		origin = sector
	}
	region.origin = origin
	return nil
}

func (region *Region) Compare(other *Region) bool {
	return region.Origin().Compare(other.Origin())
}

func (region *Region) Origin() *Sector {
	return region.origin
}

// Rebase adjusts the region such that it's new center is the sector containing the given Absolute
func (region *Region) Rebase(absolute *Absolute) {
	var err error
	region.origin, err = NewSector(
		absolute.Sector().x-((RegionSize-1)/2),
		absolute.Sector().y-((RegionSize-1)/2),
		absolute.Sector().z,
	)
	if err != nil {
		panic(err)
	}
}

// SectorDelta calculates the delta between two regions in terms of sectors
func (region *Region) SectorDelta(other *Region) (x, y, z int) {
	x = int(math.Abs(float64(other.origin.x - region.origin.x)))
	y = int(math.Abs(float64(other.origin.y - region.origin.y)))
	z = int(math.Abs(float64(other.origin.z - region.origin.z)))

	return x, y, z
}

// Local is a coordinate relative to the base of a Region
type Local struct {
	py.BaseObject

	x, y, z int
	// Region is the region which the coordinates are relative to
	region *Region
}

func (local *Local) Init(x, y, z int, region *Region) error {
	local.x = x
	local.y = y
	local.z = z
	local.region = region
	return nil
}

func (local *Local) X() int {
	return local.x
}

func (local *Local) Y() int {
	return local.y
}

func (local *Local) Z() int {
	return local.z
}

func (local *Local) Compare(other *Local) bool {
	return local.x == other.x &&
		local.y == other.y &&
		local.z == other.z &&
		local.Region().Compare(other.Region())
}

func (local *Local) Region() *Region {
	return local.region
}

// Absolute converts a local coordinate into an absolute coordinate
func (local *Local) Absolute() *Absolute {
	abs, err := NewAbsolute(
		local.x+(AreaSize*local.Region().origin.x),
		local.y+(AreaSize*local.Region().origin.y),
		local.z,
	)
	if err != nil {
		panic(err)
	}
	return abs
}
