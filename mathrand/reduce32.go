package mathrand

// ReduceUint32 returns `src` as a value smaller than `mod`.
//
// Based on: http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
//
// For fair distribution of results `src` shouldn't be reduced,
// it should be an uint32 mathrandom number [0..2^32).
func ReduceUint32(src, mod uint32) uint32 {
	// http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return uint32(((uint64)(src) * uint64(mod)) >> 32)
}
