package protocol_os_162

import (
	"io"

	"github.com/gemrs/gem/gem/protocol"
)

// +gen define_inbound:"Pkt58,SzVar8"
type InboundMouseMovement protocol.InboundMouseMovement

func (struc *InboundMouseMovement) Decode(r io.Reader, flags interface{}) {}

// +gen define_inbound:"Pkt61,SzFixed"
type InboundMouseClick protocol.InboundMouseClick

func (struc *InboundMouseClick) Decode(r io.Reader, flags interface{}) {}

// +gen define_inbound:"Pkt43,SzFixed"
type InboundPing protocol.InboundPing

func (struc *InboundPing) Decode(r io.Reader, flags interface{}) {}

// +gen define_inbound:"Pkt33,SzFixed"
type InboundWindowFocus protocol.InboundWindowFocus

func (struc *InboundWindowFocus) Decode(r io.Reader, flags interface{}) {}

// +gen define_inbound:"Pkt80,SzVar16"
type InboundKeyPress protocol.InboundKeyPress

func (struc *InboundKeyPress) Decode(r io.Reader, flags interface{}) {}

// +gen define_inbound:"Pkt48,SzFixed"
type InboundCameraMovement protocol.InboundCameraMovement

func (struc *InboundCameraMovement) Decode(r io.Reader, flags interface{}) {}
