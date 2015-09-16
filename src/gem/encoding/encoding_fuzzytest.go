// +build gofuzz
package encoding

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func Fuzz(data []byte) (ret int) {
	buffer := bytes.NewBuffer(data)
	endian := binary.BigEndian

	var tmp8 uint8

	if err := binary.Read(buffer, endian, &tmp8); err != nil {
		return -1
	}

	flags := IntegerFlag(tmp8)
	if flags&IntReverse != 0 {
		return -1
	}

	var tmp16 uint16
	var tmp32 uint32
	var tmp64 uint64

	if err := binary.Read(buffer, endian, &tmp8); err != nil {
		return -1
	} else if err := binary.Read(buffer, endian, &tmp16); err != nil {
		return -1
	} else if err := binary.Read(buffer, endian, &tmp32); err != nil {
		return -1
	} else if err := binary.Read(buffer, endian, &tmp64); err != nil {
		return -1
	}

	i8 := new(Int8)
	*i8 = Int8(tmp8)
	i16 := new(Int16)
	*i16 = Int16(tmp16)
	i32 := new(Int32)
	*i32 = Int32(tmp32)
	i64 := new(Int64)
	*i64 = Int64(tmp64)
	values := []Codable{i8, i16, i32, i64}

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Flags: %b\n", flags)
			fmt.Printf("Endian: %v\n", flags.endian())
			fmt.Printf("i8: %x\n", i8.Value())
			fmt.Printf("i16: %x\n", i16.Value())
			fmt.Printf("i32: %x\n", i32.Value())
			fmt.Printf("i64: %x\n", i64.Value())
			panic(err)
		}
	}()

	func() {
		defer func() {
			if err := recover(); err != nil {
				if err != ErrUnsupported16 && err != ErrUnsupported64 {
					panic(err)
				}
			}
		}()
		buffer = bytes.NewBuffer([]byte{})
		for _, v := range values {
			err := v.Encode(buffer, flags)
			if err != nil {
				panic(err)
			}
		}

		encoded := buffer.Bytes()

		buffer = bytes.NewBuffer(encoded)

		for _, v := range values {
			cpy := v.Value()
			err := v.Decode(buffer, flags)
			if err != nil {
				panic(err)
			}

			if v.Value() != cpy {
				panic(fmt.Errorf("value mismatch: %x %x", cpy, v.Value()))
			}
		}
	}()
	return 1
}
