package mathrand_test

import (
	"testing"

	"github.com/xaionaro-go/rand/mathrand"
)

func TestReadUint64AddRotateMultiply(t *testing.T) {
	//testRead(t, mathrand.GlobalPRNG.ReadUint64AddRotateMultiply)
}
func BenchmarkReadUint64AddRotateMultiply1(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64AddRotateMultiply, 1)
}
func BenchmarkReadUint64AddRotateMultiply16(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64AddRotateMultiply, 16)
}
func BenchmarkReadUint64AddRotateMultiply1024(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64AddRotateMultiply, 1024)
}
func BenchmarkReadUint64AddRotateMultiply65536(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64AddRotateMultiply, 65536)
}
func BenchmarkReadUint64AddRotateMultiply65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, mathrand.GlobalPRNG.ReadUint64AddRotateMultiply, 65536)
}

func TestReadUint64AddNRotateMultiply(t *testing.T) {
	//testRead(t, mathrand.GlobalPRNG.ReadUint64AddNRotateMultiply)
}
func BenchmarkReadUint64AddNRotateMultiply1(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64AddNRotateMultiply, 1)
}
func BenchmarkReadUint64AddNRotateMultiply16(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64AddNRotateMultiply, 16)
}
func BenchmarkReadUint64AddNRotateMultiply1024(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64AddNRotateMultiply, 1024)
}
func BenchmarkReadUint64AddNRotateMultiply65536(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64AddNRotateMultiply, 65536)
}
func BenchmarkReadUint64AddNRotateMultiply65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, mathrand.GlobalPRNG.ReadUint64AddNRotateMultiply, 65536)
}

func TestReadUint64MultiplyAdd(t *testing.T) {
	//testRead(t, mathrand.GlobalPRNG.ReadUint64MultiplyAdd)
}
func BenchmarkReadUint64MultiplyAdd1(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64MultiplyAdd, 1)
}
func BenchmarkReadUint64MultiplyAdd16(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64MultiplyAdd, 16)
}
func BenchmarkReadUint64MultiplyAdd1024(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64MultiplyAdd, 1024)
}
func BenchmarkReadUint64MultiplyAdd65536(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64MultiplyAdd, 65536)
}
func BenchmarkReadUint64MultiplyAdd65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, mathrand.GlobalPRNG.ReadUint64MultiplyAdd, 65536)
}

func TestReadUint64AddRotate(t *testing.T) {
	//testRead(t, mathrand.GlobalPRNG.ReadUint64AddRotate)
}
func BenchmarkReadUint64AddRotate1(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64AddRotate, 1)
}
func BenchmarkReadUint64AddRotate16(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64AddRotate, 16)
}
func BenchmarkReadUint64AddRotate1024(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64AddRotate, 1024)
}
func BenchmarkReadUint64AddRotate65536(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64AddRotate, 65536)
}
func BenchmarkReadUint64AddRotate65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, mathrand.GlobalPRNG.ReadUint64AddRotate, 65536)
}

func TestReadUint64Xorshift(t *testing.T) {
	//testRead(t, mathrand.GlobalPRNG.ReadUint64Xorshift)
}
func BenchmarkReadUint64Xorshift1(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64Xorshift, 1)
}
func BenchmarkReadUint64Xorshift16(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64Xorshift, 16)
}
func BenchmarkReadUint64Xorshift1024(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64Xorshift, 1024)
}
func BenchmarkReadUint64Xorshift65536(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64Xorshift, 65536)
}
func BenchmarkReadUint64Xorshift65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, mathrand.GlobalPRNG.ReadUint64Xorshift, 65536)
}

func TestReadUint64Xoshiro256(t *testing.T) {
	//testRead(t, mathrand.GlobalPRNG.ReadUint64Xoshiro256)
}
func BenchmarkReadUint64Xoshiro2561(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64Xoshiro256, 1)
}
func BenchmarkReadUint64Xoshiro25616(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64Xoshiro256, 16)
}
func BenchmarkReadUint64Xoshiro2561024(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64Xoshiro256, 1024)
}
func BenchmarkReadUint64Xoshiro25665536(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64Xoshiro256, 65536)
}
func BenchmarkReadUint64Xoshiro25665536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, mathrand.GlobalPRNG.ReadUint64Xoshiro256, 65536)
}

func TestReadUint64LFG(t *testing.T) {
	//testRead(t, mathrand.GlobalPRNG.ReadUint64LFG)
}
func BenchmarkReadUint64LFG1(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64LFG, 1)
}
func BenchmarkReadUint64LFG16(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64LFG, 16)
}
func BenchmarkReadUint64LFG1024(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64LFG, 1024)
}
func BenchmarkReadUint64LFG65536(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64LFG, 65536)
}
func BenchmarkReadUint64LFG65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, mathrand.GlobalPRNG.ReadUint64LFG, 65536)
}

