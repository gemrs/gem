package position

import (
	"math"
)

// A SectorHash is a 64-bit representation of a given sector in the world.
// This is useful for referencing sectors in a map for fast lookup
type SectorHash uint64

func (s SectorHash) Sector() *Sector {
	x := (s >> 0) & 0xFFFF
	y := (s >> 16) & 0xFFFF
	z := (s >> 32) & 0xFFFF
	return NewSector(int(x), int(y), int(z))
}

// A Sector is an 8x8 tile chunk of the map
type Sector struct {
	x, y, z int
}

func NewSector(x, y, z int) *Sector {
	return &Sector{x, y, z}
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

// Min returns the absolute coord of the minimum tile in this sector
func (s *Sector) Min() *Absolute {
	return NewAbsolute(
		s.x*AreaSize,
		s.y*AreaSize,
		s.z,
	)
}

// Max returns the absolute coord of the maximum tile in this sector
func (s *Sector) Max() *Absolute {
	return NewAbsolute(
		(s.x+1)*AreaSize,
		(s.y+1)*AreaSize,
		s.z,
	)
}

func (s *Sector) HashCode() SectorHash {
	x, y, z := uint16(s.X()), uint16(s.Y()), uint16(s.Z())
	hash := SectorHash(x)
	hash = hash + SectorHash(uint64(y)<<16)
	hash = hash + SectorHash(uint64(z)<<32)
	return hash
}

func (s *Sector) Compare(other *Sector) bool {
	return s.x == other.x &&
		s.y == other.y &&
		s.z == other.z
}

func (s *Sector) Delta(other *Sector) (x, y, z int) {
	x = int(math.Abs(float64(other.x - s.x)))
	y = int(math.Abs(float64(other.y - s.y)))
	z = int(math.Abs(float64(other.z - s.z)))

	return x, y, z
}
