package encoding

import (
	"io"
)

type Codable interface {
	Encodable
	Decodable
}

type Decodable interface {
	Decode(buf io.Reader, flags_ interface{}) error
}

type Encodable interface {
	Encode(buf io.Writer, flags_ interface{}) error
}

var NilFlags int = 0
