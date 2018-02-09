package protocol_os_163

import (
	"io"

	"github.com/gemrs/gem/gem/protocol"
)

// +gen define_inbound:"Pkt63,SzVar8"
type InboundMouseMovement protocol.InboundMouseMovement

func (struc *InboundMouseMovement) Decode(r io.Reader, flags interface{}) {}

// +gen define_inbound:"Pkt3,SzFixed"
type InboundMouseClick protocol.InboundMouseClick

func (struc *InboundMouseClick) Decode(r io.Reader, flags interface{}) {}

// +gen define_inbound:"Pkt89,SzFixed"
type InboundPing protocol.InboundPing

func (struc *InboundPing) Decode(r io.Reader, flags interface{}) {}

// +gen define_inbound:"Pkt5,SzFixed"
type InboundWindowFocus protocol.InboundWindowFocus

func (struc *InboundWindowFocus) Decode(r io.Reader, flags interface{}) {}

// +gen define_inbound:"Pkt1,SzVar16"
type InboundKeyPress protocol.InboundKeyPress

func (struc *InboundKeyPress) Decode(r io.Reader, flags interface{}) {}

// +gen define_inbound:"Pkt82,SzFixed"
type InboundCameraMovement protocol.InboundCameraMovement

func (struc *InboundCameraMovement) Decode(r io.Reader, flags interface{}) {}
