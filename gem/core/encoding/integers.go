package encoding

import (
	"encoding/binary"
	"io"
	"unsafe"
)

type IntegerFlag int

const (
	IntNilFlag IntegerFlag = (1 << iota)
	IntNegate
	IntInverse128
	IntOffset128
	IntLittleEndian
	IntPDPEndian
	IntRPDPEndian
	IntReverse
)

func (f IntegerFlag) reverse() IntegerFlag {
	return f | IntReverse
}

func (f IntegerFlag) endian() binary.ByteOrder {
	switch {
	case f&IntLittleEndian != 0:
		return binary.LittleEndian
	case f&IntPDPEndian != 0:
		return PDPEndian
	case f&IntRPDPEndian != 0:
		return RPDPEndian
	default:
		return binary.BigEndian
	}
	panic("never reached")
}

func (f IntegerFlag) apply(value uint64) uint64 {
	data := (*[8]byte)(unsafe.Pointer(&value))[:8]
	if f&IntNegate != 0 {
		data[0] = -data[0]
	}
	if f&IntInverse128 != 0 {
		data[0] = 128 - data[0]
	}
	if f&IntOffset128 != 0 {
		if f&IntReverse != 0 {
			data[0] = data[0] - 128
		} else {
			data[0] = data[0] + 128
		}
	}
	return value
}

type Int8 int8

func (i Int8) Encode(buf io.Writer, flags_ interface{}) {
	unsigned := Uint8(uint8(i))
	unsigned.Encode(buf, flags_)
}

func (i *Int8) Decode(buf io.Reader, flags_ interface{}) {
	var unsigned Uint8
	unsigned.Decode(buf, flags_)
	*i = Int8(unsigned)
}

func (i *Int8) Value() interface{} {
	return *i
}

type Int16 int16

func (i Int16) Encode(buf io.Writer, flags_ interface{}) {
	unsigned := Uint16(uint16(i))
	unsigned.Encode(buf, flags_)
}

func (i *Int16) Decode(buf io.Reader, flags_ interface{}) {
	var unsigned Uint16
	unsigned.Decode(buf, flags_)
	*i = Int16(unsigned)
}

func (i *Int16) Value() interface{} {
	return *i
}

type Int24 uint32

func (i Int24) Encode(buf io.Writer, flags_ interface{}) {
	unsigned := Uint24(uint32(i))
	unsigned.Encode(buf, flags_)
}

func (i *Int24) Decode(buf io.Reader, flags_ interface{}) {
	var unsigned Uint24
	unsigned.Decode(buf, flags_)
	*i = Int24(unsigned)
}

func (i *Int24) Value() interface{} {
	return *i
}

type Int32 uint32

func (i Int32) Encode(buf io.Writer, flags_ interface{}) {
	unsigned := Uint32(uint32(i))
	unsigned.Encode(buf, flags_)
}

func (i *Int32) Decode(buf io.Reader, flags_ interface{}) {
	var unsigned Uint32
	unsigned.Decode(buf, flags_)
	*i = Int32(unsigned)
}

func (i *Int32) Value() interface{} {
	return *i
}

type Int64 uint64

func (i Int64) Encode(buf io.Writer, flags_ interface{}) {
	unsigned := Uint64(uint64(i))
	unsigned.Encode(buf, flags_)
}

func (i *Int64) Decode(buf io.Reader, flags_ interface{}) {
	var unsigned Uint64
	unsigned.Decode(buf, flags_)
	*i = Int64(unsigned)
}

func (i *Int64) Value() interface{} {
	return *i
}
