package fastrand

import (
	"math/rand"
	"testing"

	valyala "github.com/valyala/fastrand"
	nebulousLabs "gitlab.com/NebulousLabs/fastrand"
)

func BenchmarkOurIntn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Intn(1000)
	}
}

func BenchmarkStandardIntn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(1000)
	}
}

func BenchmarkValyalaUintn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		valyala.Uint32n(1000)
	}
}

func BenchmarkNebulousLabsUintn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nebulousLabs.Intn(1000)
	}
}
