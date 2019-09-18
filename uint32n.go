package fastrand

import (
	"math/bits"
	"sync/atomic"
)

const (
	primeNumber32bit0 = 3948558707
	primeNumber32bit1 = 1948560947
)

// Uint32 is a fast analog of `math/rand.Uint32`. This random numbers could
// easily be predicted (it's not an analog of `crypto/rand.Read`).
//
// It's not thread-safe in formal terms (which is usually not important
// for PRNG) and it sometimes may return the same value to concurrent routines.
// If you need a non-copy guarantee among concurrent calls of this function
// (for example for `nonce`) then you can use `*Safe` functions. Or you can use
// different instances of `fastrand.PRNG` for different goroutines. Also if
// you use a random number for example for sampling or from a single goroutine
// then you can safely use this function.
func (prng *PRNG) Uint32() uint32 {
	prng.uint32nPosition += primeNumber32bit1
	prng.uint32nPosition = bits.RotateLeft32(prng.uint32nPosition, 16)
	prng.uint32nPosition *= primeNumber32bit0
	return prng.uint32nPosition
}

// Uint32 is a fast analog of `math/rand.Uint32`. This random numbers could
// easily be predicted (it's not an analog of `crypto/rand.Read`).
//
// It's not thread-safe in formal terms (which is usually not important
// for PRNG) and it sometimes may return the same value to concurrent routines.
// If you need a non-copy guarantee among concurrent calls of this function
// (for example for `nonce`) then you can use `*Safe` functions. Or you can use
// different instances of `fastrand.PRNG` for different goroutines. Also if
// you use a random number for example for sampling or from a single goroutine
// then you can safely use this function.
func Uint32() uint32 {
	return globalPRNG.Uint32()
}

// Uint32n is a fast analog of `math/rand.Intn`, but it returns `uint32`.
//
// This function is equivalent to:
// 	fastrand.ReduceUint32(prng.Uint32(), n)
func (prng *PRNG) Uint32n(n uint32) uint32 {
	return ReduceUint32(prng.Uint32(), n)
}

// Uint32n is a fast analog of `math/rand.Intn`, but it returns `uint32`.
//
// This function is equivalent to:
// 	fastrand.ReduceUint32(fastrand.Uint32(), n)
func Uint32n(n uint32) uint32 {
	return globalPRNG.Uint32n(n)
}

// Uint32Safe is a fast analog of `math/rand.Uint32`. This random numbers
// could easily be predicted (it's not an analog of `crypto/rand.Read`).
func (prng *PRNG) Uint32Safe() uint32 {
	pos := atomic.AddUint32(&prng.uint32nPositionSafe, primeNumber32bit1)
	pos *= primeNumber32bit0
	return pos
}

// Uint32Safe is a fast analog of `math/rand.Uint32`. This random numbers
// could easily be predicted (it's not an analog of `crypto/rand.Read`).
func Uint32Safe() uint32 {
	return globalPRNG.Uint32Safe()
}

// Uint32nSafe is a fast analog of `math/rand.Intn`, but it returns `uint32`.
//
// This function is equivalent to:
// 	fastrand.ReduceUint32(fastrand.Uint32Safe(), n)
func (prng *PRNG) Uint32nSafe(n uint32) uint32 {
	return ReduceUint32(prng.Uint32Safe(), n)
}

// Uint32nSafe is a fast analog of `math/rand.Intn`, but it returns `uint32`.
//
// This function is equivalent to:
// 	fastrand.ReduceUint32(fastrand.Uint32Safe(), n)
func Uint32nSafe(n uint32) uint32 {
	return globalPRNG.Uint32nSafe(n)
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
// you are not going to convert this random values then this random numbers
// should be good enough.
//
// It's not thread-safe in formal terms (which is usually not important
// for PRNG) and it sometimes may return the same value to concurrent routines.
// If you need a non-copy guarantee among concurrent calls of this function
// (for example for `nonce`) then you can use `*Safe` functions. Or you can use
// different instances of `fastrand.PRNG` for different goroutines. Also if
// you use a random number for example for sampling or from a single goroutine
// then you can safely use this function.
func (prng *PRNG) Uint32Fast() uint32 {
	prng.uint32nPosition += primeNumber32bit1
	prng.uint32nPosition = bits.RotateLeft32(prng.uint32nPosition, 16)
	return prng.uint32nPosition
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
// you are not going to convert this random values then this random numbers
// should be good enough.
//
// It's not thread-safe in formal terms (which is usually not important
// for PRNG) and it sometimes may return the same value to concurrent routines.
// If you need a non-copy guarantee among concurrent calls of this function
// (for example for `nonce`) then you can use `*Safe` functions. Or you can use
// different instances of `fastrand.PRNG` for different goroutines. Also if
// you use a random number for example for sampling or from a single goroutine
// then you can safely use this function.
func Uint32Fast() uint32 {
	return globalPRNG.Uint32Fast()
}

// Uint32nSafe is a fast analog of `math/rand.Intn`, but it returns `uint32`.
//
// This function is equivalent to:
// 	fastrand.ReduceUint32(prng.Uint32Fast(), n)
func (prng *PRNG) Uint32nFast(n uint32) uint32 {
	return ReduceUint32(prng.Uint32Fast(), n)
}

// Uint32nSafe is a fast analog of `math/rand.Intn`, but it returns `uint32`.
//
// This function is equivalent to:
// 	fastrand.ReduceUint32(fastrand.Uint32Fast(), n)
func Uint32nFast(n uint32) uint32 {
	return globalPRNG.Uint32nFast(n)
}
