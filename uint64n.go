package fastrand

import (
	"math"
	"math/bits"
)

const (
	primeNumber64bit0 = 15396334245663786197
	primeNumber64bit1 = 8963315421273233617
)

var uint64nPosition uint64

// Uint64n is a fast analog of rand.Intn, but it returns `uint64`.
//
// It's not thread-safe in formal terms so it sometimes may return the same
// value to concurrent routines.
//go:norace
func Uint64n(n uint64) uint64 {
	uint64nPosition += primeNumber64bit1
	uint64nPosition = bits.RotateLeft64(uint64nPosition, 32)
	if n == math.MaxUint64 {
		return uint64nPosition
	}
	return ReduceUint64(uint64nPosition, n)
}
