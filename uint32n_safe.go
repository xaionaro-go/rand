package fastrand

import (
	"math"
	"math/bits"
	"sync/atomic"
)

var uint32nPositionSafe uint32

// Uint32n is a fast analog of rand.Intn, but it returns `uint32`.
func Uint32nSafe(n uint32) uint32 {
	pos := atomic.AddUint32(&uint32nPositionSafe, 1948560947)
	pos = bits.RotateLeft32(pos, 16)
	if n == math.MaxUint32 {
		return pos
	}
	return pos % n
}
