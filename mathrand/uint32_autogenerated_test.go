package mathrand_test

import (
	"testing"

	"github.com/xaionaro-go/rand/mathrand"
)

func TestUint32AddRotateMultiply(t *testing.T) {
	//testUint32(t, mathrand.GlobalPRNG.Uint32AddRotateMultiply)
}
func BenchmarkUint32AddRotateMultiply(b *testing.B) {
	prng := mathrand.New()
	b.SetBytes(4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prng.Uint32AddRotateMultiply()
	}
}

func TestUint32MultiplyAdd(t *testing.T) {
	//testUint32(t, mathrand.GlobalPRNG.Uint32MultiplyAdd)
}
func BenchmarkUint32MultiplyAdd(b *testing.B) {
	prng := mathrand.New()
	b.SetBytes(4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prng.Uint32MultiplyAdd()
	}
}

func TestUint32AddRotate(t *testing.T) {
	//testUint32(t, mathrand.GlobalPRNG.Uint32AddRotate)
}
func BenchmarkUint32AddRotate(b *testing.B) {
	prng := mathrand.New()
	b.SetBytes(4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prng.Uint32AddRotate()
	}
}

func TestUint32Xorshift(t *testing.T) {
	//testUint32(t, mathrand.GlobalPRNG.Uint32Xorshift)
}
func BenchmarkUint32Xorshift(b *testing.B) {
	prng := mathrand.New()
	b.SetBytes(4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prng.Uint32Xorshift()
	}
}

func TestUint32LFG(t *testing.T) {
	//testUint32(t, mathrand.GlobalPRNG.Uint32LFG)
}
func BenchmarkUint32LFG(b *testing.B) {
	prng := mathrand.New()
	b.SetBytes(4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prng.Uint32LFG()
	}
}

func TestUint32PCG(t *testing.T) {
	//testUint32(t, mathrand.GlobalPRNG.Uint32PCG)
}
func BenchmarkUint32PCG(b *testing.B) {
	prng := mathrand.New()
	b.SetBytes(4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prng.Uint32PCG()
	}
}
