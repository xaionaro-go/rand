package mathrand

import (
	"time"
)

var (
	GlobalPRNG = New()
)

// PRNG is a instance of a pseudo-mathrandom number generator.
type PRNG struct {
	state64  [4]uint64
	state32  [4]uint32
	pcgState uint64
}

// New creates new instance of pseudo-mathrandom number generator. It
// may be useful to generate separate instances for concurrent goroutines
// to avoid using of `Safe()` functions.
func New() *PRNG {
	return NewWithSeed(uint64(time.Now().UnixNano()))
}

// NewWithSeed is the same as `New` but initializes the PRNG with predefined
// seed.
func NewWithSeed(seed uint64) *PRNG {
	prng := &PRNG{}
	i := uint64(0)
	for idx := range prng.state64 {
		prng.state64[idx] = seed + i
		i++
	}
	for idx := range prng.state32 {
		prng.state32[idx] = uint32(seed + i)
		i++
	}
	prng.pcgState = (seed << 1) + 1 // must be odd
	return prng
}
