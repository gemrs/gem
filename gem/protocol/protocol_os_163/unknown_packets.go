package protocol_os_163

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
