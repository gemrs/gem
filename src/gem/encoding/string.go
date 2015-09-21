package encoding

import (
	"bytes"
	"io"
	"fmt"
)

const (
	stringDelim byte = 0xA
)

// A string terminated by 0xA
type JString string

func (str *JString) Encode(buf *bytes.Buffer, flags_ interface{}) error {
	length := flags_.(int)

	data := []byte(*str)
	if length == 0 {
		// variable length - append stringDelim
		data = append(data, stringDelim)
	} else {
		// fixed length - pad for n bytes
		if len(data) != length {
			if len(data) > length {
				return fmt.Errorf("string out of bounds for fixed length of %v bytes", length)
			}
			padding := make([]byte, length - len(data))
			data = append(data, padding...)
		}
	}

	_, err := buf.Write(data)
	return err
}

func (str *JString) Decode(buf io.Reader, flags_ interface{}) error {
	length := flags_.(int)
	b := make([]byte, length)
	i := 0
	if length == 0 {
		// variable length array - read until stringDelim
		for b[i:i+1][0] != stringDelim {
			_, err := buf.Read(b[i:i+1])
			if err != nil {
				return err
			}
		}
	} else {
		// Fixed length array - just fill our buffer with n bytes
		buf.Read(b)
	}

	b = bytes.Trim(b, "\x00\x0A")

	*str = JString(b)
	return nil
}
