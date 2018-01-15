package rt3

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gemrs/gem/gem/encoding"
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
	err = itemCount.Decode(idxBuf, encoding.IntegerFlag(encoding.IntNilFlag))
	if err != nil {
		return nil, err
	}

	items := make([]*ItemDefinition, int(itemCount))
	offset := 2
	var tmp encoding.Uint16
	for i := range items {
		err = tmp.Decode(idxBuf, encoding.IntegerFlag(encoding.IntNilFlag))
		if err != nil {
			return nil, err
		}

		dataBuf := bytes.NewReader(objData[offset:])
		items[i], err = readItemDef(dataBuf)
		items[i].Id = i
		if err != nil {
			return nil, err
		}

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

func readItemDef(buf io.Reader) (*ItemDefinition, error) {
	nilFlags := encoding.IntegerFlag(encoding.IntNilFlag)

	definition := newItemDefinition()

	var tmp8 encoding.Uint8
	var tmp16 encoding.Uint16
	var tmp32 encoding.Uint32
	var tmpString encoding.JString
	var err error
	for {
		err = tmp8.Decode(buf, nilFlags)
		if err != nil {
			return nil, err
		}

		section := int(tmp8)
		switch {
		case section == 0:
			return definition, nil

		case section == 1:
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

			definition.ModelId = int(tmp16)

		case section == 2:
			err = tmpString.Decode(buf, 0)
			if err != nil {
				return nil, err
			}

			definition.Name = string(tmpString)

		case section == 3:
			err = tmpString.Decode(buf, 0)
			if err != nil {
				return nil, err
			}

			definition.Description = string(tmpString)

		case section == 4:
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

			definition.ModelZoom = int(tmp16)

		case section == 5:
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

			definition.ModelRotation1 = int(tmp16)

		case section == 6:
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

			definition.ModelRotation2 = int(tmp16)

		case section == 7:
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

			definition.ModelOffset1 = int(tmp16)
			if definition.ModelOffset1 > 32767 {
				definition.ModelOffset1 -= 0x10000
			}

		case section == 8:
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

			definition.ModelOffset2 = int(tmp16)
			if definition.ModelOffset2 > 32767 {
				definition.ModelOffset2 -= 0x10000
			}

		case section == 10:
			// unused
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 11:
			definition.Stackable = true

		case section == 12:
			err = tmp32.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

			definition.ShopValue = int(tmp32)

		case section == 16:
			definition.Members = true

		case section == 23:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

			err = tmp8.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 24:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 25:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

			err = tmp8.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 26:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section >= 30 && section < 35:
			err = tmpString.Decode(buf, 0)
			if err != nil {
				return nil, err
			}

			i := section - 30
			definition.GroundActions[i] = string(tmpString)
			if definition.GroundActions[i] == "hidden" {
				definition.GroundActions[i] = ""
			}

		case section >= 35 && section < 40:
			err = tmpString.Decode(buf, 0)
			if err != nil {
				return nil, err
			}

			i := section - 35
			definition.Actions[i] = string(tmpString)

		case section == 40:
			err = tmp8.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

			numColorsSwapped := int(tmp8)
			for i := 0; i < numColorsSwapped; i++ {
				err = tmp16.Decode(buf, nilFlags)
				if err != nil {
					return nil, err
				}
				value := uint16(tmp16)

				err = tmp16.Decode(buf, nilFlags)
				if err != nil {
					return nil, err
				}
				key := uint16(tmp16)

				definition.PaletteSwaps[key] = value
			}

		case section == 78:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 79:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 90:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 91:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 92:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 93:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 94:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 95:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 96:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 97:
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

			definition.ParentId = int(tmp16)

		case section == 98:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}
			definition.NotedId = int(tmp16)

		case section >= 100 && section < 110:
			i := section - 100

			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

			definition.StackIds[i] = int(tmp16)

			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

			definition.StackAmounts[i] = int(tmp16)

		case section == 110:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 111:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 112:
			//unknown
			err = tmp16.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 113:
			//unknown
			err = tmp8.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 114:
			//unknown
			err = tmp8.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

		case section == 115:
			err = tmp8.Decode(buf, nilFlags)
			if err != nil {
				return nil, err
			}

			definition.Team = int(tmp8)

		default:
			panic(fmt.Sprintf("unknown section %v", section))
		}
	}
}
