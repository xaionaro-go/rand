[![GoDoc](https://godoc.org/github.com/xaionaro-go/rand/mathrand?status.svg)](https://godoc.org/github.com/xaionaro-go/fastrand)

A collection of popular fast PRNGs (non crypto). This implementations may be used as
analog of [`math/rand`](https://golang.org/pkg/math/rand/).

[Diehard tests](https://en.wikipedia.org/wiki/Diehard_tests) result:

|func|passed|weak|failed|score|
|----|------|----|------|-----|
|[Uint32MultiplyAdd](https://en.wikipedia.org/wiki/Linear_congruential_generator)|7|3|104|[8.5](./docs/dieharder-results/prngSample-Uint32MultiplyAdd.bin.txt)|
|[Uint32PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator)|22|20|72|[32](./docs/dieharder-results/prngSample-Uint32PCG.bin.txt)|
|[Uint32Xorshift](https://en.wikipedia.org/wiki/Xorshift)|24|12|78|[30](./docs/dieharder-results/prngSample-Uint32Xorshift.bin.txt)|
|Uint64AddNRotateMultiply|17|21|76|[27.5](./docs/dieharder-results/prngSample-Uint64AddNRotateMultiply.bin.txt)|
|Uint64AddRotate|1|1|112|[1.5](./docs/dieharder-results/prngSample-Uint64AddRotate.bin.txt)|
|Uint64AddRotateMultiply|18|13|83|[24.5](./docs/dieharder-results/prngSample-Uint64AddRotateMultiply.bin.txt)|
|[Uint64MSWS](https://en.wikipedia.org/wiki/Middle-square_method)|23|16|75|[31](./docs/dieharder-results/prngSample-Uint64MSWS.bin.txt)|
|[Uint64MultiplyAdd](https://en.wikipedia.org/wiki/Linear_congruential_generator)|11|11|92|[16.5](./docs/dieharder-results/prngSample-Uint64MultiplyAdd.bin.txt)|
|[Uint64Xorshift](https://en.wikipedia.org/wiki/Xorshift)|26|11|77|[31.5](./docs/dieharder-results/prngSample-Uint64Xorshift.bin.txt)|
|[Uint64Xoshiro256](https://en.wikipedia.org/wiki/Xorshift#xoshiro_and_xoroshiro)|36|10|68|[41](./docs/dieharder-results/prngSample-Uint64Xoshiro256.bin.txt)|

Other fast PRNGs implementations:

* [`*Valyala*`](https://github.com/valyala/fastrand)
* [`*NebulousLabs*`](https://gitlab.com/NebulousLabs/fastrand)
* [`*Lukechampine*`](https://lukechampine.com/frand)
