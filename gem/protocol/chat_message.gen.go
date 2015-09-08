package protocol

import (
	"github.com/sinusoids/gem/gem/encode"
)

type ChatMessage3 AnonStruct_67

type VariableChatMessage struct {
	Length  encode.Int16
	Message encode.JString
}

func (struc *VariableChatMessage) Encode(buf *bytes.Buffer, flags_ interface{}) (err error) {
	err = struc.Length.Encode(buf, encode.IntegerFlag(encode.IntOffset128|encode.IntLittleEndian))
	if err != nil {
		return err
	}
	err = struc.Message.Encode(buf, encode.NilFlags)
	if err != nil {
		return err
	}

}

func (struc *VariableChatMessage) Decode(buf *bytes.Buffer, flags_ interface{}) (err error) {
	err = struc.Length.Decode(buf, encode.IntegerFlag(encode.IntOffset128|encode.IntLittleEndian))
	if err != nil {
		return err
	}
	err = struc.Message.Decode(buf, encode.NilFlags)
	if err != nil {
		return err
	}

}

type ChatMessage1 VariableChatMessage

type ChatMessage2 VariableChatMessage

type AnonStruct_67 struct {
	Message encode.JString
}

func (struc *AnonStruct_67) Encode(buf *bytes.Buffer, flags_ interface{}) (err error) {
	err = struc.Message.Encode(buf, encode.NilFlags)
	if err != nil {
		return err
	}

}

func (struc *AnonStruct_67) Decode(buf *bytes.Buffer, flags_ interface{}) (err error) {
	err = struc.Message.Decode(buf, encode.NilFlags)
	if err != nil {
		return err
	}

}
