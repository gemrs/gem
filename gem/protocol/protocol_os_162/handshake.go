package protocol_os_162

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/protocol"
)

// +gen define_outbound
type OutboundLoginResponse protocol.OutboundLoginResponse

func (struc OutboundLoginResponse) Encode(w io.Writer, flags interface{}) {
	buf := encoding.WrapWriter(w)

	buf.PutU8(int(struc.Response))

	if struc.Response != protocol.AuthOkay {
		return
	}

	buf.PutU8(0)
	buf.PutU32(0)
	buf.PutU8(struc.Rights)
	buf.PutU8(1)
	buf.PutU16(struc.Index)
	buf.PutU8(1)
}
