//+build amd64

package mathrand

import (
	_ "unsafe"
)

//go:linkname runtime_cputicks runtime.cputicks
func runtime_cputicks() int64

func (prng *PRNG) Uint64CpuTicks() uint64 {
	return uint64(runtime_cputicks())
}
