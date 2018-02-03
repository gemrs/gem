package rt3

import (
	"fmt"
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
)

const whirlpoolLen = 64

type ReferenceTableFlags int

const (
	RefFlagIdents    ReferenceTableFlags = (1 << 0)
	RefFlagWhirlpool                     = (1 << 1)
	RefFlagSizes                         = (1 << 2)
	RefFlagHash                          = (1 << 3)
)

type ReferenceTable struct {
	Format      int
	Version     int
	Flags       ReferenceTableFlags
	Entries     map[int]*ReferenceEntry
	Identifiers *IdentifierMap
	Capacity    int
}

func (struc *ReferenceTable) Decode(buf io.Reader, flags interface{}) {
	var tmp8 encoding.Uint8
	var tmp32 encoding.Uint32
	var tmpBytes encoding.Bytes

	tmp8.Decode(buf, encoding.IntNilFlag)
	struc.Format = int(tmp8)
	if struc.Format < 5 || struc.Format > 7 {
		panic(fmt.Errorf("unsupported format %v", struc.Format))
	}

	tmpPacked16 := &refTable16{
		Format: struc.Format,
	}

	if struc.Format >= 6 {
		tmp32.Decode(buf, encoding.IntNilFlag)
		struc.Version = int(tmp32)
	}

	tmp8.Decode(buf, encoding.IntNilFlag)
	struc.Flags = ReferenceTableFlags(tmp8)

	tmpPacked16.Decode(buf, nil)
	numFiles := int(tmpPacked16.Value)
	ids := make([]int, numFiles)
	accum := 0
	maxId := -1
	for i, _ := range ids {
		tmpPacked16.Decode(buf, nil)
		delta := int(tmpPacked16.Value)
		accum += delta
		ids[i] = accum
		if ids[i] > maxId {
			maxId = ids[i]
		}
	}
	maxId++
	struc.Capacity = maxId

	struc.Entries = make(map[int]*ReferenceEntry)
	for i, id := range ids {
		struc.Entries[id] = &ReferenceEntry{
			Id:    id,
			Index: i,
		}
	}

	identifiers := make([]int, maxId)
	if struc.Flags&RefFlagIdents != 0 {
		for _, id := range ids {
			tmp32.Decode(buf, encoding.IntNilFlag)
			identifier := int(tmp32)
			identifiers[id] = identifier
			struc.Entries[id].Identifier = identifier
		}
	}

	struc.Identifiers = NewIdentifierMap(identifiers)

	for _, id := range ids {
		tmp32.Decode(buf, encoding.IntNilFlag)
		struc.Entries[id].Crc = int(tmp32)
	}

	if struc.Flags&RefFlagHash != 0 {
		for _, id := range ids {
			tmp32.Decode(buf, encoding.IntNilFlag)
			struc.Entries[id].Hash = int(tmp32)
		}
	}

	if struc.Flags&RefFlagWhirlpool != 0 {
		for _, id := range ids {
			tmpBytes.Decode(buf, whirlpoolLen)
			struc.Entries[id].Whirlpool = make([]byte, whirlpoolLen)
			copy(struc.Entries[id].Whirlpool, []byte(tmpBytes))
		}
	}

	if struc.Flags&RefFlagSizes != 0 {
		for _, id := range ids {
			tmp32.Decode(buf, encoding.IntNilFlag)
			struc.Entries[id].Compressed = int(tmp32)
			tmp32.Decode(buf, encoding.IntNilFlag)
			struc.Entries[id].Uncompressed = int(tmp32)
		}
	}

	for _, id := range ids {
		tmp32.Decode(buf, encoding.IntNilFlag)
		struc.Entries[id].Version = int(tmp32)
	}

	members := make([][]int, maxId)
	for _, id := range ids {
		tmpPacked16.Decode(buf, nil)
		numSubEntries := int(tmpPacked16.Value)
		members[id] = make([]int, numSubEntries)
	}

	for _, id := range ids {
		accum := 0
		maxId := -1

		for i, _ := range members[id] {
			tmpPacked16.Decode(buf, nil)
			delta := int(tmpPacked16.Value)
			accum += delta
			members[id][i] = accum
			if members[id][i] > maxId {
				maxId = members[id][i]
			}
		}
		maxId++

		struc.Entries[id].Children = make(map[int]*ReferenceChildEntry)
		for index, child := range members[id] {
			struc.Entries[id].Children[child] = &ReferenceChildEntry{
				Id:    child,
				Index: index,
			}
		}
	}

	if struc.Flags&RefFlagIdents != 0 {
		for _, id := range ids {
			identifiers := make([]int, len(members[id]))
			for _, child := range members[id] {
				tmp32.Decode(buf, encoding.IntNilFlag)
				identifier := int(tmp32)
				identifiers[child] = identifier
				struc.Entries[id].Children[child].Identifier = identifier
			}
			struc.Entries[id].Identifiers = NewIdentifierMap(identifiers)
		}
	}
}

func (table *ReferenceTable) UncompressedSize() int {
	sum := 0
	for _, e := range table.Entries {
		sum += e.Uncompressed
	}
	return sum
}

type ReferenceEntry struct {
	Id           int
	Index        int
	Identifier   int
	Crc          int
	Hash         int
	Whirlpool    []byte
	Compressed   int
	Uncompressed int
	Version      int
	Children     map[int]*ReferenceChildEntry
	Identifiers  *IdentifierMap
}

type ReferenceChildEntry struct {
	Id         int
	Index      int
	Identifier int
}

type IdentifierMap struct {
	table []int
}

func NewIdentifierMap(src []int) *IdentifierMap {
	var idents IdentifierMap

	length := len(src)
	halfLength := length >> 1

	size := 1
	mask := 1
	for i := 1; i <= length+halfLength; i <<= 1 {
		mask = i
		size = i << 1
	}

	mask <<= 1
	size <<= 1

	idents.table = make([]int, size)

	for i, _ := range idents.table {
		idents.table[i] = -1
	}

	for id, identifier := range src {
		i := 0
		for i := identifier & (mask - 1); idents.table[i+i+1] != -1; i = (i + 1) & (mask - 1) {
		}

		idents.table[i+i] = identifier
		idents.table[i+i+1] = id
	}

	return &idents
}

func (m *IdentifierMap) Get(ident int) int {
	mask := (len(m.table) >> 1) - 1
	i := ident & mask

	for {
		id := m.table[i+i+1]
		if id == -1 {
			return -1
		}

		if m.table[i+i] == ident {
			return id
		}

		i = (i + 1) & mask
	}
}

type refTable16 struct {
	Value  int
	Format int
}

func (i *refTable16) Decode(buf io.Reader, flags interface{}) {
	flags = encoding.IntNilFlag
	if i.Format >= 7 {
		flags = encoding.IntPacked
	}
	var value encoding.Uint16
	value.Decode(buf, flags)
	i.Value = int(value)
}
