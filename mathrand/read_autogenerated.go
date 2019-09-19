package mathrand

import "unsafe"

// ReadUint64AddRotateMultiply is an analog of math/rand.Read. This random numbers could easily
// be predicted (it's not an analog of crypto/rand.Read).
//
// Applied PRNG method:
// Uint64AddRotateMultiply is a fast analog of `math/rand.Uint64`.
//
// The reliability on statistical tests of this method is unknown. This is
// improved LCG method, so see also Uint64MultiplyAdd.
func (prng *PRNG) ReadUint64AddRotateMultiply(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state64Temp0 := prng.state64[0]
		for i = 8; i < l; i += 8 {
			state64Temp0 += primeNumber64bit1
			state64Temp0 = rotateLeft64(state64Temp0, 32)
			state64Temp0 *= primeNumber64bit0
			*(*uint64)((unsafe.Pointer)(&b[i-8])) = state64Temp0
		}
		prng.state64[0] = state64Temp0
	}
	i -= 8
	rest := prng.Uint64AddRotateMultiply()
	if i == l {
		return
	}
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
	return
}

// XORRead XORs argument "b" with a pseudo-random value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.ReadUint64AddRotateMultiply(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func (prng *PRNG) XORReadUint64AddRotateMultiply(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state64Temp0 := prng.state64[0]
		for i = 8; i < l; i += 8 {
			state64Temp0 += primeNumber64bit1
			state64Temp0 = rotateLeft64(state64Temp0, 32)
			state64Temp0 *= primeNumber64bit0
			*(*uint64)((unsafe.Pointer)(&b[i-8])) ^= state64Temp0
		}
		prng.state64[0] = state64Temp0
	}
	i -= 8
	rest := prng.Uint64AddRotateMultiply()
	if i == l {
		return
	}
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
	return
}

// ReadUint64AddNRotateMultiply is an analog of math/rand.Read. This random numbers could easily
// be predicted (it's not an analog of crypto/rand.Read).
//
// Applied PRNG method:
// Uint64AddNRotateMultiply is a fast analog of `math/rand.Uint64`.
//
// The reliability on statistical tests of this method is unknown. This is
// improved LCG method, so see also Uint64MultiplyAdd.
func (prng *PRNG) ReadUint64AddNRotateMultiply(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state64Temp0 := prng.state64[0]
		for i = 8; i < l; i += 8 {
			state64Temp0 += primeNumber64bit1
			r := 28 + int(state64Temp0&0x7)
			state64Temp0 = rotateLeft64(state64Temp0, r)
			state64Temp0 *= primeNumber64bit0
			*(*uint64)((unsafe.Pointer)(&b[i-8])) = state64Temp0
		}
		prng.state64[0] = state64Temp0
	}
	i -= 8
	rest := prng.Uint64AddNRotateMultiply()
	if i == l {
		return
	}
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
	return
}

// XORRead XORs argument "b" with a pseudo-random value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.ReadUint64AddNRotateMultiply(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func (prng *PRNG) XORReadUint64AddNRotateMultiply(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state64Temp0 := prng.state64[0]
		for i = 8; i < l; i += 8 {
			state64Temp0 += primeNumber64bit1
			r := 28 + int(state64Temp0&0x7)
			state64Temp0 = rotateLeft64(state64Temp0, r)
			state64Temp0 *= primeNumber64bit0
			*(*uint64)((unsafe.Pointer)(&b[i-8])) ^= state64Temp0
		}
		prng.state64[0] = state64Temp0
	}
	i -= 8
	rest := prng.Uint64AddNRotateMultiply()
	if i == l {
		return
	}
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
	return
}

// ReadUint64MultiplyAdd is an analog of math/rand.Read. This random numbers could easily
// be predicted (it's not an analog of crypto/rand.Read).
//
// Applied PRNG method:
// Uint64MultiplyAdd is a fast (but week) analog of `math/rand.Uint64`.
//
// See also: https://en.wikipedia.org/wiki/Linear_congruential_generator
func (prng *PRNG) ReadUint64MultiplyAdd(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state64Temp0 := prng.state64[0]
		for i = 8; i < l; i += 8 {
			state64Temp0 *= primeNumber64bit0
			state64Temp0 += primeNumber64bit1
			*(*uint64)((unsafe.Pointer)(&b[i-8])) = state64Temp0
		}
		prng.state64[0] = state64Temp0
	}
	i -= 8
	rest := prng.Uint64MultiplyAdd()
	if i == l {
		return
	}
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
	return
}

