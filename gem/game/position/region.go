package position

// A region is a 13x13 sector (104x104 tile) chunk.
// This is the loaded area of the map around the player.
type Region struct {
	// The origin of the region is the lowest corner sector, NOT the center.
	origin *Sector
}

func NewRegion(origin *Sector) *Region {
	if origin == nil {
		sector := NewSector(0, 0, 0)
		origin = sector
	}
	return &Region{origin}
}

func (region *Region) Compare(other *Region) bool {
	return region.Origin().Compare(other.Origin())
}

func (region *Region) Origin() *Sector {
	return region.origin
}

// Rebase adjusts the region such that it's new center is the sector containing the given Absolute
func (region *Region) Rebase(absolute *Absolute) {
	region.origin = NewSector(
		absolute.Sector().x-((RegionSize-1)/2),
		absolute.Sector().y-((RegionSize-1)/2),
		absolute.Sector().z,
	)
}

// SectorDelta calculates the delta between two regions in terms of sectors
func (region *Region) SectorDelta(other *Region) (x, y, z int) {
	return region.origin.Delta(other.origin)
}
