package encoding

import (
	"bytes"
	"fmt"
	"io"
)

const (
	stringDelim byte = 0xA
)

// A string terminated by 0xA
type JString string

func (str JString) Encode(buf io.Writer, flags_ interface{}) {
	length := flags_.(int)

	data := []byte(str)
	if length == 0 {
		// variable length - append stringDelim
		data = append(data, stringDelim)
	} else {
		// fixed length - pad for n bytes
		if len(data) != length {
			if len(data) > length {
				panic(fmt.Errorf("string out of bounds for fixed length of %v bytes", length))
			}
			padding := make([]byte, length-len(data))
			data = append(data, padding...)
		}
	}

	_, err := buf.Write(data)
	if err != nil {
		panic(err)
	}
}

func (str *JString) Decode(buf io.Reader, flags_ interface{}) {
	length := flags_.(int)
	b := make([]byte, length)
	if length == 0 {
		// variable length array - read until stringDelim
		chr := make([]byte, 1)
		for {
			_, err := buf.Read(chr)
			if err != nil {
				panic(err)
			}

			if chr[0] == stringDelim {
				break
			}
			b = append(b, chr[0])
		}
	} else {
		// Fixed length array - just fill our buffer with n bytes
		buf.Read(b)
	}

	b = bytes.Trim(b, "\x00\x0A")

	*str = JString(b)
}
