//glua:bind module gem.game.position
package position

import (
	"encoding/json"
	"fmt"
)

//go:generate glua .

// Absolute is a coordinate mapping to a single tile within the world
//glua:bind
type Absolute struct {
	jsonAbsolute
}

type jsonAbsolute struct {
	X, Y, Z int
}

//glua:bind constructor Absolute
func NewAbsolute(x, y, z int) *Absolute {
	return &Absolute{jsonAbsolute{x, y, z}}
}

func (pos *Absolute) MarshalJSON() ([]byte, error) {
	return json.Marshal(pos.jsonAbsolute)
}

func (pos *Absolute) UnmarshalJSON(d []byte) error {
	return json.Unmarshal(d, &pos.jsonAbsolute)
}

func (pos *Absolute) Compare(other *Absolute) bool {
	return pos.jsonAbsolute.X == other.jsonAbsolute.X &&
		pos.jsonAbsolute.Y == other.jsonAbsolute.Y &&
		pos.jsonAbsolute.Z == other.jsonAbsolute.Z
}

func (pos *Absolute) Delta(target *Absolute) (x, y, z int) {
	return pos.X() - target.X(), pos.Y() - target.Y(), pos.Z() - target.Z()
}

func (pos *Absolute) DeltaTo(target *Absolute) (x, y, z int) {
	return target.X() - pos.X(), target.Y() - pos.Y(), target.Z() - pos.Z()
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

//glua:bind
func (pos *Absolute) X() int {
	return pos.jsonAbsolute.X
}

//glua:bind
func (pos *Absolute) Y() int {
	return pos.jsonAbsolute.Y
}

//glua:bind
func (pos *Absolute) Z() int {
	return pos.jsonAbsolute.Z
}

func (pos *Absolute) String() string {
	return fmt.Sprintf("Absolute<%v, %v, %v>", pos.X(), pos.Y(), pos.Z())
}

// Sector calculates the sector which contains an Absolute
func (pos *Absolute) Sector() *Sector {
	sector := NewSector(
		pos.X()/SectorSize,
		pos.Y()/SectorSize,
		pos.Z(),
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
		pos.X()-(SectorSize*region.origin.x),
		pos.Y()-(SectorSize*region.origin.y),
		pos.Z(),
		region,
	)
	return local
}

func (pos *Absolute) SectorLocal() (x, y int) {
	return pos.X() % SectorSize, pos.Y() % SectorSize
}
