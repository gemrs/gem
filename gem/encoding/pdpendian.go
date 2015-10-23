package encoding

import (
	"fmt"
	"unsafe"
)

var ErrUnsupported16 = fmt.Errorf("PDPEndian: unsupported type Uint16")
var ErrUnsupported64 = fmt.Errorf("PDPEndian: unsupported type Uint64")

var PDPEndian = pdpEndian{false}
var RPDPEndian = pdpEndian{true}

type pdpEndian struct {
	reverse bool
}

func (e pdpEndian) Uint16(b []byte) uint16 { panic(ErrUnsupported16) }

func (e pdpEndian) PutUint16(b []byte, v uint16) { panic(ErrUnsupported16) }

func (e pdpEndian) Uint32(orig []byte) uint32 {
	out := make([]byte, 4)
	in := make([]byte, 4)
	// Make a copy of orig so that we can reverse it
	copy(in, orig[0:4])

	if e.reverse {
		in[0], in[1], in[2], in[3] = in[3], in[2], in[1], in[0]
	}

	out[0], out[1], out[2], out[3] = in[1], in[0], in[3], in[2]

	return uint32(out[3])<<24 | uint32(out[2])<<16 | uint32(out[1])<<8 | uint32(out[0])
}

func (e pdpEndian) PutUint32(out []byte, v uint32) {
	in := (*[4]byte)(unsafe.Pointer(&v))[:4]

	out[0], out[1], out[2], out[3] = in[1], in[0], in[3], in[2]

	if e.reverse {
		out[0], out[1], out[2], out[3] = out[3], out[2], out[1], out[0]
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
