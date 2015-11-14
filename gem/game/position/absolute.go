package position

import (
	"fmt"

	"github.com/qur/gopy/lib"
)

// Absolute is a coordinate mapping to a single tile within the world
type Absolute struct {
	py.BaseObject

	x, y, z int
}

func (pos *Absolute) Init(x, y, z int) {
	pos.x = x
	pos.y = y
	pos.z = z
}

func (pos *Absolute) Compare(other *Absolute) bool {
	return pos.x == other.x &&
		pos.y == other.y &&
		pos.z == other.z
}

func (pos *Absolute) Delta(target *Absolute) (x, y, z int) {
	return pos.X() - target.X(), pos.Y() - target.Y(), pos.Z() - target.Z()
}

// NextInterpolatedStep returns the one tile closer to the given point
// used by the waypoint queue to calculate the player's position along a path
func (pos *Absolute) NextInterpolatedPoint(target *Absolute) *Absolute {
	interp := func(a, b int) int {
		if a == b {
			return a
		} else if a > b {
			a--
			return a
		} else if a < b {
			a++
			return a
		}
		return 0
	}

	abs := NewAbsolute(interp(pos.X(), target.X()), interp(pos.Y(), target.Y()), pos.Z())

	return abs
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
	sector := NewSector(
		pos.x/AreaSize,
		pos.y/AreaSize,
		pos.z,
	)
	return sector
}

// RegionOf returns a Region centered on this position
func (pos *Absolute) RegionOf() *Region {
	region := NewRegion(nil)
	region.Rebase(pos)
	return region
}

// LocalTo calculates the local coordinates of an Absolute relative to a region
func (pos *Absolute) LocalTo(region *Region) *Local {
	local := NewLocal(
		pos.x-(AreaSize*region.origin.x),
		pos.y-(AreaSize*region.origin.y),
		pos.z,
		region,
	)
	return local
}
