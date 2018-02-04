package rt3

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
)

type FSIndex struct {
	Length     int
	StartBlock int
}

func (struc *FSIndex) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)

	struc.Length = buf.GetU24()
	struc.StartBlock = buf.GetU24()
}

type FSBlock struct {
	FileID       int
	FilePosition int
	NextBlock    int
	Partition    int
	Data         []byte
}

func (struc *FSBlock) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)

	struc.FileID = buf.GetU16()
	struc.FilePosition = buf.GetU16()
	struc.NextBlock = buf.GetU24()
	struc.Partition = buf.GetU8()

	struc.Data = buf.GetBytes(512)
}

type FSBlockExt struct {
	*FSBlock
}

func (struc *FSBlockExt) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)

	struc.FileID = buf.GetU32()
	struc.FilePosition = buf.GetU16()
	struc.NextBlock = buf.GetU24()
	struc.Partition = buf.GetU8()

	struc.Data = buf.GetBytes(512)
}
