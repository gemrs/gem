package encoding

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gemrs/gem/fork/github.com/gtank/isaac"
)

//go:generate stringer -type=FrameSize
type FrameSize int

const (
	SzFixed FrameSize = iota
	SzVar8
	SzVar16
)

type PacketHeader struct {
	Number int
	Size   FrameSize
	Object interface{}
}

func (p PacketHeader) Encode(buf io.Writer, flags interface{}) {
	var err error
	rand := flags.(*isaac.ISAAC)
	num := Int8(uint32(p.Number) + rand.Rand())

	/* buffer the payload so that we can encode it's size */
	var tmpBuffer bytes.Buffer
	if e, ok := p.Object.(Encodable); ok {
		e.Encode(&tmpBuffer, flags)
	} else {
		panic("payload wasn't encodable")
	}

	flippedBuf := bytes.NewBuffer(tmpBuffer.Bytes())

	/* encode the packet number */
	num.Encode(buf, IntNilFlag)

	/* encode the packet size */
	switch p.Size {
	case SzFixed:
	case SzVar8:
		size8 := Int8(flippedBuf.Len())
		size8.Encode(buf, IntNilFlag)
	case SzVar16:
		size16 := Int16(flippedBuf.Len())
		size16.Encode(buf, IntNilFlag)
	}

	/* encode the buffered payload */
	_, err = flippedBuf.WriteTo(buf)
	if err != nil {
		panic(err)
	}
}

func (p *PacketHeader) Decode(buf io.Reader, flags interface{}) {
	rand := flags.(uint32)

	if p.Object == nil {
		panic("no destination object in packet decode")
	}

	/* decode the packet number */
	var shiftedNumber Int8
	shiftedNumber.Decode(buf, IntNilFlag)

	number := uint8(uint32(shiftedNumber) - rand)

	if int(number) != p.Number {
		panic(fmt.Sprintf("packet number mismatch. Got %v, expected %v", int(number), p.Number))
	}

	/* decode the packet size */
	//TODO: verify size matches expected
	switch p.Size {
	case SzFixed:
	case SzVar8:
		var size8 Int8
		size8.Decode(buf, IntNilFlag)
		p.Size = FrameSize(size8)
	case SzVar16:
		var size16 Int16
		size16.Decode(buf, IntNilFlag)
		p.Size = FrameSize(size16)
	}

	/* decode the payload */
	if d, ok := p.Object.(Decodable); ok {
		d.Decode(buf, &PacketHeader{
			Number: p.Number,
			Size:   p.Size,
		})
	} else {
		panic(fmt.Errorf("payload wasn't decodable: %#v", p.Object))
	}

}
