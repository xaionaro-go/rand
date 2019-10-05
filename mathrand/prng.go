package mathrand

import "time"

var (
	// GlobalPRNG is just an initialized (ready-to-use) PRNG which could
	// be used from anywhere.
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
	prng.SetSeed(seed)
	return prng
}

func (prng *PRNG) SetSeed(seed uint64) {
	i := uint64(0)
	if seed == 0 {
		seed = primeNumber64bit1
	}
	for idx := range prng.state64 {
		prng.state64[idx] = seed
		seed *= primeNumber64bit0
		if prng.state64[idx] == 0 {
			panic(`"seed+i" cannot be zero`)
		}
		i++
	}
	for idx := range prng.state32 {
		prng.state32[idx] = uint32(seed)
		seed *= primeNumber64bit0
		if prng.state32[idx] == 0 {
			panic(`"uint32(seed+i)" cannot be zero`)
		}
		i++
	}
	prng.pcgState = (seed << 1) + 1 // must be odd
}

func (prng *PRNG) Reseed() {
	prng.SetSeed(prng.Uint64Xoshiro256())
}

func (prng *PRNG) fastReseed64() {
	prng.state64[0] = prng.Uint64Xoshiro256()
}

func (prng *PRNG) fastReseed32() {
	prng.state32[0] = uint32(prng.Uint64Xoshiro256())
}

func xorShift64(a uint64) uint64 {
	a ^= a << 13
	a ^= a >> 7
	a ^= a << 17
	return a
}

func xorShift32(a uint32) uint32 {
	a ^= a << 13
	a ^= a >> 17
	a ^= a << 5
	return a
}
