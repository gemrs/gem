package encoding

import (
	"fmt"
	"sort"
)

var ErrUnsupported16 = fmt.Errorf("PDPEndian: unsupported type Uint16")
var ErrUnsupported64 = fmt.Errorf("PDPEndian: unsupported type Uint64")

var PDPEndian = pdpEndian{false}
var RPDPEndian = pdpEndian{true}

type ByteSlice []byte

func (p ByteSlice) Len() int           { return len(p) }
func (p ByteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p ByteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type pdpEndian struct {
	reverse bool
}

func (e pdpEndian) Uint16(b []byte) uint16 { panic(ErrUnsupported16) }

func (e pdpEndian) PutUint16(b []byte, v uint16) { panic(ErrUnsupported16) }

func (e pdpEndian) Uint32(orig []byte) uint32 {
	b := make([]byte, 4)
	copy(b, orig[0:4])

	if e.reverse {
		sort.Sort(sort.Reverse(ByteSlice(b[0:4])))
	}

	return uint32(b[0])<<16 | uint32(b[1])<<24 | uint32(b[2]) | uint32(b[3])<<8
}

func (e pdpEndian) PutUint32(b []byte, v uint32) {
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[0] = byte(v)
	b[1] = byte(v >> 8)

	if e.reverse {
		sort.Sort(sort.Reverse(ByteSlice(b[0:4])))
	}
}

func (e pdpEndian) Uint64(b []byte) uint64 { panic(ErrUnsupported64) }

func (e pdpEndian) PutUint64(b []byte, v uint64) { panic(ErrUnsupported64) }

func (e pdpEndian) String() string {
	switch {
	case !e.reverse:
		return "PDPEndian"
	case e.reverse:
		return "RPDPEndian"
	}
	panic("never reached")
}

func (e pdpEndian) GoString() string {
	switch {
	case !e.reverse:
		return "encoding.PDPEndian"
	case e.reverse:
		return "encoding.RPDPEndian"
	}
	panic("never reached")
}
