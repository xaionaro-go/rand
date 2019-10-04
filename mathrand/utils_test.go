package mathrand_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func prepareSample(sampleName string, readFunc func([]byte) (int, error)) {
	b := make([]byte, 1<<20)
	_, err := readFunc(b)
	panicIfError(err)
	err = ioutil.WriteFile(filepath.Join(os.TempDir(), "prngSample-"+sampleName+".bin"), b, 0755)
	panicIfError(err)
}
