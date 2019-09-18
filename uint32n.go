package fastrand

import (
	"math"
	"math/bits"
)

const (
	primeNumber32bit0 = 1948560947
)

var uint32nPosition uint32

// Uint32n is a fast analog of rand.Intn, but it returns `uint32`.
//
// It's not thread-safe in formal terms so it sometimes may return the same
// value to concurrent routines.
//go:norace
func Uint32n(n uint32) uint32 {
	uint32nPosition += primeNumber32bit0
	uint32nPosition = bits.RotateLeft32(uint32nPosition, 16)
	if n == math.MaxUint32 {
		return uint32nPosition
	}
	return ReduceUint32(uint32nPosition, n)
}
