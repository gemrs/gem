package encoding

import (
	"bytes"
	"io"

	"github.com/gtank/isaac"
)

type FrameSize int

const (
	SzFixed FrameSize = iota
	SzVar8
	SzVar16
)

type PacketHeader struct {
	Number int
	Size   FrameSize
	Object Codable
}

func (p *PacketHeader) Encode(buf io.Writer, flags interface{}) error {
	var err error
	rand := flags.(*isaac.ISAAC)
	num := Int8(uint32(p.Number) + rand.Rand())

	/* buffer the payload so that we can encode it's size */
	var tmpBuffer bytes.Buffer
	err = p.Object.Encode(&tmpBuffer, flags)
	if err != nil {
		return err
	}
	flippedBuf := bytes.NewBuffer(tmpBuffer.Bytes())

	/* encode the packet number */
	err = num.Encode(buf, IntNilFlag)
	if err != nil {
		return err
	}

	/* encode the packet size */
	switch p.Size {
	case SzFixed:
	case SzVar8:
		size8 := Int8(flippedBuf.Len())
		err = size8.Encode(buf, IntNilFlag)
		if err != nil {
			return err
		}
	case SzVar16:
		size16 := Int16(flippedBuf.Len())
		err = size16.Encode(buf, IntNilFlag)
		if err != nil {
			return err
		}
	}

	/* encode the buffered payload */
	_, err = flippedBuf.WriteTo(buf)
	return err
}

func (p *PacketHeader) Decode(buf io.Reader, flags interface{}) error {
	panic("not implemented")
}
