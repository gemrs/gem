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

func (i *Bytes) Encode(buf io.Writer, flags interface{}) error {
	if flags == nil {
		_, err := buf.Write(*i)
		return err
	} else {
		length := flags.(int)
		length = min(length, len(*i))
		_, err := buf.Write((*i)[:length])
		return err
	}
}

func (i *Bytes) Decode(buf io.Reader, flags interface{}) error {
	if flags == nil {
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
			return err
		}

		return nil
	} else {
		length := flags.(int)
		*i = make(Bytes, length)
		_, err := buf.Read((*i)[:length])
		return err
	}
}

func (i *Bytes) Value() interface{} {
	return *i
}
