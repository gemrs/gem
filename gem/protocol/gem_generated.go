// Code generated by gem_gen; DO NOT EDIT.
package protocol

import (
	"io"

	"github.com/gemrs/gem/gem/game/server"
)

func (o OutboundRemoveGroundItem) Encode(w io.Writer, flags interface{}) {
	packet := server.Proto.Encode(o)
	if packet != nil {
		packet.Encode(w, flags)
	}
}

func (o PlayerUpdate) Encode(w io.Writer, flags interface{}) {
	packet := server.Proto.Encode(o)
	if packet != nil {
		packet.Encode(w, flags)
	}
}

func (o OutboundLogout) Encode(w io.Writer, flags interface{}) {
	packet := server.Proto.Encode(o)
	if packet != nil {
		packet.Encode(w, flags)
	}
}

func (o OutboundPlayerChatMessage) Encode(w io.Writer, flags interface{}) {
	packet := server.Proto.Encode(o)
	if packet != nil {
		packet.Encode(w, flags)
	}
}

func (o OutboundSetText) Encode(w io.Writer, flags interface{}) {
	packet := server.Proto.Encode(o)
	if packet != nil {
		packet.Encode(w, flags)
	}
}

func (o OutboundCreateGlobalGroundItem) Encode(w io.Writer, flags interface{}) {
	packet := server.Proto.Encode(o)
	if packet != nil {
		packet.Encode(w, flags)
	}
}

func (o OutboundPlayerInit) Encode(w io.Writer, flags interface{}) {
	packet := server.Proto.Encode(o)
	if packet != nil {
		packet.Encode(w, flags)
	}
}

func (o OutboundCreateGroundItem) Encode(w io.Writer, flags interface{}) {
	packet := server.Proto.Encode(o)
	if packet != nil {
		packet.Encode(w, flags)
	}
}

func (o OutboundSectorUpdate) Encode(w io.Writer, flags interface{}) {
	packet := server.Proto.Encode(o)
	if packet != nil {
		packet.Encode(w, flags)
	}
}

func (o OutboundUpdateInventoryItem) Encode(w io.Writer, flags interface{}) {
	packet := server.Proto.Encode(o)
	if packet != nil {
		packet.Encode(w, flags)
	}
}

func (o OutboundLoginResponse) Encode(w io.Writer, flags interface{}) {
	packet := server.Proto.Encode(o)
	if packet != nil {
		packet.Encode(w, flags)
	}
}

func (o OutboundChatMessage) Encode(w io.Writer, flags interface{}) {
	packet := server.Proto.Encode(o)
	if packet != nil {
		packet.Encode(w, flags)
	}
}

func (o OutboundRegionUpdate) Encode(w io.Writer, flags interface{}) {
	packet := server.Proto.Encode(o)
	if packet != nil {
		packet.Encode(w, flags)
	}
}

func (o OutboundSkill) Encode(w io.Writer, flags interface{}) {
	packet := server.Proto.Encode(o)
	if packet != nil {
		packet.Encode(w, flags)
	}
}

func (o OutboundTabInterface) Encode(w io.Writer, flags interface{}) {
	packet := server.Proto.Encode(o)
	if packet != nil {
		packet.Encode(w, flags)
	}
}
