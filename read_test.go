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

func TestRead(t *testing.T) {
	testRead(t, fastrand.Read)
}

func TestReadSafe(t *testing.T) {
	testRead(t, fastrand.ReadSafe)
}

func TestXORRead(t *testing.T) {
	testRead(t, fastrand.XORRead)
}

func TestXORReadSafe(t *testing.T) {
	testRead(t, fastrand.XORReadSafe)
}

func BenchmarkOurRead1(b *testing.B) {
	benchmarkRead(b, fastrand.Read, 1)
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

func BenchmarkOurReadFast1(b *testing.B) {
	benchmarkRead(b, fastrand.ReadFast, 1)
}
func BenchmarkOurReadFast16(b *testing.B) {
	benchmarkRead(b, fastrand.ReadFast, 16)
}
func BenchmarkOurReadFast1024(b *testing.B) {
	benchmarkRead(b, fastrand.ReadFast, 1024)
}
func BenchmarkOurReadFast65536(b *testing.B) {
	benchmarkRead(b, fastrand.ReadFast, 65536)
}
func BenchmarkOurReadFast16777216(b *testing.B) {
	benchmarkRead(b, fastrand.ReadFast, 16777216)
}

func BenchmarkOurReadFastSafe1(b *testing.B) {
	benchmarkRead(b, fastrand.ReadFastSafe, 1)
}
func BenchmarkOurReadFastSafe16(b *testing.B) {
	benchmarkRead(b, fastrand.ReadFastSafe, 16)
}
func BenchmarkOurReadFastSafe1024(b *testing.B) {
	benchmarkRead(b, fastrand.ReadFastSafe, 1024)
}
func BenchmarkOurReadFastSafe65536(b *testing.B) {
	benchmarkRead(b, fastrand.ReadFastSafe, 65536)
}
func BenchmarkOurReadFastSafe16777216(b *testing.B) {
	benchmarkRead(b, fastrand.ReadFastSafe, 16777216)
}

func BenchmarkOurXORRead16(b *testing.B) {
	benchmarkRead(b, fastrand.XORRead, 16)
}
func BenchmarkOurXORRead65536(b *testing.B) {
	benchmarkRead(b, fastrand.XORRead, 65536)
}
func BenchmarkOurXORReadSafe16(b *testing.B) {
	benchmarkRead(b, fastrand.XORReadSafe, 16)
}
func BenchmarkOurXORReadSafe65536(b *testing.B) {
	benchmarkRead(b, fastrand.XORReadSafe, 65536)
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

func BenchmarkOurReadSafe65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, fastrand.ReadSafe, 65536)
}
func BenchmarkOurReadFastSafe65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, fastrand.ReadFastSafe, 65536)
}
func BenchmarkOurXORReadSafe65536Concurrent(b *testing.B) {
	benchmarkConcurrentRead(b, fastrand.XORReadSafe, 65536)
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
