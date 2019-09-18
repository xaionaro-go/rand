package fastrand

import (
	"math"
	"math/bits"
	"unsafe"
)

// Read is a fast analog of `math/rand.Read`.
//
// It's not thread-safe in formal terms so it sometimes may return the same
// value to concurrent routines.
//go:norace
func Read(b []byte) (n int, err error) {
	return read(b, uint64nPosition^uint64(uint32nPosition))
}

// Read is a fast analog of `math/rand.Read`.
//go:norace
func ReadSafe(b []byte) (n int, err error) {
	return read(b, uint64nPosition^(primeNumber64bit0*Uint64nSafe(math.MaxUint64)))
}

//go:norace
func read(b []byte, pos uint64) (n int, err error) {
	l := len(b)
	var i int
	{
		lPos := pos
		for i = 8; i < l; i += 8 {
			lPos += primeNumber64bit1
			lPos = bits.RotateLeft64(lPos, 32)
			*(*uint64)((unsafe.Pointer)(&b[i-8])) = lPos
			if i&0x3ff == 0 {
				uint64nPosition ^= lPos
			}
		}
		pos ^= primeNumber64bit0 * (lPos + primeNumber64bit1)
	}
	i -= 8
	uint64nPosition ^= pos * primeNumber64bit0
	if l > (1 << 6) {
		uint32nPosition ^= uint32(pos * primeNumber64bit1)
	}
	if i == l {
		return
	}
	rest := pos
	if l-i >= 4 {
		*(*uint32)((unsafe.Pointer)(&b[i])) = uint32(rest)
		rest >>= 32
		i += 4
	}
	if l-i >= 2 {
		*(*uint16)((unsafe.Pointer)(&b[i])) = uint16(rest)
		rest >>= 16
		i += 2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return l, nil
}