func TestReadUint64MSWS(t *testing.T) {
	//testRead(t, mathrand.GlobalPRNG.ReadUint64MSWS)
}
func BenchmarkReadUint64MSWS1(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64MSWS, 1)
}
func BenchmarkReadUint64MSWS16(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64MSWS, 16)
}
func BenchmarkReadUint64MSWS1024(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64MSWS, 1024)
}
func BenchmarkReadUint64MSWS65536(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint64MSWS, 65536)
}
func BenchmarkReadUint64MSWS65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, mathrand.GlobalPRNG.ReadUint64MSWS, 65536)
}

func TestReadUint32AddRotateMultiply(t *testing.T) {
	//testRead(t, mathrand.GlobalPRNG.ReadUint32AddRotateMultiply)
}
func BenchmarkReadUint32AddRotateMultiply1(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32AddRotateMultiply, 1)
}
func BenchmarkReadUint32AddRotateMultiply16(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32AddRotateMultiply, 16)
}
func BenchmarkReadUint32AddRotateMultiply1024(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32AddRotateMultiply, 1024)
}
func BenchmarkReadUint32AddRotateMultiply65536(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32AddRotateMultiply, 65536)
}
func BenchmarkReadUint32AddRotateMultiply65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, mathrand.GlobalPRNG.ReadUint32AddRotateMultiply, 65536)
}

func TestReadUint32MultiplyAdd(t *testing.T) {
	//testRead(t, mathrand.GlobalPRNG.ReadUint32MultiplyAdd)
}
func BenchmarkReadUint32MultiplyAdd1(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32MultiplyAdd, 1)
}
func BenchmarkReadUint32MultiplyAdd16(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32MultiplyAdd, 16)
}
func BenchmarkReadUint32MultiplyAdd1024(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32MultiplyAdd, 1024)
}
func BenchmarkReadUint32MultiplyAdd65536(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32MultiplyAdd, 65536)
}
func BenchmarkReadUint32MultiplyAdd65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, mathrand.GlobalPRNG.ReadUint32MultiplyAdd, 65536)
}

func TestReadUint32AddRotate(t *testing.T) {
	//testRead(t, mathrand.GlobalPRNG.ReadUint32AddRotate)
}
func BenchmarkReadUint32AddRotate1(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32AddRotate, 1)
}
func BenchmarkReadUint32AddRotate16(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32AddRotate, 16)
}
func BenchmarkReadUint32AddRotate1024(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32AddRotate, 1024)
}
func BenchmarkReadUint32AddRotate65536(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32AddRotate, 65536)
}
func BenchmarkReadUint32AddRotate65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, mathrand.GlobalPRNG.ReadUint32AddRotate, 65536)
}

func TestReadUint32Xorshift(t *testing.T) {
	//testRead(t, mathrand.GlobalPRNG.ReadUint32Xorshift)
}
func BenchmarkReadUint32Xorshift1(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32Xorshift, 1)
}
func BenchmarkReadUint32Xorshift16(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32Xorshift, 16)
}
func BenchmarkReadUint32Xorshift1024(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32Xorshift, 1024)
}
func BenchmarkReadUint32Xorshift65536(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32Xorshift, 65536)
}
func BenchmarkReadUint32Xorshift65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, mathrand.GlobalPRNG.ReadUint32Xorshift, 65536)
}

func TestReadUint32LFG(t *testing.T) {
	//testRead(t, mathrand.GlobalPRNG.ReadUint32LFG)
}
func BenchmarkReadUint32LFG1(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32LFG, 1)
}
func BenchmarkReadUint32LFG16(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32LFG, 16)
}
func BenchmarkReadUint32LFG1024(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32LFG, 1024)
}
func BenchmarkReadUint32LFG65536(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32LFG, 65536)
}
func BenchmarkReadUint32LFG65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, mathrand.GlobalPRNG.ReadUint32LFG, 65536)
}

func TestReadUint32PCG(t *testing.T) {
	//testRead(t, mathrand.GlobalPRNG.ReadUint32PCG)
}
func BenchmarkReadUint32PCG1(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32PCG, 1)
}
func BenchmarkReadUint32PCG16(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32PCG, 16)
}
func BenchmarkReadUint32PCG1024(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32PCG, 1024)
}
func BenchmarkReadUint32PCG65536(b *testing.B) {
	benchmarkRead(b, mathrand.GlobalPRNG.ReadUint32PCG, 65536)
}
func BenchmarkReadUint32PCG65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, mathrand.GlobalPRNG.ReadUint32PCG, 65536)
}
