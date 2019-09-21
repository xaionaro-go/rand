//+build amd64

package mathrand

import (
	_ "unsafe"
)

//go:linkname runtime_cputicks runtime.cputicks
func runtime_cputicks() int64

// Uint64CpuTicks returns current value of ticks count on AMD64 CPU
func (prng *PRNG) Uint64CpuTicks() uint64 {
	return uint64(runtime_cputicks())
}
