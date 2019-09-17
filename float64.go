package fastrand

import (
	"math"
)

// Float64 is a fast analog of `math/rand.Float64`
func Float64() float64 {
	return float64(Uint64n(math.MaxUint64)) / float64(math.MaxUint64)
}
