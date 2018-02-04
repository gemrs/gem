package rt3

import "fmt"

const maxRegion = 32768

type MapKeyLookupFunc func(region int) ([]uint32, bool)

type Map struct {
	Regions map[int]*Region
}

type Region struct {
	Region     int
	X, Y       int
	AbsX, AbsY int
	Locations  MapLocations
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

	count := 0

	for i := 0; i < maxRegion; i++ {
		region := NewRegion(i)

		locationFile := index.FileIndexByName(fmt.Sprintf("l%v_%v", region.X, region.Y))
		if locationFile == -1 {
			continue
		}

		mapKeys, ok := lookupKeys(i)
		if !ok {
			mapKeys = nil
		}

		locContainer, err := index.EncryptedContainer(locationFile, mapKeys)
		if err != nil {
			// Failed to load: keys are incorrect, or missing, or ..
			continue
		}

		region.Locations.Decode(locContainer, nil)
		count += len(region.Locations.Locations)
		m.Regions[i] = region
	}

	return nil
}
