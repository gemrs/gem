package rt3

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
)

const (
	RegionSize = 64
)

type MapLandscape struct {
	Tiles [4][RegionSize][RegionSize]MapTile
}

type MapTile struct {
	RenderType int
}

func (m *MapLandscape) Decode(r io.Reader, flags_ interface{}) {
	buf := encoding.WrapReader(r)
	for z := 0; z < 4; z++ {
		for x := 0; x < RegionSize; x++ {
			for y := 0; y < RegionSize; y++ {
			L:
				for {
					opcode := buf.GetU8()
					if opcode == 0 {

					} else if opcode == 1 {
						buf.GetU8()
						break L
					} else if opcode <= 49 {
						buf.GetU8()
					} else if opcode <= 81 {
						m.Tiles[z][x][y].RenderType = opcode - 49
					}
				}
			}
		}
	}
}
