package position

// Local is a coordinate relative to the base of a Region
type Local struct {
	x, y, z int
	// Region is the region which the coordinates are relative to
	region *Region
}

func (local *Local) Init(x, y, z int, region *Region) {
	local.x = x
	local.y = y
	local.z = z
	local.region = region
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
	abs := NewAbsolute(
		local.x+(AreaSize*local.Region().origin.x),
		local.y+(AreaSize*local.Region().origin.y),
		local.z,
	)
	return abs
}
