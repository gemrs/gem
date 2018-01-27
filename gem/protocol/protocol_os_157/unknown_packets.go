package protocol_os_157

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
	p.Number = int(shiftedNumber)

	/* decode the packet size */
	sz, ok := unknownPacketLengths[p.Number]
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
var unknownPacketLengths = map[int]int{
	0:  0,
	1:  -2,
	2:  0,
	3:  6,
	4:  -1,
	5:  1,
	6:  4,
	7:  4,
	8:  5,
	9:  10,
	10: -1,
	11: -1,
	12: 16,
	13: 16,
	14: 8,
	15: 8,
	16: 8,
	17: 8,
	18: 8,
	19: 8,
	20: 8,
	21: 8,
	22: 8,
	23: 8,
	24: 7,
	25: 7,
	26: 7,
	27: 7,
	28: 7,
	29: 13,
	30: 15,
	31: 3,
	32: 3,
	33: 3,
	34: 3,
	35: 3,
	36: 9,
	37: 11,
	38: 7,
	39: 7,
	40: 7,
	41: 7,
	42: 7,
	43: 13,
	44: 15,
	45: 3,
	46: 3,
	47: 3,
	48: 3,
	49: 3,
	50: 3,
	51: 3,
	52: 3,
	53: 9,
	54: 11,
	55: 8,
	56: 8,
	57: 8,
	58: 8,
	59: 8,
	60: 14,
	61: 16,
	62: 8,
	63: 8,
	64: 8,
	65: 8,
	66: 8,
	67: 4,
	68: 6,
	69: 0,
	70: 4,
	71: -2,
	72: -1,
	73: 5,
	74: 2,
	75: -1,
	76: -1,
	77: -1,
	78: -1,
	79: 2,
	80: 2,
	81: 2,
	82: 0,
	83: -1,
	84: -1,
	85: 9,
	86: -1,
	87: -1,
	88: 13,
	89: 3,
	90: -2,
	91: -1,
	92: -1,
	93: -1,
	94: -1,
	95: -1,
}
