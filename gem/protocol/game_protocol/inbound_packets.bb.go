// Code generated by bbc; DO NOT EDIT.
package game_protocol

import (
	"io"

	"github.com/gemrs/gem/gem/encoding"
)

type InboundChatMessage anonymous_inbound_packets_bb_1

var InboundChatMessageDefinition = encoding.PacketHeader{
	Type:   (*InboundChatMessage)(nil),
	Number: 4,
	Size:   encoding.SzVar8,
}

func (frm *InboundChatMessage) Encode(buf io.Writer, flags interface{}) (err error) {
	struc := (*anonymous_inbound_packets_bb_1)(frm)
	hdr := encoding.PacketHeader{
		Number: InboundChatMessageDefinition.Number,
		Size:   InboundChatMessageDefinition.Size,
		Object: struc,
	}
	return hdr.Encode(buf, flags)
}

func (frm *InboundChatMessage) Decode(buf io.Reader, flags interface{}) (err error) {
	struc := (*anonymous_inbound_packets_bb_1)(frm)
	hdr := encoding.PacketHeader{
		Number: InboundChatMessageDefinition.Number,
		Size:   InboundChatMessageDefinition.Size,
		Object: struc,
	}
	return hdr.Decode(buf, flags)
}

type anonymous_inbound_packets_bb_2 struct {
	Command encoding.JString
}

func (struc *anonymous_inbound_packets_bb_2) Encode(buf io.Writer, flags interface{}) (err error) {
	err = struc.Command.Encode(buf, 0)
	if err != nil {
		return err
	}

	return err
}

func (struc *anonymous_inbound_packets_bb_2) Decode(buf io.Reader, flags interface{}) (err error) {
	err = struc.Command.Decode(buf, 0)
	if err != nil {
		return err
	}

	return err
}

type InboundCommand anonymous_inbound_packets_bb_2

var InboundCommandDefinition = encoding.PacketHeader{
	Type:   (*InboundCommand)(nil),
	Number: 103,
	Size:   encoding.SzVar8,
}

func (frm *InboundCommand) Encode(buf io.Writer, flags interface{}) (err error) {
	struc := (*anonymous_inbound_packets_bb_2)(frm)
	hdr := encoding.PacketHeader{
		Number: InboundCommandDefinition.Number,
		Size:   InboundCommandDefinition.Size,
		Object: struc,
	}
	return hdr.Encode(buf, flags)
}

func (frm *InboundCommand) Decode(buf io.Reader, flags interface{}) (err error) {
	struc := (*anonymous_inbound_packets_bb_2)(frm)
	hdr := encoding.PacketHeader{
		Number: InboundCommandDefinition.Number,
		Size:   InboundCommandDefinition.Size,
		Object: struc,
	}
	return hdr.Decode(buf, flags)
}

type anonymous_inbound_packets_bb_0 struct {
}

func (struc *anonymous_inbound_packets_bb_0) Encode(buf io.Writer, flags interface{}) (err error) {

	return err
}

func (struc *anonymous_inbound_packets_bb_0) Decode(buf io.Reader, flags interface{}) (err error) {

	return err
}

type InboundPing anonymous_inbound_packets_bb_0

var InboundPingDefinition = encoding.PacketHeader{
	Type:   (*InboundPing)(nil),
	Number: 0,
	Size:   encoding.SzFixed,
}

func (frm *InboundPing) Encode(buf io.Writer, flags interface{}) (err error) {
	struc := (*anonymous_inbound_packets_bb_0)(frm)
	hdr := encoding.PacketHeader{
		Number: InboundPingDefinition.Number,
		Size:   InboundPingDefinition.Size,
		Object: struc,
	}
	return hdr.Encode(buf, flags)
}

func (frm *InboundPing) Decode(buf io.Reader, flags interface{}) (err error) {
	struc := (*anonymous_inbound_packets_bb_0)(frm)
	hdr := encoding.PacketHeader{
		Number: InboundPingDefinition.Number,
		Size:   InboundPingDefinition.Size,
		Object: struc,
	}
	return hdr.Decode(buf, flags)
}

type anonymous_inbound_packets_bb_1 struct {
	Effects        encoding.Uint8
	Colour         encoding.Uint8
	EncodedMessage encoding.Bytes
}

func (struc *anonymous_inbound_packets_bb_1) Encode(buf io.Writer, flags interface{}) (err error) {
	err = struc.Effects.Encode(buf, encoding.IntegerFlag(encoding.IntOffset128|encoding.IntReverse))
	if err != nil {
		return err
	}

	err = struc.Colour.Encode(buf, encoding.IntegerFlag(encoding.IntOffset128|encoding.IntReverse))
	if err != nil {
		return err
	}

	err = struc.EncodedMessage.Encode(buf, nil)
	if err != nil {
		return err
	}

	return err
}

func (struc *anonymous_inbound_packets_bb_1) Decode(buf io.Reader, flags interface{}) (err error) {
	err = struc.Effects.Decode(buf, encoding.IntegerFlag(encoding.IntOffset128|encoding.IntReverse))
	if err != nil {
		return err
	}

	err = struc.Colour.Decode(buf, encoding.IntegerFlag(encoding.IntOffset128|encoding.IntReverse))
	if err != nil {
		return err
	}

	err = struc.EncodedMessage.Decode(buf, nil)
	if err != nil {
		return err
	}

	return err
}
