package mathrand

const (
	primeNumber32bit0 = 3948558707
	primeNumber32bit1 = 1948560947

	pcgMultiplier = 6364136223846793005 // see https://en.wikipedia.org/wiki/Permuted_congruential_generator
	pcgIncrement  = 1442695040888963407 // see https://en.wikipedia.org/wiki/Permuted_congruential_generator
)

func rotateLeft32(x uint32, k int) uint32 {
	return x<<k | x>>(32-k)
}
func rotateRight32(x uint32, k int) uint32 {
	return x>>k | x<<(32-k)
}

// Uint32AddRotateMultiply is a fast analog of `math/rand.Uint32`.
//
// The reliability on statistical tests of this method is unknown. This is
// improved LCG method, so see also Uint32MultiplyAdd.
func (prng *PRNG) Uint32AddRotateMultiply() uint32 {
	prng.state32[0] += primeNumber32bit1
	prng.state32[0] = rotateLeft32(prng.state32[0], 32)
	prng.state32[0] *= primeNumber32bit0
	return prng.state32[0]
}

// Uint32MultiplyAdd is a fast (but week) analog of `math/rand.Uint32`.
//
// See also: https://en.wikipedia.org/wiki/Linear_congruential_generator
func (prng *PRNG) Uint32MultiplyAdd() uint32 {
	prng.state32[0] *= primeNumber32bit0
	prng.state32[0] += primeNumber32bit1
	return prng.state32[0]
}

// Uint32AddRotate is a very fast (but weak) analog of `math/rand.Uint32`.
//
// The reliability on statistical tests of this method is unknown. However
// it's known that this method generates uneven distribution if you will XOR
// two values generates by this method.
func (prng *PRNG) Uint32AddRotate() uint32 {
	prng.state32[0] += primeNumber32bit1
	prng.state32[0] = rotateLeft32(prng.state32[0], 32)
	return prng.state32[0]
}

// Uint32Xorshift is a very fast (but weak) analog of `math/rand.Uint32`.
//
// See also: https://en.wikipedia.org/wiki/Xorshift
func (prng *PRNG) Uint32Xorshift() uint32 {
	prng.state32[0] ^= prng.state32[0] << 13
	prng.state32[0] ^= prng.state32[0] >> 17
	prng.state32[0] ^= prng.state32[0] << 5
	return prng.state32[0]
}

// See also: https://en.wikipedia.org/wiki/Lagged_Fibonacci_generator
func (prng *PRNG) Uint32LFG() uint32 {
	newPosition := prng.state32[0] * prng.state32[1]
	prng.state32[1] = prng.state32[0]
	prng.state32[0] = newPosition
	return newPosition
}

// See also: https://en.wikipedia.org/wiki/Permuted_congruential_generator
func (prng *PRNG) Uint32PCG() uint32 {
	x := prng.pcgState
	count := int(x >> 59)
	prng.pcgState = x*pcgMultiplier + pcgIncrement
	x ^= x >> 18
	return rotateRight32(uint32(x>>27), count)
}
