package encoding

// Ported from hyperion

var XlateTable = []byte{
	' ', 'e', 't', 'a', 'o', 'i', 'h', 'n', 's', 'r',
	'd', 'l', 'u', 'm', 'w', 'c', 'y', 'f', 'g', 'p',
	'b', 'v', 'k', 'x', 'j', 'q', 'z', '0', '1', '2',
	'3', '4', '5', '6', '7', '8', '9', ' ', '!', '?',
	'.', ',', ':', ';', '(', ')', '-', '&', '*', '\\',
	'\'', '@', '#', '+', '=', '\243', '$', '%', '"', '[',
	']',
}

const SentenceMarkers = ".?!"

func ChatTextUnpack(packed []byte) string {
	size := len(packed)
	buf := make([]byte, 4096)
	idx := 0
	highNibble := -1
	for i := 0; i < size*2; i++ {
		val := int(packed[i/2] >> uint(4-4*(i%2)) & 0xF)
		if highNibble == -1 {
			if val < 13 {
				buf[idx] = XlateTable[val]
				idx++
			} else {
				highNibble = int(val)
			}
		} else {
			x := ((highNibble << 4) + val) - 195
			buf[idx] = XlateTable[x]
			idx++
			highNibble = -1
		}
	}
	return string(buf[:idx])
}

func ChatTextSanitize(message string) string {
	buf := ([]rune)(message)
	for i := range buf {
		valid := false
		for _, r := range XlateTable {
			if buf[i] == rune(r) {
				valid = true
				break
			}
		}
		if !valid {
			buf[i] = ' '
		}
	}
	message = string(buf)
	// TODO uppercase etc.
	//	sentences := strings.Split(message, SentenceMarkers)
	return message
}

func ChatTextPack(message string) []byte {
	if len(message) > 80 {
		message = message[:80]
	}
	packed := make([]byte, 0)

	carryOverNibble := -1
	for idx := 0; idx < len(message); idx++ {
		c := message[idx]
		tableIdx := 0
		for i := 0; i < len(XlateTable); i++ {
			if c == XlateTable[i] {
				tableIdx = i
				break
			}
		}
		if tableIdx > 12 {
			tableIdx += 195
		}
		if carryOverNibble == -1 {
			if tableIdx < 13 {
				carryOverNibble = tableIdx
			} else {
				packed = append(packed, byte(tableIdx))
			}
		} else if tableIdx < 13 {
			packed = append(packed, byte((carryOverNibble<<4)+tableIdx))
			carryOverNibble = -1
		} else {
			packed = append(packed, byte((carryOverNibble<<4)+(tableIdx>>4)))
			carryOverNibble = tableIdx & 0xF
		}
	}
	if carryOverNibble != -1 {
		packed = append(packed, byte(carryOverNibble<<4))
	}
	reversed := make([]byte, len(packed))
	for i := 0; i < len(packed); i++ {
		reversed[len(packed)-i-1] = packed[i]
	}
	return reversed
}
