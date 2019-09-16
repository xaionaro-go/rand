package fastrand_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xaionaro-go/fastrand"
)

func testRead(t *testing.T, readFunc func([]byte) (int, error)) {
	count := 1 << 20

	b := make([]byte, count)
	_, _ = readFunc(b)

	m := map[uint8]uint64{}
	var sum uint64
	for _, v := range b {
		sum += uint64(v)
		m[v]++
	}

	avg := float64(sum) / float64(len(b))
	assert.True(t, avg > 120)
	assert.True(t, avg < 136)

	for v:=0; v<(1<<8); v++ {
		c:=m[uint8(v)]
		assert.True(t, float64(c)*1.05 > float64(count / (1 << 8)), fmt.Sprintf("m[%v]==%v", v, c))
		assert.True(t, float64(c)*0.95 < float64(count / (1 << 8)), fmt.Sprintf("m[%v]==%v", v, c))
	}
}

func TestRead(t *testing.T) {
	testRead(t, fastrand.Read)
}

func TestReadSafe(t *testing.T) {
	testRead(t, fastrand.ReadSafe)
}

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

func BenchmarkOurRead65536(b *testing.B) {
	benchmarkRead(b, fastrand.Read, 65536)
}

func BenchmarkOurRead16777216(b *testing.B) {
	benchmarkRead(b, fastrand.Read, 16777216)
}


func BenchmarkOurReadSafe1(b *testing.B) {
	benchmarkRead(b, fastrand.ReadSafe, 1)
}

func BenchmarkOurReadSafe15(b *testing.B) {
	benchmarkRead(b, fastrand.ReadSafe, 15)
}

func BenchmarkOurReadSafe16(b *testing.B) {
	benchmarkRead(b, fastrand.ReadSafe, 16)
}

func BenchmarkOurReadSafe1024(b *testing.B) {
	benchmarkRead(b, fastrand.ReadSafe, 1024)
}

func BenchmarkOurReadSafe65536(b *testing.B) {
	benchmarkRead(b, fastrand.ReadSafe, 65536)
}

func BenchmarkOurReadSafe16777216(b *testing.B) {
	benchmarkRead(b, fastrand.ReadSafe, 16777216)
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

func BenchmarkStandardRead65536(b *testing.B) {
	benchmarkRead(b, rand.Read, 65536)
}

func BenchmarkStandardRead16777216(b *testing.B) {
	benchmarkRead(b, rand.Read, 16777216)
}

func benchmarkRead(b *testing.B, readFunc func([]byte) (int, error), bufSize uint) {
	buf := make([]byte, bufSize)
	b.SetBytes(int64(bufSize))
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_, _ = readFunc(buf)
	}
}