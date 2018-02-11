package position

// Local is a coordinate relative to the base of a Region
type Local struct {
	x, y, z int
	// Region is the region which the coordinates are relative to
	region *Region
}

func NewLocal(x, y, z int, region *Region) *Local {
	return &Local{x, y, z, region}
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

func (local *Local) DeltaTo(target *Local) (x, y, z int) {
	return target.X() - local.X(), target.Y() - local.Y(), target.Z() - local.Z()
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
	abs := NewAbsolute(
		local.x+(SectorSize*local.Region().origin.x),
		local.y+(SectorSize*local.Region().origin.y),
		local.z,
	)
	return abs
}
