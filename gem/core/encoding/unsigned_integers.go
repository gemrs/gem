package encoding

import (
	"io"
)

type Uint8 uint8

func (i Uint8) Encode(buf io.Writer, flags_ interface{}) {
	var flags IntegerFlag
	if flags_ != nil {
		flags = flags_.(IntegerFlag)
	}

	value := flags.apply(uint64(i))
	_, err := buf.Write([]byte{byte(value)})
	if err != nil {
		panic(err)
	}
}

func (i *Uint8) Decode(buf io.Reader, flags_ interface{}) {
	var flags IntegerFlag
	if flags_ != nil {
		flags = flags_.(IntegerFlag)
	}

	b := make([]byte, 1)
	_, err := buf.Read(b)
	if err != nil {
		panic(err)
	}

	*i = Uint8(flags.reverse().apply(uint64(b[0])))
}

func (i *Uint8) Value() interface{} {
	return *i
}

type Uint16 uint16

func (i Uint16) Encode(buf io.Writer, flags_ interface{}) {
	var flags IntegerFlag
	if flags_ != nil {
		flags = flags_.(IntegerFlag)
	}

	value := flags.apply(uint64(i))
	endian := flags.endian()

	if flags&IntPacked != 0 {
		if value < 0x80 {
			Uint8(value).Encode(buf, IntNilFlag)
			return
		}
		value += 0x8000
	}

	data := make([]byte, 2)
	endian.PutUint16(data, uint16(value))
	_, err := buf.Write(data)
	if err != nil {
		panic(err)
	}
}

func (i *Uint16) Decode(buf io.Reader, flags_ interface{}) {
	var flags IntegerFlag
	if flags_ != nil {
		flags = flags_.(IntegerFlag)
	}

	endian := flags.endian()

	data := make([]byte, 2)

	if flags&IntPacked != 0 {
		var tmp8 Uint8
		tmp8.Decode(buf, IntNilFlag)
		if byte(tmp8) < 0x80 {
			*i = Uint16(tmp8)
			return
		}
		data[0] = byte(tmp8)

		tmp8.Decode(buf, IntNilFlag)
		data[1] = byte(tmp8)
	} else {
		_, err := buf.Read(data)
		if err != nil {
			panic(err)
		}
	}

	value64 := uint64(endian.Uint16(data))
	*i = Uint16(flags.reverse().apply(value64))
}

func (i *Uint16) Value() interface{} {
	return *i
}

type Uint24 uint32

func (i Uint24) Encode(buf io.Writer, flags_ interface{}) {
	var flags IntegerFlag
	if flags_ != nil {
		flags = flags_.(IntegerFlag)
	}

	value := flags.apply(uint64(i))

	data := []byte{byte(value >> 16), byte(value >> 8), byte(value)}
	_, err := buf.Write(data)
	if err != nil {
		panic(err)
	}
}

func (i *Uint24) Decode(buf io.Reader, flags_ interface{}) {
	var flags IntegerFlag
	if flags_ != nil {
		flags = flags_.(IntegerFlag)
	}

	data := make([]byte, 3)
	_, err := buf.Read(data)
	if err != nil {
		panic(err)
	}

	value64 := uint64(data[0])<<16 | uint64(data[1])<<8 | uint64(data[2])
	*i = Uint24(flags.reverse().apply(value64))
}

func (i *Uint24) Value() interface{} {
	return *i
}

type Uint32 uint32

func (i Uint32) Encode(buf io.Writer, flags_ interface{}) {
	var flags IntegerFlag
	if flags_ != nil {
		flags = flags_.(IntegerFlag)
	}

	value := flags.apply(uint64(i))
	endian := flags.endian()

	data := make([]byte, 4)
	endian.PutUint32(data, uint32(value))
	_, err := buf.Write(data)
	if err != nil {
		panic(err)
	}
}

func (i *Uint32) Decode(buf io.Reader, flags_ interface{}) {
	var flags IntegerFlag
	if flags_ != nil {
		flags = flags_.(IntegerFlag)
	}

	endian := flags.endian()

	data := make([]byte, 4)
	_, err := buf.Read(data)
	if err != nil {
		panic(err)
	}

	value64 := uint64(endian.Uint32(data))
	*i = Uint32(flags.reverse().apply(value64))
}

func (i *Uint32) Value() interface{} {
	return *i
}

type Uint64 uint64

func (i Uint64) Encode(buf io.Writer, flags_ interface{}) {
	var flags IntegerFlag
	if flags_ != nil {
		flags = flags_.(IntegerFlag)
	}

	value := flags.apply(uint64(i))
	endian := flags.endian()

	data := make([]byte, 8)
	endian.PutUint64(data, uint64(value))
	_, err := buf.Write(data)
	if err != nil {
		panic(err)
	}
}

func (i *Uint64) Decode(buf io.Reader, flags_ interface{}) {
	var flags IntegerFlag
	if flags_ != nil {
		flags = flags_.(IntegerFlag)
	}

	endian := flags.endian()

	data := make([]byte, 8)
	_, err := buf.Read(data)
	if err != nil {
		panic(err)
	}

	value64 := uint64(endian.Uint64(data))
	*i = Uint64(flags.reverse().apply(value64))
}

func (i *Uint64) Value() interface{} {
	return *i
}
