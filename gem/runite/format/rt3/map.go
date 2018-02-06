package rt3

import (
	"errors"
	"fmt"

	"github.com/gemrs/gem/gem/core/encoding"
)

var (
	ErrNoMap = errors.New("unable to load requested map")
)

const maxRegion = 32768

type MapKeyLookupFunc func(region int) ([]uint32, bool)

type Map struct {
	regions    map[int]*Region
	index      *JagFSIndex
	lookupKeys MapKeyLookupFunc
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

func (m *Map) Region(x, y int) (*Region, error) {
	return m.RegionById((x << 8) + y)
}

func (m *Map) RegionById(id int) (*Region, error) {
	if region, ok := m.regions[id]; ok {
		return region, nil
	}

	err := m.loadRegion(id)
	if err != nil {
		return nil, err
	}

	return m.regions[id], nil
}

func (m *Map) NumRegions() int {
	return m.index.FileCount()
}

func (m *Map) Load(fs *JagFS, lookupKeys MapKeyLookupFunc) error {
	var err error
	m.index, err = fs.Index(IdxLandscapes)
	if err != nil {
		return err
	}

	m.lookupKeys = lookupKeys
	m.regions = make(map[int]*Region)

	return nil
}

func (m *Map) loadRegion(id int) error {
	region := NewRegion(id)

	mapFile := m.index.FileIndexByName(fmt.Sprintf("m%v_%v", region.X, region.Y))
	locationFile := m.index.FileIndexByName(fmt.Sprintf("l%v_%v", region.X, region.Y))
	if mapFile == -1 && locationFile == -1 {
		return ErrNoMap
	}

	mapKeys, ok := m.lookupKeys(id)
	if !ok {
		mapKeys = nil
	}

	if locationFile != -1 {
		locContainer, err := m.index.EncryptedContainer(locationFile, mapKeys)
		if err != nil {
			// Failed to load: keys are incorrect, or missing, or ..
			return ErrNoMap
		}

		encoding.TryDecode(&region.Locations, locContainer, nil)
	}

	if mapFile != -1 {
		mapContainer, err := m.index.Container(mapFile)
		if err != nil {
			// Failed to load: keys are incorrect, or missing, or ..
			return ErrNoMap
		}

		encoding.TryDecode(&region.Landscape, mapContainer, nil)
	}

	m.regions[id] = region
	return nil
}
