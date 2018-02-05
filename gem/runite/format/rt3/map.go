package rt3

import (
	"fmt"

	"github.com/gemrs/gem/gem/core/encoding"
)

const maxRegion = 32768

type MapKeyLookupFunc func(region int) ([]uint32, bool)

type Map struct {
	Regions    map[int]*Region
	MaxX, MaxY int
}

type Region struct {
	Region     int
	X, Y       int
	AbsX, AbsY int
	Locations  MapLocations
	Landscape  MapLandscape
}

func NewRegion(id int) *Region {
	return &Region{
		Region: id,
		X:      (id >> 8),
		Y:      id & 0xFF,
		AbsX:   (id >> 8 & 0xFF) << 6,
		AbsY:   (id & 0xFF) << 6,
	}
}

func (m *Map) Load(fs *JagFS, lookupKeys MapKeyLookupFunc) error {
	index, err := fs.Index(IdxLandscapes)
	if err != nil {
		return err
	}

	m.Regions = make(map[int]*Region)

	var maxX, maxY int

	for i := 0; i < maxRegion; i++ {
		region := NewRegion(i)

		mapFile := index.FileIndexByName(fmt.Sprintf("m%v_%v", region.X, region.Y))
		locationFile := index.FileIndexByName(fmt.Sprintf("l%v_%v", region.X, region.Y))
		if mapFile == -1 && locationFile == -1 {
			continue
		}

		m.Regions[i] = region

		if region.AbsX+64 > maxX {
			maxX = region.AbsX + 64
		}

		if region.AbsY+64 > maxY {
			maxY = region.AbsY + 64
		}

		mapKeys, ok := lookupKeys(i)
		if !ok {
			mapKeys = nil
		}

		if locationFile != -1 {
			locContainer, err := index.EncryptedContainer(locationFile, mapKeys)
			if err != nil {
				// Failed to load: keys are incorrect, or missing, or ..
				continue
			}

			encoding.TryDecode(&region.Locations, locContainer, nil)
		}

		if mapFile != -1 {
			mapContainer, err := index.Container(mapFile)
			if err != nil {
				// Failed to load: keys are incorrect, or missing, or ..
				continue
			}

			encoding.TryDecode(&region.Landscape, mapContainer, nil)
		}

	}

	m.MaxX = maxX
	m.MaxY = maxY

	return nil
}
