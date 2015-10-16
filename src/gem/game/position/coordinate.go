package position

import (
	"github.com/qur/gopy/lib"
)

const (
	AreaSize   int = 8
	RegionSize int = 13
)

//go:generate gopygen -type Absolute -type Sector -type Region -type Local $GOFILE

// Locatable is an object which has an absolute position in the world
type Locatable interface {
	Position() *Absolute
}

// Absolute is a coordinate mapping to a single tile within the world
type Absolute struct {
	py.BaseObject

	X, Y, Z int
}

func NewAbsolute(x, y, z int) *Absolute {
	pos, err := Absolute{X: x, Y: y, Z: z}.Alloc()
	if err != nil {
		panic(err)
	}
	return pos
}

// Sector calculates the sector which contains an Absolute
func (pos *Absolute) Sector() *Sector {
	return NewSector(
		pos.X/AreaSize,
		pos.Y/AreaSize,
		pos.Z,
	)
}

// LocalTo calculates the local coordinates of an Absolute relative to a region
func (pos *Absolute) LocalTo(region *Region) *Local {
	return NewLocal(
		pos.X-(AreaSize*region.Origin.X),
		pos.Y-(AreaSize*region.Origin.Y),
		pos.Z,
		region,
	)
}

// A Sector is an 8x8 tile chunk of the map
type Sector struct {
	py.BaseObject

	X, Y, Z int
}

func NewSector(x, y, z int) *Sector {
	pos, err := Sector{X: x, Y: y, Z: z}.Alloc()
	if err != nil {
		panic(err)
	}
	return pos
}

// A region is a 13x13 sector (104x104 tile) chunk.
// This is the loaded area of the map around the player.
type Region struct {
	py.BaseObject

	// The Origin of the region is the lowest corner sector, NOT the center.
	Origin *Sector
}

func NewRegion(origin *Sector) *Region {
	pos, err := Region{Origin: origin}.Alloc()
	if err != nil {
		panic(err)
	}
	return pos
}

// Rebase adjusts the region such that it's new center is the sector containing the given Absolute
func (region *Region) Rebase(absolute *Absolute) {
	region.Origin = NewSector(
		absolute.Sector().X-((RegionSize-1)/2),
		absolute.Sector().Y-((RegionSize-1)/2),
		absolute.Sector().Z,
	)
}

// Local is a coordinate relative to the base of a Region
type Local struct {
	py.BaseObject

	X, Y, Z int
	// Region is the region which the coordinates are relative to
	Region *Region
}

func NewLocal(x, y, z int, region *Region) *Local {
	pos, err := Local{
		X: x, Y: y, Z: z,
		Region: region,
	}.Alloc()
	if err != nil {
		panic(err)
	}
	return pos
}

// Absolute converts a local coordinate into an absolute coordinate
func (local *Local) Absolute() *Absolute {
	return NewAbsolute(
		local.X+(AreaSize*local.Region.Origin.X),
		local.Y+(AreaSize*local.Region.Origin.Y),
		local.Z,
	)
}
