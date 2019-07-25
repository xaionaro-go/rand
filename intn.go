package fastrand

import (
	"math"
)

var intnPosition uint32

// Intn is a fast analog of rand.Intn. It's not thread-safe in formal terms, but it's not important :)
func Intn(n uint32) uint32 {
	// Just two arbitrary prime numbers:
	intnPosition = 3948558707 * (intnPosition + 1948560947)
	if n == math.MaxUint32 {
		return intnPosition
	}
	return intnPosition % n
}
