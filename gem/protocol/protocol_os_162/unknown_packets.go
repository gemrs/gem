package protocol_os_162

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
	// FIXME: my 157 client has ISAAC disabled for some reason
	//p.Number = int(shiftedNumber)

	/* decode the packet size */
	sz, ok := inboundPacketLengths[p.Number]
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

/* This horrible table is the list of fixed length packets we don't know about and their lengths */
var inboundPacketLengths = map[int]int{
	0:  -2,
	1:  3,
	2:  3,
	3:  16,
	4:  8,
	5:  8,
	6:  3,
	7:  4,
	8:  2,
	9:  15,
	10: 0,
	11: 2,
	12: 3,
	13: 16,
	14: 8,
	15: 5,
	16: -1,
	17: -1,
	18: 8,
	19: -1,
	20: 8,
	21: -1,
	22: 16,
	23: 8,
	24: 7,
	25: 8,
	26: 6,
	27: 13,
	28: 7,
	29: 7,
	30: 10,
	31: 3,
	32: 4,
	33: 1,
	34: 3,
	35: 11,
	36: 8,
	37: 7,
	38: 2,
	39: 7,
	40: 13,
	41: 7,
	42: 7,
	43: 0,
	44: 7,
	45: -1,
	46: 8,
	47: 7,
	48: 4,
	49: -1,
	50: 3,
	51: -1,
	52: 3,
	53: 3,
	54: -1,
	55: 8,
	56: 8,
	57: 9,
	58: -1,
	59: 0,
	60: 14,
	61: 6,
	62: 3,
	63: 8,
	64: 3,
	65: 8,
	66: 5,
	67: 8,
	68: 8,
	69: 0,
	70: -2,
	71: 11,
	72: -1,
	73: -1,
	74: 4,
	75: 8,
	76: 9,
	77: -1,
	78: -1,
	79: 13,
	80: -2,
	81: -1,
	82: 8,
	83: 3,
	84: -1,
	85: 8,
	86: 8,
	87: 15,
	88: 3,
	89: 2,
	90: 8,
	91: 7,
	92: -1,
	93: -1,
	94: 9,
	95: 3,
}
