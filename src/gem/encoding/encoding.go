package encoding

import (
	"io"
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
	Encode(buf io.Writer, flags_ interface{}) error
}

var NilFlags int = 0
