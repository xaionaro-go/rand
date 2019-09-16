package fastrand

import (
	"math"
	"unsafe"
)

// Read is a fast analog of `math.Read`.
// It's not thread-safe in formal terms, but it's not important :)
//go:norace
func Read(b []byte) (n int, err error) {
	return read(b, uint64nPosition^uint64(uint32nPosition))
}

// Read is a fast analog of `math.Read`.
//go:norace
func ReadSafe(b []byte) (n int, err error) {
	return read(b, uint64nPosition ^ (15396334245663786197 * Uint64nSafe(math.MaxUint64)))
}

//go:norace
func read(b []byte, pos uint64) (n int, err error) {
	l := len(b)
	var i int
	{
		lPos := pos
		for i = 8; i < l; i += 8 {
			lPos = 15396334245663786197 * (lPos + 8963315421273233617)
			*(*uint64)((unsafe.Pointer)(&b[i-8])) = lPos
			if i&0x3f == 0 {
				uint64nPosition ^= lPos
			}
		}
		pos = 15396334245663786197 * (lPos + 8963315421273233617)
	}
	i -= 8
	uint64nPosition ^= pos * 15396334245663786197
	if l > (1 << 6) {
		uint32nPosition ^= uint32(pos * 8963315421273233617)
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
