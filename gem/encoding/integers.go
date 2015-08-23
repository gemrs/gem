package encoding

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

type IntegerFlag int

const (
	IntNilFlag IntegerFlag = (1 << iota)
	IntNegative
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
	if f&IntNegative != 0 {
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

type Int8 uint8

func (i *Int8) Encode(buf *bytes.Buffer, flags_ interface{}) error {
	flags := flags_.(IntegerFlag)
	value := flags.apply(uint64(*i))
	return buf.WriteByte(byte(value))
}

func (i *Int8) Decode(buf *bytes.Buffer, flags_ interface{}) error {
	flags := flags_.(IntegerFlag)
	value, err := buf.ReadByte()
	if err != nil {
		return err
	}

	*i = Int8(flags.reverse().apply(uint64(value)))
	return nil
}

func (i *Int8) Value() interface{} {
	return *i
}

type Int16 uint16

func (i *Int16) Encode(buf *bytes.Buffer, flags_ interface{}) error {
	flags := flags_.(IntegerFlag)
	value := flags.apply(uint64(*i))
	endian := flags.endian()

	data := make([]byte, 2)
	endian.PutUint16(data, uint16(value))
	_, err := buf.Write(data)
	return err
}

func (i *Int16) Decode(buf *bytes.Buffer, flags_ interface{}) error {
	flags := flags_.(IntegerFlag)
	endian := flags.endian()

	data := make([]byte, 2)
	_, err := buf.Read(data)
	if err != nil {
		return err
	}

	value64 := uint64(endian.Uint16(data))
	*i = Int16(flags.reverse().apply(value64))
	return nil
}

func (i *Int16) Value() interface{} {
	return *i
}

type Int32 uint32

func (i *Int32) Encode(buf *bytes.Buffer, flags_ interface{}) error {
	flags := flags_.(IntegerFlag)
	value := flags.apply(uint64(*i))
	endian := flags.endian()

	data := make([]byte, 4)
	endian.PutUint32(data, uint32(value))
	_, err := buf.Write(data)
	return err
}

func (i *Int32) Decode(buf *bytes.Buffer, flags_ interface{}) error {
	flags := flags_.(IntegerFlag)
	endian := flags.endian()

	data := make([]byte, 4)
	_, err := buf.Read(data)
	if err != nil {
		return err
	}

	value64 := uint64(endian.Uint32(data))
	*i = Int32(flags.reverse().apply(value64))
	return nil
}

func (i *Int32) Value() interface{} {
	return *i
}

type Int64 uint64

func (i *Int64) Encode(buf *bytes.Buffer, flags_ interface{}) error {
	flags := flags_.(IntegerFlag)
	value := flags.apply(uint64(*i))
	endian := flags.endian()

	data := make([]byte, 8)
	endian.PutUint64(data, uint64(value))
	_, err := buf.Write(data)
	return err
}

func (i *Int64) Decode(buf *bytes.Buffer, flags_ interface{}) error {
	flags := flags_.(IntegerFlag)
	endian := flags.endian()

	data := make([]byte, 8)
	_, err := buf.Read(data)
	if err != nil {
		return err
	}

	value64 := uint64(endian.Uint64(data))
	*i = Int64(flags.reverse().apply(value64))
	return nil
}

func (i *Int64) Value() interface{} {
	return *i
}
