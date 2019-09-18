package mathrand_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	lukechampine "lukechampine.com/frand"
)

const (
	allowedAvgError = 1 << 10
)

func testRead(t *testing.T, readFunc func([]byte) (int, error)) {
	errCount := 0
	for i := 6; i <= 20; i += 2 {
		count := 1 << i

		b := make([]byte, count)

		for j := 0; j < 21-i; j++ {
			_, _ = readFunc(b)

			m := map[uint8]uint64{}
			var sum uint64
			for _, v := range b {
				sum += uint64(v)
				m[v]++
			}

			avg := float64(sum) / float64(len(b))
			if !assert.True(t, avg > 128-allowedAvgError/float64(i*i), fmt.Sprintf("%v: %v", i, avg)) {
				errCount++
			}
			if !assert.True(t, avg < 128+allowedAvgError/float64(i*i), fmt.Sprintf("%v: %v", i, avg)) {
				errCount++
			}

			for v := 0; v < (1 << 8); v++ {
				c := m[uint8(v)]
				if !assert.True(t, float64(c)*1.15+50 > float64(count)/(1<<8), fmt.Sprintf("%v: m[%v]==%v far from %v", i, v, c, count/256)) {
					errCount++
				}
				if !assert.True(t, float64(c)*0.85-30 < float64(count)/(1<<8), fmt.Sprintf("%v: m[%v]==%v far from %v", i, v, c, count/256)) {
					errCount++
				}
			}
		}
	}
	if errCount > 0 {
		t.Fatalf("errCount == %v", errCount)
	}
}

func TestStandardRead(t *testing.T) {
	testRead(t, rand.Read)
}

func BenchmarkStandardRead1(b *testing.B) {
	benchmarkRead(b, rand.Read, 1)
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
func BenchmarkLukechampine1024(b *testing.B) {
	benchmarkRead(b, lukechampine.Read, 1024)
}
func BenchmarkLukechampine65536(b *testing.B) {
	benchmarkRead(b, lukechampine.Read, 65536)
}
func BenchmarkLukechampine16777216(b *testing.B) {
	benchmarkRead(b, lukechampine.Read, 16777216)
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
