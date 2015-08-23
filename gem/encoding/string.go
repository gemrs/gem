package encoding

import (
	"bytes"
)

const (
	stringDelim byte = 0xA
)

// A string terminated by 0xA
type JString string

func (str *JString) Encode(buf *bytes.Buffer, flags_ interface{}) error {
	data := []byte(*str)
	data = append(data, stringDelim)

	_, err := buf.Write(data)
	return err
}

func (str *JString) Decode(buf *bytes.Buffer, flags_ interface{}) error {
	line, err := buf.ReadString(stringDelim)
	*str = JString(line)
	return err
}
