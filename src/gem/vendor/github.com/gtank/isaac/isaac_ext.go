package isaac

func (r *ISAAC) SeedArray(key []uint32) {
	for i, k := range key {
		r.randrsl[i] = k
	}
	r.randInit(true)
}
