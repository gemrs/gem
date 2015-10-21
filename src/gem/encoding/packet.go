package encoding

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gtank/isaac"
)

//go:generate stringer -type=FrameSize
type FrameSize int

const (
	SzFixed FrameSize = iota
	SzVar8
	SzVar16
)

type PacketHeader struct {
	Type   Codable
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
	var err error
	rand := flags.(uint32)

	if p.Object == nil {
		panic("no destination object in packet decode")
	}

	/* decode the packet number */
	var shiftedNumber Int8
	err = shiftedNumber.Decode(buf, IntNilFlag)
	if err != nil {
		return err
	}

	number := uint8(uint32(shiftedNumber) - rand)

	if int(number) != p.Number {
		return fmt.Errorf("packet number mismatch. Got %v, expected %v", int(number), p.Number)
	}

	/* decode the packet size */
	//TODO: verify size matches expected
	switch p.Size {
	case SzFixed:
	case SzVar8:
		var size8 Int8
		err = size8.Decode(buf, IntNilFlag)
		if err != nil {
			return err
		}
		p.Size = FrameSize(size8)
	case SzVar16:
		var size16 Int16
		err = size16.Decode(buf, IntNilFlag)
		if err != nil {
			return err
		}
		p.Size = FrameSize(size16)
	}

	/* decode the payload */
	return p.Object.Decode(buf, &PacketHeader{
		Number: p.Number,
		Size:   p.Size,
	})
}
