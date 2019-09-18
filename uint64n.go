package fastrand

import (
	"math/bits"
	"sync/atomic"
)

const (
	primeNumber64bit0 = 15396334245663786197
	primeNumber64bit1 = 8963315421273233617
)

// Uint64 is a fast analog of `math/rand.Uint64`. This random numbers could
// easily be predicted (it's not an analog of `crypto/rand.Read`).
//
// It's not thread-safe in formal terms (which is usually not important
// for PRNG) and it sometimes may return the same value to concurrent routines.
// If you need a non-copy guarantee among concurrent calls of this function
// (for example for `nonce`) then you can use `*Safe` functions. Or you can use
// different instances of `fastrand.PRNG` for different goroutines. Also if
// you use a random number for example for sampling or from a single goroutine
// then you can safely use this function.
func (prng *PRNG) Uint64() uint64 {
	prng.uint64nPosition += primeNumber64bit1
	prng.uint64nPosition = bits.RotateLeft64(prng.uint64nPosition, 32)
	prng.uint64nPosition *= primeNumber64bit0
	return prng.uint64nPosition
}

// Uint64 is a fast analog of `math/rand.Uint64`. This random numbers could
// easily be predicted (it's not an analog of `crypto/rand.Read`).
//
// It's not thread-safe in formal terms (which is usually not important
// for PRNG) and it sometimes may return the same value to concurrent routines.
// If you need a non-copy guarantee among concurrent calls of this function
// (for example for `nonce`) then you can use `*Safe` functions. Or you can use
// different instances of `fastrand.PRNG` for different goroutines. Also if
// you use a random number for example for sampling or from a single goroutine
// then you can safely use this function.
func Uint64() uint64 {
	return globalPRNG.Uint64()
}

// Uint64n is a fast analog of `math/rand.Intn`, but it returns `uint64`.
//
// This function is equivalent to:
// 	fastrand.ReduceUint64(fastrand.Uint64(), n)
func (prng *PRNG) Uint64n(n uint64) uint64 {
	return ReduceUint64(prng.Uint64(), n)
}

// Uint64n is a fast analog of `math/rand.Intn`, but it returns `uint64`.
//
// This function is equivalent to:
// 	fastrand.ReduceUint64(fastrand.Uint64(), n)
func Uint64n(n uint64) uint64 {
	return globalPRNG.Uint64n(n)
}

// Uint64Safe is a fast analog of `math/rand.Uint64`. This random numbers
// could easily be predicted (it's not an analog of `crypto/rand.Read`).
func (prng *PRNG) Uint64Safe() uint64 {
	pos := atomic.AddUint64(&prng.uint64nPositionSafe, primeNumber64bit1)
	pos *= primeNumber64bit0
	return pos
}

// Uint64Safe is a fast analog of `math/rand.Uint64`. This random numbers
// could easily be predicted (it's not an analog of `crypto/rand.Read`).
func Uint64Safe() uint64 {
	return globalPRNG.Uint64Safe()
}

// Uint64nSafe is a fast analog of `math/rand.Intn`, but it returns `uint64`.
//
// This function is equivalent to:
// 	fastrand.ReduceUint64(prng.Uint64Safe(), n)
func (prng *PRNG) Uint64nSafe(n uint64) uint64 {
	return ReduceUint64(prng.Uint64Safe(), n)
}

// Uint64nSafe is a fast analog of `math/rand.Intn`, but it returns `uint64`.
//
// This function is equivalent to:
// 	fastrand.ReduceUint64(fastrand.Uint64Safe(), n)
func Uint64nSafe(n uint64) uint64 {
	return globalPRNG.Uint64nSafe(n)
}

// Uint64Fast is a very fast (but inaccurate) analog of `math/rand.Uint64`.
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
// (for example for `nonce`) then you can use `*Safe` functions. However if
// you use a random number for example for sampling or from a single goroutine
// then you can safely use this function.
func (prng *PRNG) Uint64Fast() uint64 {
	prng.uint64nPosition += primeNumber64bit1
	prng.uint64nPosition = bits.RotateLeft64(prng.uint64nPosition, 32)
	return prng.uint64nPosition
}

// Uint64Fast is a very fast (but inaccurate) analog of `math/rand.Uint64`.
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
// (for example for `nonce`) then you can use `*Safe` functions. However if
// you use a random number for example for sampling or from a single goroutine
// then you can safely use this function.
func Uint64Fast() uint64 {
	return globalPRNG.Uint64Fast()
}

// Uint64nSafe is a fast analog of `math/rand.Intn`, but it returns `uint64`.
//
// This function is equivalent to:
// 	fastrand.ReduceUint64(prng.Uint64Fast(), n)
func (prng *PRNG) Uint64nFast(n uint64) uint64 {
	return ReduceUint64(Uint64Fast(), n)
}

// Uint64nSafe is a fast analog of `math/rand.Intn`, but it returns `uint64`.
//
// This function is equivalent to:
// 	fastrand.ReduceUint64(fastrand.Uint64Fast(), n)
func Uint64nFast(n uint64) uint64 {
	return globalPRNG.Uint64nFast(n)
}
