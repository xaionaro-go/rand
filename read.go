package fastrand

import (
	"encoding/binary"
	"math"
)

var (
	byteOrder = binary.LittleEndian
)

func Read(b []byte) (n int, err error) {
	var i int
	l := len(b)
	for i=8; i<l; i+=8 {
		byteOrder.PutUint64(b[i-8:], Uint64n(math.MaxUint64))
	}
	i -= 8
	if i == l {
		return
	}
	rest := Uint64n(math.MaxUint64)
	if l-i >= 4 {
		byteOrder.PutUint32(b[i:], uint32(rest))
		rest >>= 32
		i+=4
	}
	if l-i >= 2 {
		byteOrder.PutUint16(b[i:], uint16(rest))
		rest >>= 16
		i+=2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return l, nil
}
