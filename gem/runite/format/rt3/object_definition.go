package rt3

import (
	"fmt"
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
)

type ObjectDefinition struct {
	Id               int
	Name             string
	SizeX            int
	SizeY            int
	OffsetX          int
	OffsetY          int
	OffsetZ          int
	Actions          [5]string
	Solid            bool
	BlocksProjectile bool
	InteractType     int
	ModelSizeX       int
	ModelSizeY       int
	ModelSizeZ       int
	MapSceneId       int
	MapAreaId        int
}

func NewObjectDefinition(id int) *ObjectDefinition {
	return &ObjectDefinition{
		Id: id,
	}
}

func (d *ObjectDefinition) Decode(r io.Reader, flags_ interface{}) {
	buf := encoding.WrapReader(r)
L:
	for {
		opcode := buf.GetU8()

		switch opcode {
		case 0:
			break L

		case 1:
			len := buf.GetU8()
			for i := 0; i < len; i++ {
				buf.GetU16()
				buf.GetU8()
			}

		case 2:
			d.Name = buf.GetStringZ()

		case 5:
			len := buf.GetU8()
			for i := 0; i < len; i++ {
				buf.GetU16()
			}

		case 14:
			d.SizeX = buf.GetU8()

		case 15:
			d.SizeY = buf.GetU8()

		case 17:
			d.InteractType = 0
			d.BlocksProjectile = false

		case 18:
			d.BlocksProjectile = false

		case 19:
			buf.GetU8()

		case 21:
		case 22:
		case 23:
		case 24:
			buf.GetU16()

		case 27:
		case 28:
			buf.GetU8()

		case 29:
			buf.Get8()

		case 39:
			buf.Get8()

		case 30:
			d.Actions[0] = buf.GetStringZ()

		case 31:
			d.Actions[1] = buf.GetStringZ()

		case 32:
			d.Actions[2] = buf.GetStringZ()

		case 33:
			d.Actions[3] = buf.GetStringZ()

		case 34:
			d.Actions[4] = buf.GetStringZ()

		case 40:
			len := buf.GetU8()
			for i := 0; i < len; i++ {
				buf.GetU16()
				buf.GetU16()
			}

		case 41:
			len := buf.GetU8()
			for i := 0; i < len; i++ {
				buf.GetU16()
				buf.GetU16()
			}

		case 62:
		case 64:
		case 65:
			d.ModelSizeX = buf.GetU16()

		case 66:
			d.ModelSizeZ = buf.GetU16()

		case 67:
			d.ModelSizeY = buf.GetU16()

		case 68:
			d.MapSceneId = buf.GetU16()

		case 69:
			buf.GetU8()

		case 70:
			d.OffsetX = buf.GetU16()

		case 71:
			d.OffsetZ = buf.GetU16()

		case 72:
			d.OffsetY = buf.GetU16()

		case 73:
		case 74:
		case 75:
			buf.GetU8()

		case 77:
			buf.GetU16()
			buf.GetU16()
			len := buf.GetU8()
			for i := 0; i <= len; i++ {
				buf.GetU16()
			}

		case 78:
			buf.GetU16()
			buf.GetU8()

		case 79:
			buf.GetU16()
			buf.GetU16()
			buf.GetU8()
			len := buf.GetU8()
			for i := 0; i < len; i++ {
				buf.GetU16()
			}

		case 81:
			buf.GetU8()

		case 82:
			d.MapAreaId = buf.GetU16()

		case 92:
			buf.GetU16()
			buf.GetU16()
			buf.GetU16()
			len := buf.GetU8()
			for i := 0; i <= len; i++ {
				buf.GetU16()
			}

		case 249:
			len := buf.GetU8()
			for i := 0; i < len; i++ {
				isStr := buf.GetU8() == 1
				buf.GetU24()
				if isStr {
					buf.GetStringZ()
				} else {
					buf.GetU32()
				}
			}

		default:
			panic(fmt.Errorf("unrecognized object opcode"))
		}
	}
}
