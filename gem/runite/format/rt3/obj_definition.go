package rt3

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
)

type ItemDefinition struct {
	Id            int
	Name          string
	Description   string
	Noted         bool
	Notable       bool
	Stackable     bool
	ParentId      int
	NotedId       int
	Members       bool
	ShopValue     int
	GroundActions []string
	Actions       []string
	StackIds      []int
	StackAmounts  []int
	Team          int

	ModelId        int
	ModelZoom      int
	ModelRotation1 int
	ModelRotation2 int
	ModelOffset1   int
	ModelOffset2   int
	PaletteSwaps   map[uint16]uint16
}

func LoadItemDefinitions(fs *JagFS) ([]*ItemDefinition, error) {
	idx, err := fs.Index(0)
	if err != nil {
		return nil, err
	}

	archiveFs := NewArchiveFS(idx)
	configArchive, err := archiveFs.Archive("config")
	if err != nil {
		return nil, err
	}

	objIdx, err := configArchive.File("obj.idx")
	if err != nil {
		return nil, err
	}

	objData, err := configArchive.File("obj.dat")
	if err != nil {
		return nil, err
	}

	idxBuf := bytes.NewReader(objIdx)

	var itemCount encoding.Uint16
	itemCount.Decode(idxBuf, encoding.IntegerFlag(encoding.IntNilFlag))

	items := make([]*ItemDefinition, int(itemCount))
	offset := 2
	var tmp encoding.Uint16
	for i := range items {
		tmp.Decode(idxBuf, encoding.IntegerFlag(encoding.IntNilFlag))

		dataBuf := bytes.NewReader(objData[offset:])
		items[i] = readItemDef(dataBuf)
		items[i].Id = i

		offset += int(tmp)
	}

	return items, nil
}

func newItemDefinition() *ItemDefinition {
	return &ItemDefinition{
		Actions:       make([]string, 5),
		GroundActions: make([]string, 5),
		PaletteSwaps:  make(map[uint16]uint16),
		StackIds:      make([]int, 10),
		StackAmounts:  make([]int, 10),
	}
}

func readItemDef(buf io.Reader) *ItemDefinition {
	nilFlags := encoding.IntegerFlag(encoding.IntNilFlag)

	definition := newItemDefinition()

	var tmp8 encoding.Uint8
	var tmp16 encoding.Uint16
	var tmp32 encoding.Uint32
	var tmpString encoding.JString
	for {
		tmp8.Decode(buf, nilFlags)

		section := int(tmp8)
		switch {
		case section == 0:
			return definition

		case section == 1:
			tmp16.Decode(buf, nilFlags)
			definition.ModelId = int(tmp16)

		case section == 2:
			tmpString.Decode(buf, 0)
			definition.Name = string(tmpString)

		case section == 3:
			tmpString.Decode(buf, 0)
			definition.Description = string(tmpString)

		case section == 4:
			tmp16.Decode(buf, nilFlags)
			definition.ModelZoom = int(tmp16)

		case section == 5:
			tmp16.Decode(buf, nilFlags)
			definition.ModelRotation1 = int(tmp16)

		case section == 6:
			tmp16.Decode(buf, nilFlags)
			definition.ModelRotation2 = int(tmp16)

		case section == 7:
			tmp16.Decode(buf, nilFlags)
			definition.ModelOffset1 = int(tmp16)
			if definition.ModelOffset1 > 32767 {
				definition.ModelOffset1 -= 0x10000
			}

		case section == 8:
			tmp16.Decode(buf, nilFlags)
			definition.ModelOffset2 = int(tmp16)
			if definition.ModelOffset2 > 32767 {
				definition.ModelOffset2 -= 0x10000
			}

		case section == 10:
			// unused
			tmp16.Decode(buf, nilFlags)

		case section == 11:
			definition.Stackable = true

		case section == 12:
			tmp32.Decode(buf, nilFlags)

			definition.ShopValue = int(tmp32)

		case section == 16:
			definition.Members = true

		case section == 23:
			//unknown
			tmp16.Decode(buf, nilFlags)

			tmp8.Decode(buf, nilFlags)

		case section == 24:
			//unknown
			tmp16.Decode(buf, nilFlags)

		case section == 25:
			//unknown
			tmp16.Decode(buf, nilFlags)

			tmp8.Decode(buf, nilFlags)

		case section == 26:
			//unknown
			tmp16.Decode(buf, nilFlags)

		case section >= 30 && section < 35:
			tmpString.Decode(buf, 0)

			i := section - 30
			definition.GroundActions[i] = string(tmpString)
			if definition.GroundActions[i] == "hidden" {
				definition.GroundActions[i] = ""
			}

		case section >= 35 && section < 40:
			tmpString.Decode(buf, 0)

			i := section - 35
			definition.Actions[i] = string(tmpString)

		case section == 40:
			tmp8.Decode(buf, nilFlags)

			numColorsSwapped := int(tmp8)
			for i := 0; i < numColorsSwapped; i++ {
				tmp16.Decode(buf, nilFlags)
				value := uint16(tmp16)

				tmp16.Decode(buf, nilFlags)
				key := uint16(tmp16)

				definition.PaletteSwaps[key] = value
			}

		case section == 78:
			//unknown
			tmp16.Decode(buf, nilFlags)

		case section == 79:
			//unknown
			tmp16.Decode(buf, nilFlags)

		case section == 90:
			//unknown
			tmp16.Decode(buf, nilFlags)

		case section == 91:
			//unknown
			tmp16.Decode(buf, nilFlags)

		case section == 92:
			//unknown
			tmp16.Decode(buf, nilFlags)

		case section == 93:
			//unknown
			tmp16.Decode(buf, nilFlags)

		case section == 94:
			//unknown
			tmp16.Decode(buf, nilFlags)

		case section == 95:
			//unknown
			tmp16.Decode(buf, nilFlags)

		case section == 96:
			//unknown
			tmp16.Decode(buf, nilFlags)

		case section == 97:
			tmp16.Decode(buf, nilFlags)
			definition.ParentId = int(tmp16)

		case section == 98:
			//unknown
			tmp16.Decode(buf, nilFlags)
			definition.NotedId = int(tmp16)

		case section >= 100 && section < 110:
			i := section - 100

			tmp16.Decode(buf, nilFlags)

			definition.StackIds[i] = int(tmp16)

			tmp16.Decode(buf, nilFlags)

			definition.StackAmounts[i] = int(tmp16)

		case section == 110:
			//unknown
			tmp16.Decode(buf, nilFlags)

		case section == 111:
			//unknown
			tmp16.Decode(buf, nilFlags)

		case section == 112:
			//unknown
			tmp16.Decode(buf, nilFlags)

		case section == 113:
			//unknown
			tmp8.Decode(buf, nilFlags)

		case section == 114:
			//unknown
			tmp8.Decode(buf, nilFlags)

		case section == 115:
			tmp8.Decode(buf, nilFlags)

			definition.Team = int(tmp8)

		default:
			panic(fmt.Sprintf("unknown section %v", section))
		}
	}
}
