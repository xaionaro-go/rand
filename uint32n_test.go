package fastrand_test

import (
	"math"
	"math/rand"
	"testing"

	valyala "github.com/valyala/fastrand"
	nebulousLabs "gitlab.com/NebulousLabs/fastrand"

	"github.com/xaionaro-go/fastrand"
)

func BenchmarkOurUint32nMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fastrand.Uint32n(math.MaxUint32)
	}
}

func BenchmarkOurUint32n(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fastrand.Uint32n(1000)
	}
}

func BenchmarkOurUint64n(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fastrand.Uint64n(1000)
	}
}

func BenchmarkStandardIntn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(1000)
	}
}

func BenchmarkValyalaUint32n(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valyala.Uint32n(1000)
	}
}

func BenchmarkNebulousLabsIntn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nebulousLabs.Intn(1000)
	}
}
