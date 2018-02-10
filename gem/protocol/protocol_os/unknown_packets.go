package protocol_os

import (
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/protocol"
)

// UnknownPacket is a Codable which we can use to handle unknown packets gracefully
type UnknownPacket protocol.UnknownPacket

func (p *UnknownPacket) String() string {
	hexStrs := []string{}
	for _, b := range p.Payload {
		hexStrs = append(hexStrs, hex.EncodeToString([]byte{b}))
	}

	return fmt.Sprintf("Packet [%v] size [%v] sizetype [%v] payload [%v]", p.Number, p.Size, p.SizeType, strings.Join(hexStrs, " "))
}

func (p *UnknownPacket) Encode(buf io.Writer, flags interface{}) {
	panic("not implemented")
}

func (p *UnknownPacket) Decode(buf io.Reader, flags interface{}) {
	var err error
	rand := flags.(uint32)

	/* decode the packet number */
	var shiftedNumber encoding.Int8
	shiftedNumber.Decode(buf, encoding.IntNilFlag)

	p.Number = int(uint8(uint32(shiftedNumber) - rand))

	/* decode the packet size */
	sz, ok := InboundSizes[p.Number]
	if !ok {
		panic(fmt.Errorf("unknown length for packet %v", p.Number))
	}

	switch sz {
	case -2:
		var size16 encoding.Int16
		size16.Decode(buf, encoding.IntNilFlag)
		p.Size = int(size16)
		p.SizeType = int(SzVar16)

	case -1:
		var size8 encoding.Int8
		size8.Decode(buf, encoding.IntNilFlag)
		p.Size = int(size8)
		p.SizeType = int(SzVar8)

	default:
		p.SizeType = int(SzFixed)
		p.Size = sz
	}

	/* decode the payload */
	p.Payload = make([]byte, p.Size)
	i, err := buf.Read(p.Payload)
	if err != nil {
		panic(err)
	}
	if i != p.Size {
		panic(io.EOF)
	}
}
