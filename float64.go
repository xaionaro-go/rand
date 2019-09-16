package fastrand

import (
	"math"
)

func Float64() float64 {
	return float64(Uint64n(math.MaxUint64)) / float64(math.MaxUint64)
}
