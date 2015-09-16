package encoding

import (
	"bytes"
)

type Codable interface {
	Encoder
	Decoder
	Value() interface{}
}

type Decoder interface {
	Decode(buf *bytes.Buffer, flags_ interface{}) error
}

type Encoder interface {
	Encode(buf *bytes.Buffer, flags_ interface{}) error
}

var NilFlags int = 0
