package protocol_317

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/protocol"
)

type InboundSecureLoginBlock struct {
	Magic     encoding.Uint8
	ISAACSeed [4]encoding.Uint32
	ClientUID encoding.Uint32
	Username  encoding.JString
	Password  encoding.JString
}

func (struc InboundSecureLoginBlock) Encode(buf io.Writer, flags interface{}) {
	panic("not implemented")
}

func (struc *InboundSecureLoginBlock) Decode(buf io.Reader, flags interface{}) {
	struc.Magic.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	for i := 0; i < 4; i++ {
		struc.ISAACSeed[i].Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	}

	struc.ClientUID.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	struc.Username.Decode(buf, 0)
	struc.Password.Decode(buf, 0)
}

// +gen define_outbound
type OutboundLoginResponse protocol.OutboundLoginResponse

func (struc OutboundLoginResponse) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint8(struc.Response).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))

	if struc.Response == protocol.AuthOkay {
		encoding.Uint8(struc.Rights).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
		flagged := 0
		if struc.Flagged {
			flagged = 1
		}
		encoding.Uint8(flagged).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	}
}

type InboundLoginBlock struct {
	LoginType       encoding.Uint8
	LoginLen        encoding.Uint8
	Magic           encoding.Uint8
	Revision        encoding.Uint16
	MemType         encoding.Uint8
	ArchiveCRCs     [9]encoding.Uint32
	SecureBlockSize encoding.Uint8
}

func (struc *InboundLoginBlock) Decode(buf io.Reader, flags interface{}) {
	struc.LoginType.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	struc.LoginLen.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	struc.Magic.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	struc.Revision.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	struc.MemType.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	for i := 0; i < 9; i++ {
		struc.ArchiveCRCs[i].Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	}
	struc.SecureBlockSize.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
}
