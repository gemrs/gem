package protocol_os_162

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol"
)

type OutboundGameHandshake struct {
	UID             encoding.Uint8
	ServerISAACSeed [2]encoding.Uint32
}

func (struc *OutboundGameHandshake) Encode(buf io.Writer, flags interface{}) {
	struc.UID.Encode(buf, encoding.IntNilFlag)

	for i := 0; i < 2; i++ {
		struc.ServerISAACSeed[i].Encode(buf, encoding.IntNilFlag)
	}
}

type InboundServiceSelect struct {
	Service encoding.Uint8
}

func (struc *InboundServiceSelect) Decode(buf io.Reader, flags interface{}) {
	struc.Service.Decode(buf, encoding.IntNilFlag)
}

type InboundUpdateHandshake struct {
	Revision encoding.Uint32
}

func (struc *InboundUpdateHandshake) Decode(buf io.Reader, flags interface{}) {
	struc.Revision.Decode(buf, encoding.IntNilFlag)
}

type OutboundUpdateHandshake struct {
	Response int
}

func (struc *OutboundUpdateHandshake) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint8(struc.Response).Encode(buf, encoding.IntNilFlag)
}

func (p protocolImpl) Handshake(conn *server.Connection) int {
	var serviceSelect InboundServiceSelect
	serviceSelect.Decode(conn.NetConn(), nil)

	service := int(serviceSelect.Service)

	response := 0
	if service == UpdateServiceId {
		var handshake InboundUpdateHandshake
		handshake.Decode(conn.NetConn(), nil)
		if handshake.Revision != Revision {
			response = int(protocol.AuthUpdates)
		}

		var handshakeResponse OutboundUpdateHandshake
		handshakeResponse.Response = response
		handshakeResponse.Encode(conn.WriteBuffer, nil)
		conn.FlushWriteBuffer()
	}

	return service
}
