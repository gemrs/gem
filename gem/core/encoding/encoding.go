package encoding

import (
	"bytes"
	"io"
)

type Codable interface {
	Encodable
	Decodable
}

type Decodable interface {
	Decode(buf io.Reader, flags_ interface{})
}

type Encodable interface {
	Encode(buf io.Writer, flags_ interface{})
}

type Encoded struct {
	bytes.Buffer
}

func (e *Encoded) Read(p []byte) (n int, err error) {
	return e.Buffer.Read(p)
}

var NilFlags int = 0
