package protocol_os_161

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/protocol"
)

type InboundRsaLoginBlock struct {
	Magic     encoding.Uint8
	AuthType  encoding.Uint8
	ISAACSeed [4]encoding.Uint32
	AuthData  [8]encoding.Uint8
	Password  encoding.String
}

func (struc InboundRsaLoginBlock) Encode(buf io.Writer, flags interface{}) {
	panic("not implemented")
}

func (struc *InboundRsaLoginBlock) Decode(buf io.Reader, flags interface{}) {
	struc.Magic.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	struc.AuthType.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	for i := 0; i < 4; i++ {
		struc.ISAACSeed[i].Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	}

	for i := 0; i < 8; i++ {
		struc.AuthData[i].Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	}

	struc.Password.Decode(buf, 0)
}

type InboundXteaLoginBlock struct {
	Username encoding.String
}

func (struc InboundXteaLoginBlock) Encode(buf io.Writer, flags interface{}) {
	panic("not implemented")
}

func (struc *InboundXteaLoginBlock) Decode(buf io.Reader, flags interface{}) {
	struc.Username.Decode(buf, 0)
}

// +gen define_outbound
type OutboundLoginResponse protocol.OutboundLoginResponse

func (struc OutboundLoginResponse) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint8(struc.Response).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	if struc.Response != protocol.AuthOkay {
		return
	}

	encoding.Uint8(0).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	encoding.Uint32(0).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	encoding.Uint8(struc.Rights).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	encoding.Uint8(1).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	encoding.Uint16(struc.Index).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	encoding.Uint8(1).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
}

type InboundLoginBlock struct {
	LoginType       encoding.Uint8
	LoginLen        encoding.Uint16
	Revision        encoding.Uint32
	SecureBlockSize encoding.Uint16
}

func (struc *InboundLoginBlock) Decode(buf io.Reader, flags interface{}) {
	struc.LoginType.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	struc.LoginLen.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	struc.Revision.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	struc.SecureBlockSize.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
}
