package mathrand

const (
	primeNumber64bit0 = 15396334245663786197
	primeNumber64bit1 = 8963315421273233617

	mswsSeed = 0xb5ad4eceda1ce2a9
)

func rotateLeft64(x uint64, k int) uint64 {
	return x<<k | x>>(64-k)
}

// Uint64AddRotateMultiply is a fast analog of `math/rand.Uint64`. This mathrandom
// numbers could easily be predicted (it's not an analog of `crypto/mathrand.Read`).
//
// The reliability on statistical tests of this method is unknown. This is
// improved LCG method, so see also Uint64MultiplyAdd.
func (prng *PRNG) Uint64AddRotateMultiply() uint64 {
	// This is my own method. I did not find how it should be called.
	prng.state64[0] += primeNumber64bit1
	prng.state64[0] = rotateLeft64(prng.state64[0], 32)
	prng.state64[0] *= primeNumber64bit0
	return prng.state64[0]
}

// Uint64AddNRotateMultiply is a fast analog of `math/rand.Uint64`. This mathrandom
// numbers could easily be predicted (it's not an analog of `crypto/mathrand.Read`).
//
// The reliability on statistical tests of this method is unknown. This is
// improved LCG method, so see also Uint64MultiplyAdd.
func (prng *PRNG) Uint64AddNRotateMultiply() uint64 {
	// This is my own method. I did not find how it should be called.
	prng.state64[0] += primeNumber64bit1
	r := 28 + int(prng.state64[0]&0x7)
	prng.state64[0] = rotateLeft64(prng.state64[0], r)
	prng.state64[0] *= primeNumber64bit0
	return prng.state64[0]
}

// Uint64MultiplyAdd is a fast (but week) analog of `math/rand.Uint64`. This
// mathrandom numbers could easily be predicted (it's not an analog of `crypto/mathrand.Read`).
//
// See also: https://en.wikipedia.org/wiki/Linear_congruential_generator
func (prng *PRNG) Uint64MultiplyAdd() uint64 {
	prng.state64[0] *= primeNumber64bit0
	prng.state64[0] += primeNumber64bit1
	return prng.state64[0]
}

// Uint64AddRotate is a very fast (but weak) analog of `math/rand.Uint64`.
// This mathrandom numbers could easily be predicted (it's not an analog of
// `crypto/mathrand.Read`).
//
// The reliability on statistical tests of this method is unknown. However
// it's known that this method generates uneven distribution if you will XOR
// two values generates by this method.
func (prng *PRNG) Uint64AddRotate() uint64 {
	prng.state64[0] += primeNumber64bit1
	prng.state64[0] = rotateLeft64(prng.state64[0], 32)
	return prng.state64[0]
}

// Uint64Xorshift is a very fast (but weak) analog of `math/rand.Uint64`.
// This mathrandom numbers could easily be predicted (it's not an analog of
// `crypto/mathrand.Read`).
//
// See also: https://en.wikipedia.org/wiki/Xorshift
func (prng *PRNG) Uint64Xorshift() uint64 {
	prng.state64[0] ^= prng.state64[0] << 13
	prng.state64[0] ^= prng.state64[0] >> 7
	prng.state64[0] ^= prng.state64[0] << 17
	return prng.state64[0]
}

// See also: https://en.wikipedia.org/wiki/Xorshift#xoshiro_and_xoroshiro
func (prng *PRNG) Uint64Xoshiro256() (result uint64) {
	result = rotateLeft64(prng.state64[1]*5, 7) * 9
	t := prng.state64[1] << 17
	prng.state64[2] ^= prng.state64[0]
	prng.state64[3] ^= prng.state64[1]
	prng.state64[1] ^= prng.state64[2]
	prng.state64[0] ^= prng.state64[3]
	prng.state64[2] ^= t
	prng.state64[3] ^= rotateLeft64(prng.state64[3], 45)
	return
}

// See also: https://en.wikipedia.org/wiki/Lagged_Fibonacci_generator
func (prng *PRNG) Uint64LFG() uint64 {
	newPosition := prng.state64[0] * prng.state64[1]
	prng.state64[1] = prng.state64[0]
	prng.state64[0] = newPosition
	return newPosition
}

// See also: https://en.wikipedia.org/wiki/Middle-square_method
func (prng *PRNG) Uint64MSWS() uint64 {
	prng.state64[0] *= prng.state64[0]
	prng.state64[1] += mswsSeed
	prng.state64[0] += prng.state64[1]
	prng.state64[0] = rotateLeft64(prng.state64[0], 32)
	return prng.state64[0]
}
