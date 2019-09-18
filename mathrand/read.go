package mathrand


// Read is a fast analog of `math/rand.Read`. This mathrandom numbers could easily
// be predicted (it's not an analog of `crypto/mathrand.Read`).
//
// It's not thread-safe in formal terms (which is usually not important
// for PRNG) and it sometimes may return the same value to concurrent routines.
// If you need a non-copy guarantee among concurrent calls of this function
// (for example for `nonce`) then you can use `*Safe` functions. Or you can use
// different instances of `mathrand.PRNG` for different goroutines. Also if
// you use a mathrandom number for example for sampling or from a single goroutine
// then you can safely use this function.

/*
// XORRead XORs argument `b` with a pseudo-mathrandom value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.Read(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
//
// It's not thread-safe in formal terms (which is usually not important
// for PRNG) and it sometimes may return the same value to concurrent routines.
// If you need a non-copy guarantee among concurrent calls of this function
// (for example for `nonce`) then you can use `*Safe` functions. Or you can use
// different instances of `mathrand.PRNG` for different goroutines. Also if
// you use a mathrandom number for example for sampling or from a single goroutine
// then you can safely use this function.
func (prng *PRNG) XORRead(b []byte) (n int, err error) {
	return prng.xorRead(b, prng.uint64nPosition^uint64(prng.uint32nPosition))
}

// XORRead XORs argument `b` with a pseudo-mathrandom value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.Read(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
//
// It's not thread-safe in formal terms (which is usually not important
// for PRNG) and it sometimes may return the same value to concurrent routines.
// If you need a non-copy guarantee among concurrent calls of this function
// (for example for `nonce`) then you can use `*Safe` functions. Or you can use
// different instances of `mathrand.PRNG` for different goroutines. Also if
// you use a mathrandom number for example for sampling or from a single goroutine
// then you can safely use this function.
func XORRead(b []byte) (n int, err error) {
	return globalPRNG.XORRead(b)
}

// XORReadSafe XORs argument `b` with a pseudo-mathrandom value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.ReadSafe(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func (prng *PRNG) XORReadSafe(b []byte) (n int, err error) {
	return prng.xorRead(b, prng.uint64nPosition^(primeNumber64bit0*Uint64nSafe(math.MaxUint64)))
}

// XORReadSafe XORs argument `b` with a pseudo-mathrandom value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.ReadSafe(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func XORReadSafe(b []byte) (n int, err error) {
	return globalPRNG.XORReadSafe(b)
}

//go:norace
func (prng *PRNG) readFast(b []byte, pos uint64) (n int, err error) {
	l := len(b)
	var i int
	{
		lPos := pos
		for i = 8; i < l; i += 8 {
			lPos += primeNumber64bit1
			lPos = bits.RotateLeft64(lPos, 32)
			*(*uint64)((unsafe.Pointer)(&b[i-8])) = lPos
		}
		pos ^= primeNumber64bit0 * (lPos + primeNumber64bit1)
	}
	i -= 8
	prng.uint64nPosition ^= pos * primeNumber64bit0
	if l > (1 << 6) {
		prng.uint32nPosition ^= uint32(pos * primeNumber32bit0)
	}
	rest := pos
	if i == l {
		return
	}
	if l-i >= 4 {
		*(*uint32)((unsafe.Pointer)(&b[i])) = uint32(rest)
		rest >>= 32
		i += 4
	}
	if l-i >= 2 {
		*(*uint16)((unsafe.Pointer)(&b[i])) = uint16(rest)
		rest >>= 16
		i += 2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return l, nil
}
*/

/*
//go:norace
*/
