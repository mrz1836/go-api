package parameters

// UniqueUint64 removes duplicates from uint64 arrays
func UniqueUint64(in []uint64) []uint64 {
	found := make(map[uint64]struct{})
	out := make([]uint64, 0, len(in))
	for _, v := range in {
		if _, exists := found[v]; !exists {
			found[v] = struct{}{}
			out = append(out, v)
		}

	}
	return out
}
