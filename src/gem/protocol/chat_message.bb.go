// Generated by bbc; DO NOT EDIT
package protocol

import (
	"gem/encoding"
	"io"
)

type AnonStruct_X struct {
	Message encoding.JString
}

func (struc *AnonStruct_X) Encode(buf io.Writer, flags interface{}) (err error) {
	err = struc.Message.Encode(buf, 0)
	if err != nil {
		return err
	}

	return err
}

func (struc *AnonStruct_X) Decode(buf io.Reader, flags interface{}) (err error) {
	err = struc.Message.Decode(buf, 0)
	if err != nil {
		return err
	}

	return err
}

type OutboundChatMessage AnonStruct_X

var OutboundChatMessageDefinition = encoding.PacketHeader{
	Number: 253,
	Size:   encoding.SzVar8,
}

func (frm *OutboundChatMessage) Encode(buf io.Writer, flags interface{}) (err error) {
	struc := (*AnonStruct_X)(frm)
	hdr := encoding.PacketHeader{
		Number: OutboundChatMessageDefinition.Number,
		Size:   OutboundChatMessageDefinition.Size,
		Object: struc,
	}
	return hdr.Encode(buf, flags)
}

func (frm *OutboundChatMessage) Decode(buf io.Reader, flags interface{}) (err error) {
	struc := (*AnonStruct_X)(frm)
	hdr := encoding.PacketHeader{
		Number: OutboundChatMessageDefinition.Number,
		Size:   OutboundChatMessageDefinition.Size,
		Object: struc,
	}
	return hdr.Decode(buf, flags)
}
