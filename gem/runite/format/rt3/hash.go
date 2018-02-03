package rt3

func Djb2Hash(s string) uint32 {
	hash := uint32(0)
	for _, c := range s {
		hash = uint32(c) + ((hash << 5) - hash)
	}
	return hash
}
