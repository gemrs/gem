package protocol_os

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
)

type InboundRsaLoginBlock struct {
	Magic     int
	AuthType  int
	ISAACSeed [4]uint32
	AuthData  [8]int
	Password  string
}

func (struc InboundRsaLoginBlock) Encode(buf io.Writer, flags interface{}) {
	panic("not implemented")
}

func (struc *InboundRsaLoginBlock) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)

	struc.Magic = buf.GetU8()
	struc.AuthType = buf.GetU8()

	for i := 0; i < 4; i++ {
		struc.ISAACSeed[i] = uint32(buf.GetU32())
	}

	for i := 0; i < 8; i++ {
		struc.AuthData[i] = buf.GetU8()
	}

	struc.Password = buf.GetStringZ()
}

type InboundXteaLoginBlock struct {
	Username string
}

func (struc InboundXteaLoginBlock) Encode(buf io.Writer, flags interface{}) {
	panic("not implemented")
}

func (struc *InboundXteaLoginBlock) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)

	struc.Username = buf.GetStringZ()
}

type InboundLoginBlock struct {
	LoginType       int
	LoginLen        int
	Revision        int
	SecureBlockSize int
}

func (struc *InboundLoginBlock) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)

	struc.LoginType = buf.GetU8()
	struc.LoginLen = buf.GetU16()
	struc.Revision = buf.GetU32()
	struc.SecureBlockSize = buf.GetU16()
}
