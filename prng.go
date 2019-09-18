package fastrand

import (
	"time"
)

var (
	globalPRNG = New()
)

type PRNG struct {
	uint64nPosition     uint64
	uint64nPositionSafe uint64
	uint32nPosition     uint32
	uint32nPositionSafe uint32
}

func New() *PRNG {
	return NewWithSeed(uint64(time.Now().UnixNano()))
}

func NewWithSeed(seed uint64) *PRNG {
	return &PRNG{
		seed,
		seed + 1,
		uint32(seed + 2),
		uint32(seed + 3),
	}
}
