package protocol_317

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
)

type OutboundGameHandshake struct {
	ignored         [8]encoding.Uint8
	loginRequest    encoding.Uint8
	ServerISAACSeed [2]encoding.Uint32
}

func (struc *OutboundGameHandshake) Encode(buf io.Writer, flags interface{}) {
	for i := 0; i < 8; i++ {
		struc.ignored[i].Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	}

	struc.loginRequest.Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	for i := 0; i < 2; i++ {
		struc.ServerISAACSeed[i].Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	}
}

type OutboundUpdateHandshake struct {
	ignored [8]encoding.Uint8
}

func (struc *OutboundUpdateHandshake) Encode(buf io.Writer, flags interface{}) {
	for i := 0; i < 8; i++ {
		struc.ignored[i].Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	}
}

type InboundServiceSelect struct {
	Service encoding.Uint8
}

func (struc *InboundServiceSelect) Decode(buf io.Reader, flags interface{}) {
	struc.Service.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
}

type InboundGameHandshake struct {
	NameHash encoding.Uint8
}

func (struc *InboundGameHandshake) Decode(buf io.Reader, flags interface{}) {
	struc.NameHash.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
}
