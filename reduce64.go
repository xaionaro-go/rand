package fastrand

// ReduceUint64 returns `src` as a value smaller than `mod`.
func ReduceUint64(src, mod uint64) uint64 {
	mask := mod
	mask |= mask >> 1
	mask |= mask >> 2
	mask |= mask >> 4
	mask |= mask >> 8
	mask |= mask >> 16
	mask |= mask >> 32
	src &= mask
	if src < mod {
		return src
	}
	return src & (mask >> 1)
}
