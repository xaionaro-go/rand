package fastrand

import (
	"math"
	"math/bits"
	"sync/atomic"
)

var uint64nPositionSafe uint64

// Uint64n is a fast analog of rand.Intn, but it returns `uint64`.
func Uint64nSafe(n uint64) uint64 {
	pos := atomic.AddUint64(&uint64nPositionSafe, 8963315421273233617)
	pos = bits.RotateLeft64(pos, 32)
	if n == math.MaxUint64 {
		return pos
	}
	return ReduceUint64(pos, n)
}