// XORRead XORs argument "b" with a pseudo-random value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.ReadUint64MultiplyAdd(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func (prng *PRNG) XORReadUint64MultiplyAdd(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state64Temp0 := prng.state64[0]
		for i = 8; i < l; i += 8 {
			state64Temp0 *= primeNumber64bit0
			state64Temp0 += primeNumber64bit1
			*(*uint64)((unsafe.Pointer)(&b[i-8])) ^= state64Temp0
		}
		prng.state64[0] = state64Temp0
	}
	i -= 8
	rest := prng.Uint64MultiplyAdd()
	if i == l {
		return
	}
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
	return
}

// ReadUint64AddRotate is an analog of math/rand.Read. This random numbers could easily
// be predicted (it's not an analog of crypto/rand.Read).
//
// Applied PRNG method:
// Uint64AddRotate is a very fast (but weak) analog of `math/rand.Uint64`.
//
// The reliability on statistical tests of this method is unknown. However
// it's known that this method generates uneven distribution if you will XOR
// two values generates by this method.
func (prng *PRNG) ReadUint64AddRotate(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state64Temp0 := prng.state64[0]
		for i = 8; i < l; i += 8 {
			state64Temp0 += primeNumber64bit1
			state64Temp0 = rotateLeft64(state64Temp0, 32)
			*(*uint64)((unsafe.Pointer)(&b[i-8])) = state64Temp0
		}
		prng.state64[0] = state64Temp0
	}
	i -= 8
	rest := prng.Uint64AddRotate()
	if i == l {
		return
	}
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
	return
}

// XORRead XORs argument "b" with a pseudo-random value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.ReadUint64AddRotate(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func (prng *PRNG) XORReadUint64AddRotate(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state64Temp0 := prng.state64[0]
		for i = 8; i < l; i += 8 {
			state64Temp0 += primeNumber64bit1
			state64Temp0 = rotateLeft64(state64Temp0, 32)
			*(*uint64)((unsafe.Pointer)(&b[i-8])) ^= state64Temp0
		}
		prng.state64[0] = state64Temp0
	}
	i -= 8
	rest := prng.Uint64AddRotate()
	if i == l {
		return
	}
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
	return
}

// ReadUint64Xorshift is an analog of math/rand.Read. This random numbers could easily
// be predicted (it's not an analog of crypto/rand.Read).
//
// Applied PRNG method:
// Uint64Xorshift is a very fast (but weak) analog of `math/rand.Uint64`.
//
// See also: https://en.wikipedia.org/wiki/Xorshift
func (prng *PRNG) ReadUint64Xorshift(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state64Temp0 := prng.state64[0]
		for i = 8; i < l; i += 8 {
			state64Temp0 ^= state64Temp0 << 13
			state64Temp0 ^= state64Temp0 >> 7
			state64Temp0 ^= state64Temp0 << 17
			*(*uint64)((unsafe.Pointer)(&b[i-8])) = state64Temp0
		}
		prng.state64[0] = state64Temp0
	}
	i -= 8
	rest := prng.Uint64Xorshift()
	if i == l {
		return
	}
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
	return
}

// XORRead XORs argument "b" with a pseudo-random value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.ReadUint64Xorshift(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func (prng *PRNG) XORReadUint64Xorshift(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state64Temp0 := prng.state64[0]
		for i = 8; i < l; i += 8 {
			state64Temp0 ^= state64Temp0 << 13
			state64Temp0 ^= state64Temp0 >> 7
			state64Temp0 ^= state64Temp0 << 17
			*(*uint64)((unsafe.Pointer)(&b[i-8])) ^= state64Temp0
		}
		prng.state64[0] = state64Temp0
	}
	i -= 8
	rest := prng.Uint64Xorshift()
	if i == l {
		return
	}
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
	return
}

