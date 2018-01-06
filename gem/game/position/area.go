package position

// An Area is a rectangular grouping of sectors, with an arbitrary width and height.
type Area struct {
	sectors [][]*Sector
	w, h    int
}

func (area *Area) Init(origin *Sector, w, h int) {
	area.w, area.h = w, h
	area.sectors = make([][]*Sector, w)
	for x := 0; x < w; x++ {
		row := make([]*Sector, h)
		for y := 0; y < h; y++ {
			row[y] = NewSector(origin.x+x, origin.y+y, origin.z)
		}
		area.sectors[x] = row
	}
}

// MinSector returns the minimum sector of this area
func (area *Area) MinSector() *Sector {
	return area.sectors[0][0]
}

// MaxSector returns the maximum sector of this area
func (area *Area) MaxSector() *Sector {
	return area.sectors[area.w-1][area.h-1]
}

// Min returns the absolute coord of the minimum tile in this area
func (area *Area) Min() *Absolute {
	return area.MinSector().Min()
}

// Max returns the absolute coord of the maximum tile in this area
func (area *Area) Max() *Absolute {
	return area.MaxSector().Max()
}

// Contains determines if a given tile is within this area.
func (area *Area) Contains(t *Absolute) bool {
	p1, p2 := area.Min(), area.Max()

	return t.X() >= p1.X() && t.Y() >= p1.Y() && t.Z() >= p1.Z() &&
		t.X() <= p2.X() && t.Y() <= p2.Y() && t.Z() <= p2.Z()
}
