package fastrand_test

import (
	"math/rand"
	"testing"
	"unsafe"

	valyala "github.com/valyala/fastrand"
	nebulousLabs "gitlab.com/NebulousLabs/fastrand"

	"github.com/xaionaro-go/fastrand"
)

func BenchmarkOurUint32n(b *testing.B) {
	b.SetBytes(4)
	for i := 0; i < b.N; i++ {
		fastrand.Uint32n(1000)
	}
}

func BenchmarkOurUint64n(b *testing.B) {
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		fastrand.Uint64n(1000)
	}
}

func BenchmarkOurUint32nSafe(b *testing.B) {
	b.SetBytes(4)
	for i := 0; i < b.N; i++ {
		fastrand.Uint32nSafe(1000)
	}
}

func BenchmarkOurUint64nSafe(b *testing.B) {
	b.SetBytes(8)
	for i := 0; i < b.N; i++ {
		fastrand.Uint64nSafe(1000)
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

func BenchmarkNebulousLabsIntn(b *testing.B) {
	b.SetBytes(int64(unsafe.Sizeof(int(0))))
	for i := 0; i < b.N; i++ {
		nebulousLabs.Intn(1000)
	}
}