// ReadUint64Xoshiro256 is an analog of math/rand.Read. This random numbers could easily
// be predicted (it's not an analog of crypto/rand.Read).
//
// Applied PRNG method:
// See also: https://en.wikipedia.org/wiki/Xorshift#xoshiro_and_xoroshiro
func (prng *PRNG) ReadUint64Xoshiro256(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		var result uint64
		state64Temp1 := prng.state64[1]
		state64Temp2 := prng.state64[2]
		state64Temp0 := prng.state64[0]
		state64Temp3 := prng.state64[3]
		for i = 8; i < l; i += 8 {
			result = rotateLeft64(state64Temp1*5, 7) * 9
			t := state64Temp1 << 17
			state64Temp2 ^= state64Temp0
			state64Temp3 ^= state64Temp1
			state64Temp1 ^= state64Temp2
			state64Temp0 ^= state64Temp3
			state64Temp2 ^= t
			state64Temp3 ^= rotateLeft64(state64Temp3, 45)
			*(*uint64)((unsafe.Pointer)(&b[i-8])) = result
		}
		prng.state64[1] = state64Temp1
		prng.state64[2] = state64Temp2
		prng.state64[0] = state64Temp0
		prng.state64[3] = state64Temp3
	}
	i -= 8
	rest := prng.Uint64Xoshiro256()
	if i == l {
		return
	}
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
	return
}

// XORRead XORs argument "b" with a pseudo-random value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.ReadUint64Xoshiro256(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func (prng *PRNG) XORReadUint64Xoshiro256(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		var result uint64
		state64Temp1 := prng.state64[1]
		state64Temp2 := prng.state64[2]
		state64Temp0 := prng.state64[0]
		state64Temp3 := prng.state64[3]
		for i = 8; i < l; i += 8 {
			result = rotateLeft64(state64Temp1*5, 7) * 9
			t := state64Temp1 << 17
			state64Temp2 ^= state64Temp0
			state64Temp3 ^= state64Temp1
			state64Temp1 ^= state64Temp2
			state64Temp0 ^= state64Temp3
			state64Temp2 ^= t
			state64Temp3 ^= rotateLeft64(state64Temp3, 45)
			*(*uint64)((unsafe.Pointer)(&b[i-8])) ^= result
		}
		prng.state64[1] = state64Temp1
		prng.state64[2] = state64Temp2
		prng.state64[0] = state64Temp0
		prng.state64[3] = state64Temp3
	}
	i -= 8
	rest := prng.Uint64Xoshiro256()
	if i == l {
		return
	}
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
	return
}

// ReadUint64LFG is an analog of math/rand.Read. This random numbers could easily
// be predicted (it's not an analog of crypto/rand.Read).
//
// Applied PRNG method:
// See also: https://en.wikipedia.org/wiki/Lagged_Fibonacci_generator
func (prng *PRNG) ReadUint64LFG(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state64Temp1 := prng.state64[1]
		state64Temp0 := prng.state64[0]
		for i = 8; i < l; i += 8 {
			newPosition := state64Temp0 * state64Temp1
			state64Temp1 = state64Temp0
			state64Temp0 = newPosition
			*(*uint64)((unsafe.Pointer)(&b[i-8])) = newPosition
		}
		prng.state64[1] = state64Temp1
		prng.state64[0] = state64Temp0
	}
	i -= 8
	rest := prng.Uint64LFG()
	if i == l {
		return
	}
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
	return
}

// XORRead XORs argument "b" with a pseudo-random value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.ReadUint64LFG(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func (prng *PRNG) XORReadUint64LFG(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state64Temp1 := prng.state64[1]
		state64Temp0 := prng.state64[0]
		for i = 8; i < l; i += 8 {
			newPosition := state64Temp0 * state64Temp1
			state64Temp1 = state64Temp0
			state64Temp0 = newPosition
			*(*uint64)((unsafe.Pointer)(&b[i-8])) ^= newPosition
		}
		prng.state64[1] = state64Temp1
		prng.state64[0] = state64Temp0
	}
	i -= 8
	rest := prng.Uint64LFG()
	if i == l {
		return
	}
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
	return
}

