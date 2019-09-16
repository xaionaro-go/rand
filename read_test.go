package fastrand_test

import (
	"math/rand"
	"testing"

	"github.com/xaionaro-go/fastrand"
)

func BenchmarkOurRead1(b *testing.B) {
	benchmarkRead(b, fastrand.Read, 1)
}

func BenchmarkOurRead15(b *testing.B) {
	benchmarkRead(b, fastrand.Read, 15)
}

func BenchmarkOurRead16(b *testing.B) {
	benchmarkRead(b, fastrand.Read, 16)
}

func BenchmarkOurRead1024(b *testing.B) {
	benchmarkRead(b, fastrand.Read, 1024)
}

func BenchmarkStandardRead1(b *testing.B) {
	benchmarkRead(b, rand.Read, 1)
}

func BenchmarkStandardRead15(b *testing.B) {
	benchmarkRead(b, rand.Read, 15)
}

func BenchmarkStandardRead16(b *testing.B) {
	benchmarkRead(b, rand.Read, 16)
}

func BenchmarkStandardRead1024(b *testing.B) {
	benchmarkRead(b, rand.Read, 1024)
}

func benchmarkRead(b *testing.B, readFunc func([]byte) (int, error), bufSize uint) {
	buf := make([]byte, bufSize)
	b.SetBytes(int64(bufSize))
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_, _ = readFunc(buf)
	}
}