package fastrand

import (
	"math"
)

var uint64nPosition uint64

// Uint64n is a fast analog of rand.Intn, but it returns `uint64`.
//
// It's not thread-safe in formal terms so it sometimes may return the same
// value to concurrent routines.
//go:norace
func Uint64n(n uint64) uint64 {
	// Just two arbitrary prime numbers:
	uint64nPosition = 15396334245663786197 * (uint64nPosition + 8963315421273233617)
	if n == math.MaxUint64 {
		return uint64nPosition
	}
	return uint64nPosition % n
}
