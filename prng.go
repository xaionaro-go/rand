package fastrand

import (
	"time"
)

var (
	globalPRNG = New()
)

// PRNG is a instance of a pseudo-random number generator.
type PRNG struct {
	uint64nPosition     uint64
	uint64nPositionSafe uint64
	uint32nPosition     uint32
	uint32nPositionSafe uint32
}

// New creates new instance of pseudo-random number generator. It
// may be useful to generate separate instances for concurrent goroutines
// to avoid using of `Safe()` functions.
func New() *PRNG {
	return NewWithSeed(uint64(time.Now().UnixNano()))
}

// NewWithSeed is the same as `New` but initializes the PRNG with predefined
// seed.
func NewWithSeed(seed uint64) *PRNG {
	return &PRNG{
		seed,
		seed + 1,
		uint32(seed + 2),
		uint32(seed + 3),
	}
}
