package encoding

import (
	"bytes"
	"io"
)

type Bytes []byte

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (i *Bytes) Encode(buf *bytes.Buffer, flags interface{}) error {
	length := flags.(int)
	length = min(length, len(*i))
	_, err := buf.Write((*i)[:length])
	return err
}

func (i *Bytes) Decode(buf io.Reader, flags interface{}) error {
	length := flags.(int)
	*i = make(Bytes, length)
	_, err := buf.Read((*i)[:length])
	return err
}

func (i *Bytes) Value() interface{} {
	return *i
}
