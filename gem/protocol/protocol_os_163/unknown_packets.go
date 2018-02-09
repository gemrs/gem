package protocol_os_163

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
	0:  8,
	1:  -2,
	2:  7,
	3:  6,
	4:  -1,
	5:  1,
	6:  10,
	7:  -1,
	8:  2,
	9:  -1,
	10: -1,
	11: -1,
	12: -1,
	13: -1,
	14: -1,
	15: -1,
	16: 8,
	17: 8,
	18: 16,
	19: -1,
	20: 9,
	21: 3,
	22: 7,
	23: 4,
	24: 3,
	25: 7,
	26: 7,
	27: 11,
	28: -1,
	29: 7,
	30: 3,
	31: 2,
	32: 3,
	33: 7,
	34: 8,
	35: 8,
	36: 3,
	37: 2,
	38: 9,
	39: 8,
	40: 16,
	41: 8,
	42: 8,
	43: 13,
	44: 8,
	45: 7,
	46: 13,
	47: 4,
	48: 8,
	49: 11,
	50: -1,
	51: 7,
	52: 3,
	53: 3,
	54: 3,
	55: 8,
	56: 3,
	57: 8,
	58: 8,
	59: 8,
	60: 16,
	61: 15,
	62: 8,
	63: -1,
	64: -1,
	65: 3,
	66: 3,
	67: 4,
	68: 8,
	69: 5,
	70: 9,
	71: -2,
	72: 2,
	73: 3,
	74: -2,
	75: 3,
	76: 8,
	77: -1,
	78: 7,
	79: 6,
	80: 8,
	81: 7,
	82: 4,
	83: 3,
	84: 0,
	85: 8,
	86: 15,
	87: -1,
	88: 5,
	89: 0,
	90: 0,
	91: 8,
	92: 0,
	93: 14,
	94: -1,
	95: 13,
}
