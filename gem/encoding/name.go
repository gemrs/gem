package encoding

func NameHash(name string) Uint64 {
	var hash uint64
	runes := ([]rune)(name)
	for i, c := range runes {
		if i >= 12 {
			break
		}

		hash *= 37
		if c >= 'A' && c <= 'Z' {
			hash += (1 + uint64(c)) - 65
		} else if c >= 'a' && c <= 'z' {
			hash += (1 + uint64(c)) - 97
		} else if c >= '0' && c <= '9' {
			hash += (27 + uint64(c)) - 48
		}
	}
	for hash%37 == 0 && hash != 0 {
		hash /= 37
	}
	return Uint64(hash)
}
