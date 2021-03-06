package protocol_os_162

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/protocol"
)

// +gen define_inbound:"Pkt45,SzVar8"
type InboundPlayerWalk protocol.InboundPlayerWalk

func (struc *InboundPlayerWalk) Decode(buf io.Reader, flags interface{}) {
	decodeWalk(buf, flags, (*protocol.InboundPlayerWalk)(struc), false)
}

// +gen define_inbound:"Pkt84,SzVar8,InboundPlayerWalk"
type InboundPlayerWalkMap protocol.InboundPlayerWalk

func (struc *InboundPlayerWalkMap) Decode(buf io.Reader, flags interface{}) {
	decodeWalk(buf, flags, (*protocol.InboundPlayerWalk)(struc), true)
}

func decodeWalk(r io.Reader, flags interface{}, struc *protocol.InboundPlayerWalk, mapClick bool) {
	buf := encoding.WrapReader(r)

	struc.Y = buf.GetU16(encoding.IntLittleEndian)
	struc.X = buf.GetU16(encoding.IntOffset128)
	runMode := buf.GetU8(encoding.IntNegate)
	// FIXME runMode can also be 2 sometimes?
	struc.Running = runMode == 1

	if mapClick {
		// ignore the extra 13 bytes for now
		buf.Read(make([]byte, 13))
	}
}
