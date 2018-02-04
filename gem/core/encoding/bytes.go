package encoding

import (
	"io"
)

type Bytes []byte

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (i Bytes) Encode(buf io.Writer, flags interface{}) {
	length, ok := flags.(int)
	if !ok {
		length = -1
	}

	if length == -1 {
		_, err := buf.Write(i)
		if err != nil {
			panic(err)
		}
	} else {
		length = min(length, len(i))
		_, err := buf.Write(i[:length])
		if err != nil {
			panic(err)
		}
	}
}

func (i *Bytes) Decode(buf io.Reader, flags interface{}) {
	length, ok := flags.(int)
	if !ok {
		length = -1
	}

	if length == -1 {
		// All remaining data
		result := make(Bytes, 0)
		buffer := make(Bytes, 128)
		read := 0
		var err error
		for err == nil {
			read, err = buf.Read(buffer)
			result = append(result, buffer[:read]...)
		}

		*i = result

		if err != io.EOF {
			panic(err)
		}
	} else {
		*i = make(Bytes, length)
		_, err := buf.Read((*i)[:length])
		if err != nil {
			panic(err)
		}
	}
}

func (i *Bytes) Value() interface{} {
	return *i
}
