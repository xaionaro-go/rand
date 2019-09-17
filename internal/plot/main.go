package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"unsafe"

	"github.com/xaionaro-go/fastrand"
)

func main() {
	writeCsv(`/tmp/fastrand`, fastrand.Read)
	writeCsv(`/tmp/mathrand`, rand.Read)
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func writeCsv(filePath string, readFunc func([]byte) (int, error)) {
	buf := make([]byte, 1<<20)
	_, err := readFunc(buf)
	panicIfError(err)

	{
		m := map[uint8]uint64{}
		var sum uint64
		for _, v := range buf {
			sum += uint64(v)
			m[v]++
		}

		s := rand.Perm(256)
		sort.Slice(s, func(i, j int) bool {
			return m[uint8(s[i])] > m[uint8(s[j])]
		})

		f, err := os.Create(filePath + `-byte.csv`)
		panicIfError(err)
		for _, v := range s {
			_, err := fmt.Fprintf(f, "%v,%v\n", v, m[uint8(v)])
			panicIfError(err)
		}
	}

	{
		m := map[uint64]uint64{}
		for i := 0; i < len(buf); i++ {
			m[*(*uint64)((unsafe.Pointer)(&buf[i]))]++
		}

		f, err := os.Create(filePath + `-values.csv`)
		panicIfError(err)
		for v, c := range m {
			_, err := fmt.Fprintf(f, "%v,%v\n", v, c)
			panicIfError(err)
		}
	}
}
