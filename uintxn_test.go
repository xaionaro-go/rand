package fastrand_test

import (
	"math/rand"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	valyala "github.com/valyala/fastrand"
	nebulousLabs "gitlab.com/NebulousLabs/fastrand"
	lukechampine "lukechampine.com/frand"

	"github.com/xaionaro-go/fastrand"
)

func BenchmarkOurUint32n(b *testing.B) {
	b.SetBytes(4)
	for i := 0; i < b.N; i++ {
		fastrand.Uint32n(1000)
	}
}

func BenchmarkOurUint32nSafe(b *testing.B) {
	b.SetBytes(4)
	for i := 0; i < b.N; i++ {
		fastrand.Uint32nSafe(1000)
	}
}

func BenchmarkOurUint32nFast(b *testing.B) {
	b.SetBytes(4)
	for i := 0; i < b.N; i++ {
		fastrand.Uint32nFast(1000)
	}
}

func BenchmarkOurUint64n(b *testing.B) {
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		fastrand.Uint64n(1000)
	}
}

func BenchmarkOurUint64nSafe(b *testing.B) {
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		fastrand.Uint64nSafe(1000)
	}
}

func BenchmarkOurUint64nFast(b *testing.B) {
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		fastrand.Uint64nFast(1000)
	}
}

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

func testUint32n(t *testing.T, fn func(uint32) uint32) {
	var count int
	for i := 0; i < 1000; i++ {
		if fn(4) == 0 {
			count++
		}
	}
	assert.True(t, count > 220, count)
	assert.True(t, count < 280, count)

	for i := 0; i < 9000; i++ {
		if fn(4) == 0 {
			count++
		}
	}
	assert.True(t, count > 2300, count)
	assert.True(t, count < 2700, count)
}

func TestUint32n(t *testing.T) {
	testUint32n(t, fastrand.Uint32n)
}
func TestUint32nSafe(t *testing.T) {
	testUint32n(t, fastrand.Uint32nSafe)
}
func TestUint32nFast(t *testing.T) {
	testUint32n(t, fastrand.Uint32nFast)
}

func testUint64n(t *testing.T, fn func(uint64) uint64) {
	var count int
	for i := 0; i < 1000; i++ {
		if fn(4) == 0 {
			count++
		}
	}
	assert.True(t, count > 220, count)
	assert.True(t, count < 280, count)

	for i := 0; i < 9000; i++ {
		if fn(4) == 0 {
			count++
		}
	}
	assert.True(t, count > 2300, count)
	assert.True(t, count < 2700, count)
}

func TestUint64n(t *testing.T) {
	testUint64n(t, fastrand.Uint64n)
}
func TestUint64nSafe(t *testing.T) {
	testUint64n(t, fastrand.Uint64nSafe)
}
func TestUint64nFast(t *testing.T) {
	testUint64n(t, fastrand.Uint64nFast)
}
