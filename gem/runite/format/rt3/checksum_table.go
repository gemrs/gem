package rt3

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
)

type ChecksumTable struct {
	Entries []ChecksumEntry
}

func (c *ChecksumTable) AddEntry(e ChecksumEntry) {
	c.Entries = append(c.Entries, e)
}

func (c *ChecksumTable) Encode(buf io.Writer, flags_ interface{}) {
	whirlpool := flags_.(bool)

	if whirlpool {
		panic("whirlpool checksum table not implemented")
	}

	for _, e := range c.Entries {
		e.Encode(buf, nil)
	}
}

type ChecksumEntry struct {
	Crc       uint32
	Version   int
	FileCount int
	Size      int
	Whirlpool []byte
}

func (entry *ChecksumEntry) Encode(buf io.Writer, flags_ interface{}) {
	encoding.Uint32(entry.Crc).Encode(buf, encoding.IntNilFlag)
	encoding.Uint32(entry.Version).Encode(buf, encoding.IntNilFlag)
}
