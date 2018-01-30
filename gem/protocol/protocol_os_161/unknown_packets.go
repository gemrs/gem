package protocol_os_161

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
	0:  3,
	1:  4,
	2:  9,
	3:  6,
	4:  -1,
	5:  7,
	6:  -1,
	7:  2,
	8:  15,
	9:  10,
	10: 16,
	11: 9,
	12: 2,
	13: 16,
	14: 8,
	15: 3,
	16: 8,
	17: 15,
	18: 11,
	19: 3,
	20: -1,
	21: 14,
	22: 7,
	23: 8,
	24: 7,
	25: 0,
	26: 7,
	27: 3,
	28: 2,
	29: 11,
	30: 5,
	31: 4,
	32: 3,
	33: 0,
	34: 8,
	35: -2,
	36: 3,
	37: 13,
	38: 2,
	39: 8,
	40: 7,
	41: 1,
	42: 7,
	43: 8,
	44: 4,
	45: 3,
	46: 3,
	47: 3,
	48: 13,
	49: 5,
	50: -1,
	51: 8,
	52: 3,
	53: 9,
	54: -1,
	55: 8,
	56: -1,
	57: 8,
	58: 16,
	59: 3,
	60: 7,
	61: 3,
	62: 7,
	63: 0,
	64: 7,
	65: -1,
	66: 13,
	67: 8,
	68: 6,
	69: 8,
	70: -1,
	71: -2,
	72: -1,
	73: 7,
	74: 8,
	75: 3,
	76: 8,
	77: 8,
	78: 3,
	79: -1,
	80: 8,
	81: 8,
	82: 0,
	83: -1,
	84: 4,
	85: -2,
	86: -1,
	87: 8,
	88: -1,
	89: 8,
	90: 8,
	91: -1,
	92: -1,
	93: -1,
	94: -1,
	95: 8,
}
