package encoding

import (
	"io"
	"bytes"
)

type Codable interface {
	Encoder
	Decoder
	Value() interface{}
}

type Decoder interface {
	Decode(buf io.Reader, flags_ interface{}) error
}

type Encoder interface {
	Encode(buf *bytes.Buffer, flags_ interface{}) error
}

var NilFlags int = 0
