package fastrand

import (
	"math"
)

var uint32nPosition uint32

// Uint32n is a fast analog of rand.Intn, but it returns `uint32`.
// It's not thread-safe in formal terms, but it's not important :)
//go:norace
//go:nosplit
func Uint32n(n uint32) uint32 {
	// Just two arbitrary prime numbers:
	uint32nPosition = 3948558707 * (uint32nPosition + 1948560947)
	if n == math.MaxUint32 {
		return uint32nPosition
	}
	return uint32nPosition % n
}
