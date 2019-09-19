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

// New creates new instance of pseudo-mathrandom number generator.
//
// Methods of PRNG are not thread-safe in formal terms (which is usually
// not important for an PRNG) and it sometimes may return the same value to
// concurrent routines. If you need a non-copy guarantee among concurrent
// calls of a method (for example for "nonce") then you can use different
// instances of mathrand.PRNG for different goroutines. Also if you use a
// random number for example for sampling then you can safely use this
// function (because such errors won't make any negative effect to your
// application).
//
// And example for multiple goroutines:
//
// 	var prngPool = sync.Pool{New: func() interface{}{ return mathrand.New() }}
// 	...
// 	prng := prngPool.Get().(*mathrand.PRNG).
// 	prng.ReadUint64Xoshiro256(b)
//  prngPoo.Put(prng)
//
// Note: Random numbers of this PRNG could easily be predicted (it's not an
// analog of crypto/rand).
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
