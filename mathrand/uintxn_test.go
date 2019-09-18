package mathrand_test

import (
	"math/rand"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
	valyala "github.com/valyala/fastrand"
	nebulousLabs "gitlab.com/NebulousLabs/fastrand"
	lukechampine "lukechampine.com/frand"

	"github.com/xaionaro-go/rand/mathrand"
)

func BenchmarkStandardIntn(b *testing.B) {
	b.SetBytes(int64(unsafe.Sizeof(int(0))))
	for i := 0; i < b.N; i++ {
		rand.Intn(1000)
	}
}

func BenchmarkValyalaUint32n(b *testing.B) {
	b.SetBytes(4)
	for i := 0; i < b.N; i++ {
		valyala.Uint32n(1000)
	}
}

func BenchmarkValyalaUint32n_withPreparedRNG(b *testing.B) {
	prng := &valyala.RNG{}
	b.SetBytes(4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prng.Uint32n(1000)
	}
}

func BenchmarkNebulousLabsIntn(b *testing.B) {
	b.SetBytes(int64(unsafe.Sizeof(int(0))))
	for i := 0; i < b.N; i++ {
		nebulousLabs.Intn(1000)
	}
}

func BenchmarkLukechampineUint64n(b *testing.B) {
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		lukechampine.Uint64n(1000)
	}
}

func BenchmarkTimeNowUnixNano(b *testing.B) {
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		time.Now().UnixNano()
	}
}

func testUint32(t *testing.T, fn func() uint32) {
	var count int
	for i := 0; i < 1000; i++ {
		if mathrand.ReduceUint32(fn(), 4) == 0 {
			count++
		}
	}
	assert.True(t, count > 220, count)
	assert.True(t, count < 280, count)

	for i := 0; i < 9000; i++ {
		if mathrand.ReduceUint32(fn(), 4) == 0 {
			count++
		}
	}
	assert.True(t, count > 2300, count)
	assert.True(t, count < 2700, count)
}

func testUint64(t *testing.T, fn func() uint64) {
	var count int
	for i := 0; i < 1000; i++ {
		if mathrand.ReduceUint64(fn(), 4) == 0 {
			count++
		}
	}
	assert.True(t, count > 220, count)
	assert.True(t, count < 280, count)

	for i := 0; i < 9000; i++ {
		if mathrand.ReduceUint64(fn(), 4) == 0 {
			count++
		}
	}
	assert.True(t, count > 2300, count)
	assert.True(t, count < 2700, count)
}
