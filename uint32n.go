package fastrand

import (
	"math/bits"
	"sync/atomic"
)

const (
	primeNumber32bit0 = 3948558707
	primeNumber32bit1 = 1948560947
)

var (
	uint32nPosition     uint32
	uint32nPositionSafe uint32
)

// Uint32 is a fast analog of `math/rand.Uint32`. This random numbers could
// easily be predicted (it's not an analog of `crypto/rand.Read`).
//
// It's not thread-safe in formal terms (which is usually not important
// for PRNG) and it sometimes may return the same value to concurrent routines.
// If you need a non-copy guarantee among concurrent calls of this function
// (for example for `nonce`) then you can use `*Safe` functions. However if
// you use a random number for example for sampling or from a single goroutine
// then you can safely use this function.
func Uint32() uint32 {
	uint32nPosition += primeNumber32bit1
	uint32nPosition = bits.RotateLeft32(uint32nPosition, 16)
	uint32nPosition *= primeNumber32bit0
	return uint32nPosition
}

// Uint32n is a fast analog of `math/rand.Intn`, but it returns `uint32`.
//
// This function is equivalent to `fastrand.ReduceUint32(fastrand.Uint32(), n)`
func Uint32n(n uint32) uint32 {
	return ReduceUint32(Uint32(), n)
}

// Uint32Safe is a fast analog of `math/rand.Uint32`. This random numbers
// could easily be predicted (it's not an analog of `crypto/rand.Read`).
func Uint32Safe() uint32 {
	pos := atomic.AddUint32(&uint32nPositionSafe, primeNumber32bit1)
	pos *= primeNumber32bit0
	return pos
}

// Uint32nSafe is a fast analog of `math/rand.Intn`, but it returns `uint32`.
//
// This function is equivalent to `fastrand.ReduceUint32(fastrand.Uint32Safe(), n)`
func Uint32nSafe(n uint32) uint32 {
	return ReduceUint32(Uint32Safe(), n)
}

// Uint32Fast is a very fast (but inaccurate) analog of `math/rand.Uint32`.
// This random numbers could easily be predicted (it's not an analog of
// `crypto/rand.Read`).
//
// "Inaccurate" means that the resulting random numbers may easily reveal
// some anomalies after some conversions. For example if you will sequentially
// generate two blocks of such random numbers and XOR them then the resulting
// (XOR-ed) block will reveal an anomaly (the distribution of values will be
// essentially less even than expected on good random numbers). However if
// not going to convert this random values then this random numbers should be
// good enough.
//
// It's not thread-safe in formal terms (which is usually not important
// for PRNG) and it sometimes may return the same value to concurrent routines.
// If you need a non-copy guarantee among concurrent calls of this function
// (for example for `nonce`) then you can use `*Safe` functions. However if
// you use a random number for example for sampling or from a single goroutine
// then you can safely use this function.
func Uint32Fast() uint32 {
	uint32nPosition += primeNumber32bit1
	uint32nPosition = bits.RotateLeft32(uint32nPosition, 16)
	return uint32nPosition
}

// Uint32nSafe is a fast analog of `math/rand.Intn`, but it returns `uint32`.
//
// This function is equivalent to `fastrand.ReduceUint32(fastrand.Uint32Fast(), n)`
func Uint32nFast(n uint32) uint32 {
	return ReduceUint32(Uint32Fast(), n)
}
