package rt3

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
)

type MapLocations struct {
	Locations []MapLocation
}

type MapLocation struct {
	Id          int
	Type        int
	Orientation int
	LocalX      int
	LocalY      int
	LocalZ      int
}

func getSignedPackedInt(buf *encoding.Buffer) int {
	highByte := buf.GetU8()
	if highByte < 0x80 {
		return highByte
	}

	lowByte := buf.GetU8()
	return ((highByte << 8) + lowByte) - 0x8000

}

func (loc *MapLocations) Decode(r io.Reader, flags_ interface{}) {
	buf := encoding.WrapReader(r)

	id := -1
	delta := 0
	for {
		delta = getSignedPackedInt(buf.(*encoding.Buffer))
		if delta == 0 {
			break
		}

		id += delta

		position := 0
		positionDelta := 0

		for {
			positionDelta = getSignedPackedInt(buf.(*encoding.Buffer))
			if positionDelta == 0 {
				break
			}

			position += positionDelta - 1

			localY := position & 0x3F
			localX := position >> 6 & 0x3F
			height := position >> 12 & 0x3

			attributes := buf.GetU8()
			typ := attributes >> 2
			orientation := attributes & 0x3

			loc.Locations = append(loc.Locations, MapLocation{
				Id:          id,
				Type:        typ,
				Orientation: orientation,
				LocalX:      localX,
				LocalY:      localY,
				LocalZ:      height,
			})
		}
	}
}