// ReadUint64MSWS is an analog of math/rand.Read. This random numbers could easily
// be predicted (it's not an analog of crypto/rand.Read).
//
// Applied PRNG method:
// See also: https://en.wikipedia.org/wiki/Middle-square_method
func (prng *PRNG) ReadUint64MSWS(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state64Temp0 := prng.state64[0]
		state64Temp1 := prng.state64[1]
		for i = 8; i < l; i += 8 {
			state64Temp0 *= state64Temp0
			state64Temp1 += mswsSeed
			state64Temp0 += state64Temp1
			state64Temp0 = rotateLeft64(state64Temp0, 32)
			*(*uint64)((unsafe.Pointer)(&b[i-8])) = state64Temp0
		}
		prng.state64[0] = state64Temp0
		prng.state64[1] = state64Temp1
	}
	i -= 8
	rest := prng.Uint64MSWS()
	if i == l {
		return
	}
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
	return
}

// XORRead XORs argument "b" with a pseudo-random value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.ReadUint64MSWS(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func (prng *PRNG) XORReadUint64MSWS(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state64Temp0 := prng.state64[0]
		state64Temp1 := prng.state64[1]
		for i = 8; i < l; i += 8 {
			state64Temp0 *= state64Temp0
			state64Temp1 += mswsSeed
			state64Temp0 += state64Temp1
			state64Temp0 = rotateLeft64(state64Temp0, 32)
			*(*uint64)((unsafe.Pointer)(&b[i-8])) ^= state64Temp0
		}
		prng.state64[0] = state64Temp0
		prng.state64[1] = state64Temp1
	}
	i -= 8
	rest := prng.Uint64MSWS()
	if i == l {
		return
	}
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
	return
}

// ReadUint32AddRotateMultiply is an analog of math/rand.Read. This random numbers could easily
// be predicted (it's not an analog of crypto/rand.Read).
//
// Applied PRNG method:
// Uint32AddRotateMultiply is a fast analog of `math/rand.Uint32`.
//
// The reliability on statistical tests of this method is unknown. This is
// improved LCG method, so see also Uint32MultiplyAdd.
func (prng *PRNG) ReadUint32AddRotateMultiply(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state32Temp0 := prng.state32[0]
		for i = 4; i < l; i += 4 {
			state32Temp0 += primeNumber32bit1
			state32Temp0 = rotateLeft32(state32Temp0, 32)
			state32Temp0 *= primeNumber32bit0
			*(*uint32)((unsafe.Pointer)(&b[i-4])) = state32Temp0
		}
		prng.state32[0] = state32Temp0
	}
	i -= 4
	rest := prng.Uint32AddRotateMultiply()
	if i == l {
		return
	}
	if l-i >= 2 {
		*(*uint16)((unsafe.Pointer)(&b[i])) = uint16(rest)
		rest >>= 16
		i += 2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return
}

// XORRead XORs argument "b" with a pseudo-random value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.ReadUint32AddRotateMultiply(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func (prng *PRNG) XORReadUint32AddRotateMultiply(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state32Temp0 := prng.state32[0]
		for i = 4; i < l; i += 4 {
			state32Temp0 += primeNumber32bit1
			state32Temp0 = rotateLeft32(state32Temp0, 32)
			state32Temp0 *= primeNumber32bit0
			*(*uint32)((unsafe.Pointer)(&b[i-4])) = state32Temp0
		}
		prng.state32[0] = state32Temp0
	}
	i -= 4
	rest := prng.Uint32AddRotateMultiply()
	if i == l {
		return
	}
	if l-i >= 2 {
		*(*uint16)((unsafe.Pointer)(&b[i])) = uint16(rest)
		rest >>= 16
		i += 2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return
}

// ReadUint32MultiplyAdd is an analog of math/rand.Read. This random numbers could easily
// be predicted (it's not an analog of crypto/rand.Read).
//
// Applied PRNG method:
// Uint32MultiplyAdd is a fast (but week) analog of `math/rand.Uint32`.
//
// See also: https://en.wikipedia.org/wiki/Linear_congruential_generator
func (prng *PRNG) ReadUint32MultiplyAdd(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state32Temp0 := prng.state32[0]
		for i = 4; i < l; i += 4 {
			state32Temp0 *= primeNumber32bit0
			state32Temp0 += primeNumber32bit1
			*(*uint32)((unsafe.Pointer)(&b[i-4])) = state32Temp0
		}
		prng.state32[0] = state32Temp0
	}
	i -= 4
	rest := prng.Uint32MultiplyAdd()
	if i == l {
		return
	}
	if l-i >= 2 {
		*(*uint16)((unsafe.Pointer)(&b[i])) = uint16(rest)
		rest >>= 16
		i += 2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return
}

