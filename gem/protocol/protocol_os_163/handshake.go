package protocol_os_163

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol"
)

type OutboundGameHandshake struct {
	UID             int
	ServerISAACSeed [2]int
}

func (struc *OutboundGameHandshake) Encode(w io.Writer, flags interface{}) {
	buf := encoding.WrapWriter(w)
	buf.PutU8(struc.UID)

	for i := 0; i < 2; i++ {
		buf.PutU32(struc.ServerISAACSeed[i])
	}
}

type InboundServiceSelect struct {
	Service int
}

func (struc *InboundServiceSelect) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Service = buf.GetU8()
}

type InboundUpdateHandshake struct {
	Revision int
}

func (struc *InboundUpdateHandshake) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Revision = buf.GetU32()
}

type OutboundUpdateHandshake struct {
	Response int
}

func (struc *OutboundUpdateHandshake) Encode(w io.Writer, flags interface{}) {
	buf := encoding.WrapWriter(w)
	buf.PutU8(struc.Response)
}

func (p protocolImpl) Handshake(conn *server.Connection) int {
	var serviceSelect InboundServiceSelect
	serviceSelect.Decode(conn.NetConn(), nil)

	service := serviceSelect.Service

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
