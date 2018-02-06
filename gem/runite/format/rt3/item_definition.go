package rt3

import (
	"fmt"
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
)

type ItemDefinition struct {
	Id               int
	Name             string
	Stackable        bool
	ShopValue        int
	GroundActions    [5]string
	InventoryActions [5]string
	Note             int
	NotedId          int
	UnnotedId        int
	NotedTemplate    int
}

func NewItemDefinition(id int) *ItemDefinition {
	return &ItemDefinition{
		Id:        id,
		Name:      "null",
		ShopValue: 1,
		GroundActions: [5]string{
			"",
			"",
			"Take",
			"",
			"",
		},
		InventoryActions: [5]string{
			"",
			"",
			"",
			"",
			"Drop",
		},
		NotedId:       -1,
		NotedTemplate: -1,
		UnnotedId:     -1,
	}
}

func (d *ItemDefinition) Decode(r io.Reader, flags_ interface{}) {
	buf := encoding.WrapReader(r)
L:
	for {
		opcode := buf.GetU8()

		switch opcode {
		case 0:
			break L

		case 1:
			buf.GetU16()

		case 2:
			d.Name = buf.GetStringZ()

		case 4:
			fallthrough
		case 5:
			fallthrough
		case 6:
			fallthrough
		case 7:
			fallthrough
		case 8:
			buf.GetU16()

		case 11:
			d.Stackable = true

		case 12:
			d.ShopValue = buf.GetU32()

		case 16:
		case 23:
			buf.GetU16()
			buf.GetU8()
		case 24:
			buf.GetU16()
		case 25:
			buf.GetU16()
			buf.GetU8()
		case 26:
			buf.GetU16()

		case 30:
			d.GroundActions[0] = buf.GetStringZ()
		case 31:
			d.GroundActions[1] = buf.GetStringZ()
		case 32:
			d.GroundActions[2] = buf.GetStringZ()
		case 33:
			d.GroundActions[3] = buf.GetStringZ()
		case 34:
			d.GroundActions[4] = buf.GetStringZ()

		case 35:
			d.InventoryActions[0] = buf.GetStringZ()
		case 36:
			d.InventoryActions[1] = buf.GetStringZ()
		case 37:
			d.InventoryActions[2] = buf.GetStringZ()
		case 38:
			d.InventoryActions[3] = buf.GetStringZ()
		case 39:
			d.InventoryActions[4] = buf.GetStringZ()

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

		case 42:
			buf.GetU8()

		case 65:

		case 78:
			fallthrough
		case 79:
			fallthrough
		case 90:
			fallthrough
		case 91:
			fallthrough
		case 92:
			fallthrough
		case 93:
			fallthrough
		case 95:
			buf.GetU16()

		case 97:
			d.Note = buf.GetU16()

		case 100:
			fallthrough
		case 101:
			fallthrough
		case 102:
			fallthrough
		case 103:
			fallthrough
		case 104:
			fallthrough
		case 105:
			fallthrough
		case 106:
			fallthrough
		case 107:
			fallthrough
		case 108:
			fallthrough
		case 109:
			buf.GetU16()
			buf.GetU16()

		case 110:
			fallthrough
		case 111:
			fallthrough
		case 112:
			buf.GetU16()

		case 113:
			fallthrough
		case 115:
			fallthrough
		case 114:
			buf.GetU8()

		case 139:
			d.UnnotedId = buf.GetU16()
		case 140:
			d.NotedId = buf.GetU16()

		case 148:
			fallthrough
		case 149:
			buf.GetU16()
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
			panic(fmt.Errorf("unrecognized item opcode"))
		}
	}
}