// XORRead XORs argument "b" with a pseudo-random value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.ReadUint32MultiplyAdd(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func (prng *PRNG) XORReadUint32MultiplyAdd(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state32Temp0 := prng.state32[0]
		for i = 4; i < l; i += 4 {
			state32Temp0 *= primeNumber32bit0
			state32Temp0 += primeNumber32bit1
			*(*uint32)((unsafe.Pointer)(&b[i-4])) = state32Temp0
		}
		prng.state32[0] = state32Temp0
	}
	i -= 4
	rest := prng.Uint32MultiplyAdd()
	if i == l {
		return
	}
	if l-i >= 2 {
		*(*uint16)((unsafe.Pointer)(&b[i])) = uint16(rest)
		rest >>= 16
		i += 2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return
}

// ReadUint32AddRotate is an analog of math/rand.Read. This random numbers could easily
// be predicted (it's not an analog of crypto/rand.Read).
//
// Applied PRNG method:
// Uint32AddRotate is a very fast (but weak) analog of `math/rand.Uint32`.
//
// The reliability on statistical tests of this method is unknown. However
// it's known that this method generates uneven distribution if you will XOR
// two values generates by this method.
func (prng *PRNG) ReadUint32AddRotate(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state32Temp0 := prng.state32[0]
		for i = 4; i < l; i += 4 {
			state32Temp0 += primeNumber32bit1
			state32Temp0 = rotateLeft32(state32Temp0, 32)
			*(*uint32)((unsafe.Pointer)(&b[i-4])) = state32Temp0
		}
		prng.state32[0] = state32Temp0
	}
	i -= 4
	rest := prng.Uint32AddRotate()
	if i == l {
		return
	}
	if l-i >= 2 {
		*(*uint16)((unsafe.Pointer)(&b[i])) = uint16(rest)
		rest >>= 16
		i += 2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return
}

// XORRead XORs argument "b" with a pseudo-random value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.ReadUint32AddRotate(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func (prng *PRNG) XORReadUint32AddRotate(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state32Temp0 := prng.state32[0]
		for i = 4; i < l; i += 4 {
			state32Temp0 += primeNumber32bit1
			state32Temp0 = rotateLeft32(state32Temp0, 32)
			*(*uint32)((unsafe.Pointer)(&b[i-4])) = state32Temp0
		}
		prng.state32[0] = state32Temp0
	}
	i -= 4
	rest := prng.Uint32AddRotate()
	if i == l {
		return
	}
	if l-i >= 2 {
		*(*uint16)((unsafe.Pointer)(&b[i])) = uint16(rest)
		rest >>= 16
		i += 2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return
}

// ReadUint32Xorshift is an analog of math/rand.Read. This random numbers could easily
// be predicted (it's not an analog of crypto/rand.Read).
//
// Applied PRNG method:
// Uint32Xorshift is a very fast (but weak) analog of `math/rand.Uint32`.
//
// See also: https://en.wikipedia.org/wiki/Xorshift
func (prng *PRNG) ReadUint32Xorshift(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state32Temp0 := prng.state32[0]
		for i = 4; i < l; i += 4 {
			state32Temp0 ^= state32Temp0 << 13
			state32Temp0 ^= state32Temp0 >> 17
			state32Temp0 ^= state32Temp0 << 5
			*(*uint32)((unsafe.Pointer)(&b[i-4])) = state32Temp0
		}
		prng.state32[0] = state32Temp0
	}
	i -= 4
	rest := prng.Uint32Xorshift()
	if i == l {
		return
	}
	if l-i >= 2 {
		*(*uint16)((unsafe.Pointer)(&b[i])) = uint16(rest)
		rest >>= 16
		i += 2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return
}

// XORRead XORs argument "b" with a pseudo-random value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.ReadUint32Xorshift(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func (prng *PRNG) XORReadUint32Xorshift(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state32Temp0 := prng.state32[0]
		for i = 4; i < l; i += 4 {
			state32Temp0 ^= state32Temp0 << 13
			state32Temp0 ^= state32Temp0 >> 17
			state32Temp0 ^= state32Temp0 << 5
			*(*uint32)((unsafe.Pointer)(&b[i-4])) = state32Temp0
		}
		prng.state32[0] = state32Temp0
	}
	i -= 4
	rest := prng.Uint32Xorshift()
	if i == l {
		return
	}
	if l-i >= 2 {
		*(*uint16)((unsafe.Pointer)(&b[i])) = uint16(rest)
		rest >>= 16
		i += 2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return
}

