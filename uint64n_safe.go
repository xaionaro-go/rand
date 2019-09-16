package fastrand

import (
	"math"
	"sync/atomic"
)

var uint64nPositionSafe uint64

// Uint64n is a fast analog of rand.Intn, but it returns `uint64`.
func Uint64nSafe(n uint64) uint64 {
	// Just two arbitrary prime numbers:
	pos := atomic.AddUint64(&uint64nPositionSafe, 8963315421273233617)
	pos *= 15396334245663786197
	if n == math.MaxUint64 {
		return pos
	}
	return pos % n
}
