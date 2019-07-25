package fastrand

import (
	"math"
)

func Float64() float64 {
	return float64(Intn(math.MaxUint32)) / float64(math.MaxUint32)
}
