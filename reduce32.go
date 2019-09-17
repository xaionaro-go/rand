package fastrand

func ReduceUint32(src, mod uint32) uint32 {
	// http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return uint32(((uint64)(src) * uint64(mod)) >> 32)
}