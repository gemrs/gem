package protocol

import (
	"github.com/sinusoids/gem/gem/encode"
)

type EmbeddedStruct struct {
	A encode.Int32
	B encode.Int32
	C encode.Int32
}

func (struc *EmbeddedStruct) Encode(buf *bytes.Buffer, flags_ interface{}) (err error) {
	err = struc.A.Encode(buf, encode.IntegerFlag(encode.IntLittleEndian))
	if err != nil {
		return err
	}
	err = struc.B.Encode(buf, encode.IntegerFlag(encode.IntOffset128|encode.IntPDPEndian))
	if err != nil {
		return err
	}
	err = struc.C.Encode(buf, encode.IntegerFlag())
	if err != nil {
		return err
	}

}

func (struc *EmbeddedStruct) Decode(buf *bytes.Buffer, flags_ interface{}) (err error) {
	err = struc.A.Decode(buf, encode.IntegerFlag(encode.IntLittleEndian))
	if err != nil {
		return err
	}
	err = struc.B.Decode(buf, encode.IntegerFlag(encode.IntOffset128|encode.IntPDPEndian))
	if err != nil {
		return err
	}
	err = struc.C.Decode(buf, encode.IntegerFlag())
	if err != nil {
		return err
	}

}

type AnonStruct_45 struct {
	Message  encode.JString
	Values8  [4]encode.Int16
	Values16 [2]encode.Int16
	Struc    EmbeddedStruct
}

func (struc *AnonStruct_45) Encode(buf *bytes.Buffer, flags_ interface{}) (err error) {
	err = struc.Message.Encode(buf, 256)
	if err != nil {
		return err
	}
	for i := 0; i < 4; i++ {
		err = struc.Values8[i].Encode(buf, encode.IntegerFlag())
		if err != nil {
			return err
		}
	}
	for i := 0; i < 2; i++ {
		err = struc.Values16[i].Encode(buf, encode.IntegerFlag())
		if err != nil {
			return err
		}
	}
	err = struc.Struc.Encode(buf, encode.NilFlags)
	if err != nil {
		return err
	}

}

func (struc *AnonStruct_45) Decode(buf *bytes.Buffer, flags_ interface{}) (err error) {
	err = struc.Message.Decode(buf, 256)
	if err != nil {
		return err
	}
	for i := 0; i < 4; i++ {
		err = struc.Values8[i].Decode(buf, encode.IntegerFlag())
		if err != nil {
			return err
		}
	}
	for i := 0; i < 2; i++ {
		err = struc.Values16[i].Decode(buf, encode.IntegerFlag())
		if err != nil {
			return err
		}
	}
	err = struc.Struc.Decode(buf, encode.NilFlags)
	if err != nil {
		return err
	}

}

type TestFrame AnonStruct_45
