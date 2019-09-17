package fastrand_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	lukechampine "lukechampine.com/frand"

	"github.com/xaionaro-go/fastrand"
)

func testRead(t *testing.T, readFunc func([]byte) (int, error)) {
	for i := 4; i <= 20; i++ {
		count := 1 << i

		b := make([]byte, count)
		_, _ = readFunc(b)

		m := map[uint8]uint64{}
		var sum uint64
		for _, v := range b {
			sum += uint64(v)
			m[v]++
		}

		avg := float64(sum) / float64(len(b))
		assert.True(t, avg > 128-(20*(1<<8)/float64(count))-float64(count), fmt.Sprintf("%v: %v", i, avg))
		assert.True(t, avg < 136+(20*(1<<8)/float64(count))+float64(count), fmt.Sprintf("%v: %v", i, avg))

		for v := 0; v < (1 << 8); v++ {
			c := m[uint8(v)]
			assert.True(t, float64(c)*(1+1/float64(i))+20 > float64(count/(1<<8))-(20*(1<<8)/float64(count)), fmt.Sprintf("%v: m[%v]==%v", i, v, c))
			assert.True(t, float64(c)*(1-1/float64(i))-20 < float64(count/(1<<8))+(20*(1<<8)/float64(count)), fmt.Sprintf("%v: m[%v]==%v", i, v, c))
		}
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

func BenchmarkLukechampine1(b *testing.B) {
	benchmarkRead(b, lukechampine.Read, 1)
}
func BenchmarkLukechampine15(b *testing.B) {
	benchmarkRead(b, lukechampine.Read, 15)
}
func BenchmarkLukechampine16(b *testing.B) {
	benchmarkRead(b, lukechampine.Read, 16)
}
func BenchmarkLukechampine1024(b *testing.B) {
	benchmarkRead(b, lukechampine.Read, 1024)
}
func BenchmarkLukechampine65536(b *testing.B) {
	benchmarkRead(b, lukechampine.Read, 65536)
}
func BenchmarkLukechampine16777216(b *testing.B) {
	benchmarkRead(b, lukechampine.Read, 16777216)
}

func BenchmarkOurRead65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, fastrand.Read, 65536)
}
func BenchmarkLukechampineRead65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, lukechampine.Read, 65536)
}

func benchmarkRead(b *testing.B, readFunc func([]byte) (int, error), bufSize uint) {
	buf := make([]byte, bufSize)
	b.SetBytes(int64(bufSize))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = readFunc(buf)
	}
}

func benchmarkConcurrentRead(b *testing.B, readFunc func([]byte) (int, error), bufSize uint) {
	bufPool := sync.Pool{New: func() interface{} { return &[][]byte{make([]byte, bufSize)}[0] }}
	b.SetBytes(int64(bufSize))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		buf := bufPool.Get().(*[]byte)
		for pb.Next() {
			_, _ = readFunc(*buf)
		}
		bufPool.Put(buf)
	})
}