// ReadUint32LFG is an analog of math/rand.Read. This random numbers could easily
// be predicted (it's not an analog of crypto/rand.Read).
//
// Applied PRNG method:
// See also: https://en.wikipedia.org/wiki/Lagged_Fibonacci_generator
func (prng *PRNG) ReadUint32LFG(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state32Temp1 := prng.state32[1]
		state32Temp0 := prng.state32[0]
		for i = 4; i < l; i += 4 {
			newPosition := state32Temp0 * state32Temp1
			state32Temp1 = state32Temp0
			state32Temp0 = newPosition
			*(*uint32)((unsafe.Pointer)(&b[i-4])) = newPosition
		}
		prng.state32[1] = state32Temp1
		prng.state32[0] = state32Temp0
	}
	i -= 4
	rest := prng.Uint32LFG()
	if i == l {
		return
	}
	if l-i >= 2 {
		*(*uint16)((unsafe.Pointer)(&b[i])) = uint16(rest)
		rest >>= 16
		i += 2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return
}

// XORRead XORs argument "b" with a pseudo-random value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.ReadUint32LFG(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func (prng *PRNG) XORReadUint32LFG(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		state32Temp1 := prng.state32[1]
		state32Temp0 := prng.state32[0]
		for i = 4; i < l; i += 4 {
			newPosition := state32Temp0 * state32Temp1
			state32Temp1 = state32Temp0
			state32Temp0 = newPosition
			*(*uint32)((unsafe.Pointer)(&b[i-4])) = newPosition
		}
		prng.state32[1] = state32Temp1
		prng.state32[0] = state32Temp0
	}
	i -= 4
	rest := prng.Uint32LFG()
	if i == l {
		return
	}
	if l-i >= 2 {
		*(*uint16)((unsafe.Pointer)(&b[i])) = uint16(rest)
		rest >>= 16
		i += 2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return
}

// ReadUint32PCG is an analog of math/rand.Read. This random numbers could easily
// be predicted (it's not an analog of crypto/rand.Read).
//
// Applied PRNG method:
// See also: https://en.wikipedia.org/wiki/Permuted_congruential_generator
func (prng *PRNG) ReadUint32PCG(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		pcgStateTemp := prng.pcgState
		for i = 4; i < l; i += 4 {
			x := pcgStateTemp
			count := int(x >> 59)
			pcgStateTemp = x*pcgMultiplier + pcgIncrement
			x ^= x >> 18
			*(*uint32)((unsafe.Pointer)(&b[i-4])) = rotateRight32(uint32(x>>27), count)
		}
		prng.pcgState = pcgStateTemp
	}
	i -= 4
	rest := prng.Uint32PCG()
	if i == l {
		return
	}
	if l-i >= 2 {
		*(*uint16)((unsafe.Pointer)(&b[i])) = uint16(rest)
		rest >>= 16
		i += 2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return
}

// XORRead XORs argument "b" with a pseudo-random value. The result is
// the same (but faster) as:
//
// 	x := make([]byte, len(b))
// 	mathrand.ReadUint32PCG(x)
// 	for i := range b {
// 		b[i] ^= x[i]
// 	}
func (prng *PRNG) XORReadUint32PCG(b []byte) (l int, err error) {
	l = len(b)
	var i int
	{
		pcgStateTemp := prng.pcgState
		for i = 4; i < l; i += 4 {
			x := pcgStateTemp
			count := int(x >> 59)
			pcgStateTemp = x*pcgMultiplier + pcgIncrement
			x ^= x >> 18
			*(*uint32)((unsafe.Pointer)(&b[i-4])) = rotateRight32(uint32(x>>27), count)
		}
		prng.pcgState = pcgStateTemp
	}
	i -= 4
	rest := prng.Uint32PCG()
	if i == l {
		return
	}
	if l-i >= 2 {
		*(*uint16)((unsafe.Pointer)(&b[i])) = uint16(rest)
		rest >>= 16
		i += 2
	}
	if l-i >= 1 {
		b[i] = uint8(rest)
	}
	return
}
