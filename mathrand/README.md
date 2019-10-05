[![GoDoc](https://godoc.org/github.com/xaionaro-go/rand/mathrand?status.svg)](https://godoc.org/github.com/xaionaro-go/rand/mathrand)

A collection of popular fast PRNGs (non crypto). This implementations may be used as
analog of [`math/rand`](https://golang.org/pkg/math/rand/).

[Diehard tests](https://en.wikipedia.org/wiki/Diehard_tests) result:

|func|passed|weak|failed|score|
|----|------|----|------|-----|
|Uint32AddIfShiftXOR| |1|113|[0.5](./docs/dieharder-results/prngSample-Uint32AddIfShiftXOR.bin.txt)|
|Uint32AddIfShiftXORWithReseed|9|4|101|[11](./docs/dieharder-results/prngSample-Uint32AddIfShiftXORWithReseed.bin.txt)|
|Uint32AddRotate|1|1|112|[1.5](./docs/dieharder-results/prngSample-Uint32AddRotate.bin.txt)|
|Uint32AddRotateMultiply|9|1|104|[9.5](./docs/dieharder-results/prngSample-Uint32AddRotateMultiply.bin.txt)|
|Uint32AddRotateMultiplyWithReseed|24|17|73|[32.5](./docs/dieharder-results/prngSample-Uint32AddRotateMultiplyWithReseed.bin.txt)|
|Uint32AddRotateWithReseed|17|15|82|[24.5](./docs/dieharder-results/prngSample-Uint32AddRotateWithReseed.bin.txt)|
|[Uint32MultiplyAdd](https://en.wikipedia.org/wiki/Linear_congruential_generator)|7|3|104|[8.5](./docs/dieharder-results/prngSample-Uint32MultiplyAdd.bin.txt)|
|Uint32MultiplyAddWithReseed|17|26|71|[30](./docs/dieharder-results/prngSample-Uint32MultiplyAddWithReseed.bin.txt)|
|[Uint32PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator)|22|20|72|[32](./docs/dieharder-results/prngSample-Uint32PCG.bin.txt)|
|Uint32PCGWithReseed|18|22|74|[29](./docs/dieharder-results/prngSample-Uint32PCGWithReseed.bin.txt)|
|[Uint32Xorshift](https://en.wikipedia.org/wiki/Xorshift)|24|12|78|[30](./docs/dieharder-results/prngSample-Uint32Xorshift.bin.txt)|
|Uint32XorshiftWithReseed|29|14|71|[36](./docs/dieharder-results/prngSample-Uint32XorshiftWithReseed.bin.txt)|
|Uint64AddIfShiftXOR| |3|111|[1.5](./docs/dieharder-results/prngSample-Uint64AddIfShiftXOR.bin.txt)|
|Uint64AddIfShiftXORWithReseed|9|4|101|[11](./docs/dieharder-results/prngSample-Uint64AddIfShiftXORWithReseed.bin.txt)|
|Uint64AddNRotateMultiply|17|21|76|[27.5](./docs/dieharder-results/prngSample-Uint64AddNRotateMultiply.bin.txt)|
|Uint64AddNRotateMultiplyWithReseed|21|16|77|[29](./docs/dieharder-results/prngSample-Uint64AddNRotateMultiplyWithReseed.bin.txt)|
|Uint64AddRotate|1|1|112|[1.5](./docs/dieharder-results/prngSample-Uint64AddRotate.bin.txt)|
|Uint64AddRotateMultiply|18|13|83|[24.5](./docs/dieharder-results/prngSample-Uint64AddRotateMultiply.bin.txt)|
|Uint64AddRotateMultiplyWithReseed|20|18|76|[29](./docs/dieharder-results/prngSample-Uint64AddRotateMultiplyWithReseed.bin.txt)|
|Uint64AddRotateWithReseed|24|8|82|[28](./docs/dieharder-results/prngSample-Uint64AddRotateWithReseed.bin.txt)|
|[Uint64MSWS](https://en.wikipedia.org/wiki/Middle-square_method)|23|16|75|[31](./docs/dieharder-results/prngSample-Uint64MSWS.bin.txt)|
|Uint64MSWSWithReseed|16|20|78|[26](./docs/dieharder-results/prngSample-Uint64MSWSWithReseed.bin.txt)|
|[Uint64MultiplyAdd](https://en.wikipedia.org/wiki/Linear_congruential_generator)|11|11|92|[16.5](./docs/dieharder-results/prngSample-Uint64MultiplyAdd.bin.txt)|
|Uint64MultiplyAddWithReseed|20|15|79|[27.5](./docs/dieharder-results/prngSample-Uint64MultiplyAddWithReseed.bin.txt)|
|[Uint64Xorshift](https://en.wikipedia.org/wiki/Xorshift)|26|11|77|[31.5](./docs/dieharder-results/prngSample-Uint64Xorshift.bin.txt)|
|Uint64XorshiftWithReseed|19|16|79|[27](./docs/dieharder-results/prngSample-Uint64XorshiftWithReseed.bin.txt)|
|[Uint64Xoshiro256](https://en.wikipedia.org/wiki/Xorshift#xoshiro_and_xoroshiro)|36|10|68|[41](./docs/dieharder-results/prngSample-Uint64Xoshiro256.bin.txt)|
|Uint64Xoshiro256WithReseed|16|17|81|[24.5](./docs/dieharder-results/prngSample-Uint64Xoshiro256WithReseed.bin.txt)|
|XORUint32AddIfShiftXOR| |1|113|[0.5](./docs/dieharder-results/prngSample-XORUint32AddIfShiftXOR.bin.txt)|
|XORUint32AddIfShiftXORWithReseed|9|4|101|[11](./docs/dieharder-results/prngSample-XORUint32AddIfShiftXORWithReseed.bin.txt)|
|XORUint32AddRotate|1|1|112|[1.5](./docs/dieharder-results/prngSample-XORUint32AddRotate.bin.txt)|
|XORUint32AddRotateMultiply|9|1|104|[9.5](./docs/dieharder-results/prngSample-XORUint32AddRotateMultiply.bin.txt)|
|XORUint32AddRotateMultiplyWithReseed|24|17|73|[32.5](./docs/dieharder-results/prngSample-XORUint32AddRotateMultiplyWithReseed.bin.txt)|
|XORUint32AddRotateWithReseed|17|15|82|[24.5](./docs/dieharder-results/prngSample-XORUint32AddRotateWithReseed.bin.txt)|
|XORUint32MultiplyAdd|8|2|104|[9](./docs/dieharder-results/prngSample-XORUint32MultiplyAdd.bin.txt)|
|XORUint32MultiplyAddWithReseed|17|26|71|[30](./docs/dieharder-results/prngSample-XORUint32MultiplyAddWithReseed.bin.txt)|
|XORUint32PCG|26|15|73|[33.5](./docs/dieharder-results/prngSample-XORUint32PCG.bin.txt)|
|XORUint32PCGWithReseed|18|22|74|[29](./docs/dieharder-results/prngSample-XORUint32PCGWithReseed.bin.txt)|
|XORUint32Xorshift|25|11|78|[30.5](./docs/dieharder-results/prngSample-XORUint32Xorshift.bin.txt)|
|XORUint32XorshiftWithReseed|29|14|71|[36](./docs/dieharder-results/prngSample-XORUint32XorshiftWithReseed.bin.txt)|
|XORUint64AddIfShiftXOR| |3|111|[1.5](./docs/dieharder-results/prngSample-XORUint64AddIfShiftXOR.bin.txt)|
|XORUint64AddIfShiftXORWithReseed|9|4|101|[11](./docs/dieharder-results/prngSample-XORUint64AddIfShiftXORWithReseed.bin.txt)|
|XORUint64AddNRotateMultiply|17|21|76|[27.5](./docs/dieharder-results/prngSample-XORUint64AddNRotateMultiply.bin.txt)|
|XORUint64AddNRotateMultiplyWithReseed|21|16|77|[29](./docs/dieharder-results/prngSample-XORUint64AddNRotateMultiplyWithReseed.bin.txt)|
|XORUint64AddRotate|1|1|112|[1.5](./docs/dieharder-results/prngSample-XORUint64AddRotate.bin.txt)|
|XORUint64AddRotateMultiply|19|10|85|[24](./docs/dieharder-results/prngSample-XORUint64AddRotateMultiply.bin.txt)|
|XORUint64AddRotateMultiplyWithReseed|20|18|76|[29](./docs/dieharder-results/prngSample-XORUint64AddRotateMultiplyWithReseed.bin.txt)|
|XORUint64AddRotateWithReseed|24|8|82|[28](./docs/dieharder-results/prngSample-XORUint64AddRotateWithReseed.bin.txt)|
|XORUint64MSWS|18|10|86|[23](./docs/dieharder-results/prngSample-XORUint64MSWS.bin.txt)|
|XORUint64MSWSWithReseed|16|20|78|[26](./docs/dieharder-results/prngSample-XORUint64MSWSWithReseed.bin.txt)|
|XORUint64MultiplyAdd|11|11|92|[16.5](./docs/dieharder-results/prngSample-XORUint64MultiplyAdd.bin.txt)|
|XORUint64MultiplyAddWithReseed|20|15|79|[27.5](./docs/dieharder-results/prngSample-XORUint64MultiplyAddWithReseed.bin.txt)|
|XORUint64Xorshift|26|11|77|[31.5](./docs/dieharder-results/prngSample-XORUint64Xorshift.bin.txt)|
|XORUint64XorshiftWithReseed|19|16|79|[27](./docs/dieharder-results/prngSample-XORUint64XorshiftWithReseed.bin.txt)|
|XORUint64Xoshiro256|20|19|75|[29.5](./docs/dieharder-results/prngSample-XORUint64Xoshiro256.bin.txt)|
|XORUint64Xoshiro256WithReseed|16|17|81|[24.5](./docs/dieharder-results/prngSample-XORUint64Xoshiro256WithReseed.bin.txt)|

If you need a fast PRNG then choose the fastest implementation which does not not fail on important-to-you tests.

Benchmarks:

```
BenchmarkReadUint64AddRotateMultiplyWithReseed65536Concurrent-12     	goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/rand/mathrand
BenchmarkReadUint64AddRotateMultiply1-12                             	1000000000	         5.91 ns/op	 169.28 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotateMultiply16-12                            	1000000000	         5.94 ns/op	2692.28 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotateMultiply1024-12                          	83157897	       144 ns/op	7096.43 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotateMultiply65536-12                         	 1319334	      9101 ns/op	7201.32 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotateMultiply65536Concurrent-12               	 8676526	      1344 ns/op	48746.93 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotateMultiplyWithReseed1-12                   	1000000000	         5.67 ns/op	 176.32 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotateMultiplyWithReseed16-12                  	1000000000	         5.81 ns/op	2753.05 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotateMultiplyWithReseed1024-12                	78509712	       152 ns/op	6739.85 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotateMultiplyWithReseed65536-12               	 1236812	      9887 ns/op	6628.26 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotateMultiplyWithReseed65536Concurrent-12     	 5922577	      1973 ns/op	33218.97 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddNRotateMultiply1-12                            	1000000000	         7.69 ns/op	 130.08 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddNRotateMultiply16-12                           	1000000000	         8.79 ns/op	1819.23 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddNRotateMultiply1024-12                         	30040552	       393 ns/op	2602.58 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddNRotateMultiply65536-12                        	  478203	     24378 ns/op	2688.32 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddNRotateMultiply65536Concurrent-12              	 3086086	      3743 ns/op	17507.67 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddNRotateMultiplyWithReseed1-12                  	1000000000	         7.34 ns/op	 136.22 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddNRotateMultiplyWithReseed16-12                 	1000000000	         9.17 ns/op	1743.97 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddNRotateMultiplyWithReseed1024-12               	29574066	       406 ns/op	2519.95 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddNRotateMultiplyWithReseed65536-12              	  467214	     25746 ns/op	2545.47 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddNRotateMultiplyWithReseed65536Concurrent-12    	 2858484	      4047 ns/op	16193.68 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MultiplyAdd1-12                                   	1000000000	         5.51 ns/op	 181.65 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MultiplyAdd16-12                                  	1000000000	         5.37 ns/op	2979.11 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MultiplyAdd1024-12                                	100000000	       114 ns/op	8963.90 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MultiplyAdd65536-12                               	 1639879	      7346 ns/op	8921.67 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MultiplyAdd65536Concurrent-12                     	 8702646	      1346 ns/op	48698.28 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MultiplyAddWithReseed1-12                         	1000000000	         5.34 ns/op	 187.42 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MultiplyAddWithReseed16-12                        	1000000000	         5.51 ns/op	2903.46 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MultiplyAddWithReseed1024-12                      	96030442	       124 ns/op	8243.72 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MultiplyAddWithReseed65536-12                     	 1509570	      7828 ns/op	8372.51 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MultiplyAddWithReseed65536Concurrent-12           	 8741253	      1344 ns/op	48747.50 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotate1-12                                     	1000000000	         5.63 ns/op	 177.61 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotate16-12                                    	1000000000	         5.42 ns/op	2952.00 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotate1024-12                                  	182573494	        67.5 ns/op	15159.29 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotate65536-12                                 	 3237300	      3904 ns/op	16787.75 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotate65536Concurrent-12                       	12520366	       895 ns/op	73215.25 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotateWithReseed1-12                           	1000000000	         5.52 ns/op	 181.08 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotateWithReseed16-12                          	1000000000	         5.89 ns/op	2717.95 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotateWithReseed1024-12                        	100000000	       109 ns/op	9368.62 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotateWithReseed65536-12                       	 1971318	      5898 ns/op	11111.25 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddRotateWithReseed65536Concurrent-12             	 8502072	      1377 ns/op	47593.25 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddIfShiftXOR1-12                                 	1000000000	         6.84 ns/op	 146.21 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddIfShiftXOR16-12                                	1000000000	         5.75 ns/op	2783.81 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddIfShiftXOR1024-12                              	96096104	       118 ns/op	8650.74 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddIfShiftXOR65536-12                             	 1580743	      7537 ns/op	8694.85 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddIfShiftXOR65536Concurrent-12                   	 8580079	      1357 ns/op	48300.82 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddIfShiftXORWithReseed1-12                       	1000000000	         6.85 ns/op	 145.93 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddIfShiftXORWithReseed16-12                      	1000000000	         6.51 ns/op	2457.85 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddIfShiftXORWithReseed1024-12                    	93999710	       127 ns/op	8060.02 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddIfShiftXORWithReseed65536-12                   	 1491421	      7904 ns/op	8291.92 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64AddIfShiftXORWithReseed65536Concurrent-12         	 7076581	      1643 ns/op	39884.89 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64Xorshift1-12                                      	1000000000	         5.60 ns/op	 178.64 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64Xorshift16-12                                     	1000000000	         6.03 ns/op	2653.13 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64Xorshift1024-12                                   	65962977	       178 ns/op	5755.88 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64Xorshift65536-12                                  	 1000000	     11162 ns/op	5871.43 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64Xorshift65536Concurrent-12                        	 8347202	      1406 ns/op	46619.89 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64XorshiftWithReseed1-12                            	1000000000	         5.59 ns/op	 178.74 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64XorshiftWithReseed16-12                           	1000000000	         6.32 ns/op	2532.31 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64XorshiftWithReseed1024-12                         	64752655	       184 ns/op	5556.25 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64XorshiftWithReseed65536-12                        	 1000000	     11752 ns/op	5576.69 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64XorshiftWithReseed65536Concurrent-12              	 6884772	      1707 ns/op	38402.60 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64Xoshiro2561-12                                    	1000000000	         8.77 ns/op	 113.99 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64Xoshiro25616-12                                   	1000000000	         7.19 ns/op	2225.54 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64Xoshiro2561024-12                                 	86924043	       135 ns/op	7585.89 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64Xoshiro25665536-12                                	 1419520	      8245 ns/op	7948.93 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64Xoshiro25665536Concurrent-12                      	 5786520	      2001 ns/op	32754.42 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64Xoshiro256WithReseed1-12                          	1000000000	         9.06 ns/op	 110.36 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64Xoshiro256WithReseed16-12                         	1000000000	         7.57 ns/op	2114.84 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64Xoshiro256WithReseed1024-12                       	74240482	       162 ns/op	6338.42 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64Xoshiro256WithReseed65536-12                      	 1242190	      9781 ns/op	6700.03 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64Xoshiro256WithReseed65536Concurrent-12            	 4955490	      2349 ns/op	27900.09 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MSWS1-12                                          	1000000000	         6.16 ns/op	 162.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MSWS16-12                                         	1000000000	         5.74 ns/op	2785.31 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MSWS1024-12                                       	68370486	       175 ns/op	5846.41 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MSWS65536-12                                      	 1000000	     11141 ns/op	5882.29 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MSWS65536Concurrent-12                            	 8380882	      1411 ns/op	46448.32 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MSWSWithReseed1-12                                	1000000000	         5.94 ns/op	 168.49 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MSWSWithReseed16-12                               	1000000000	         5.92 ns/op	2702.21 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MSWSWithReseed1024-12                             	63563898	       179 ns/op	5709.31 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MSWSWithReseed65536-12                            	 1000000	     11522 ns/op	5688.06 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint64MSWSWithReseed65536Concurrent-12                  	 7868062	      1495 ns/op	43825.05 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotateMultiply1-12                             	1000000000	         5.32 ns/op	 187.81 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotateMultiply16-12                            	1000000000	         5.81 ns/op	2752.89 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotateMultiply1024-12                          	50419069	       232 ns/op	4406.42 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotateMultiply65536-12                         	  803712	     14584 ns/op	4493.60 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotateMultiply65536Concurrent-12               	 6675978	      1776 ns/op	36909.88 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotateMultiplyWithReseed1-12                   	1000000000	         5.38 ns/op	 185.81 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotateMultiplyWithReseed16-12                  	1000000000	         6.22 ns/op	2570.96 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotateMultiplyWithReseed1024-12                	50033872	       240 ns/op	4266.25 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotateMultiplyWithReseed65536-12               	  769466	     15656 ns/op	4185.96 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotateMultiplyWithReseed65536Concurrent-12     	 6265228	      1862 ns/op	35190.64 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32MultiplyAdd1-12                                   	1000000000	         5.12 ns/op	 195.33 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32MultiplyAdd16-12                                  	1000000000	         6.20 ns/op	2580.01 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32MultiplyAdd1024-12                                	51695251	       230 ns/op	4459.82 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32MultiplyAdd65536-12                               	  809083	     14668 ns/op	4468.03 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32MultiplyAdd65536Concurrent-12                     	 6688881	      1745 ns/op	37553.21 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32MultiplyAddWithReseed1-12                         	1000000000	         5.19 ns/op	 192.58 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32MultiplyAddWithReseed16-12                        	1000000000	         6.51 ns/op	2457.16 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32MultiplyAddWithReseed1024-12                      	49158019	       240 ns/op	4275.36 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32MultiplyAddWithReseed65536-12                     	  784038	     14898 ns/op	4398.93 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32MultiplyAddWithReseed65536Concurrent-12           	 6654824	      1775 ns/op	36927.91 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotate1-12                                     	1000000000	         5.58 ns/op	 179.24 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotate16-12                                    	1000000000	         6.17 ns/op	2591.94 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotate1024-12                                  	146871367	        82.6 ns/op	12395.59 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotate65536-12                                 	 2419917	      4974 ns/op	13176.65 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotate65536Concurrent-12                       	10127445	      1147 ns/op	57153.43 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotateWithReseed1-12                           	1000000000	         5.31 ns/op	 188.28 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotateWithReseed16-12                          	1000000000	         6.22 ns/op	2573.78 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotateWithReseed1024-12                        	100000000	       118 ns/op	8674.39 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotateWithReseed65536-12                       	 1759989	      6783 ns/op	9662.51 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddRotateWithReseed65536Concurrent-12             	 6268934	      1703 ns/op	38474.10 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddIfShiftXOR1-12                                 	1000000000	         6.77 ns/op	 147.70 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddIfShiftXOR16-12                                	1000000000	         6.78 ns/op	2359.86 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddIfShiftXOR1024-12                              	51157076	       234 ns/op	4378.01 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddIfShiftXOR65536-12                             	  799681	     14744 ns/op	4444.93 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddIfShiftXOR65536Concurrent-12                   	 4693885	      2540 ns/op	25799.37 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddIfShiftXORWithReseed1-12                       	1000000000	         6.78 ns/op	 147.50 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddIfShiftXORWithReseed16-12                      	1000000000	         7.80 ns/op	2050.32 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddIfShiftXORWithReseed1024-12                    	47704569	       247 ns/op	4153.08 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddIfShiftXORWithReseed65536-12                   	  771753	     15423 ns/op	4249.19 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32AddIfShiftXORWithReseed65536Concurrent-12         	 3737464	      3095 ns/op	21173.48 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32Xorshift1-12                                      	1000000000	         5.57 ns/op	 179.44 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32Xorshift16-12                                     	1000000000	         7.44 ns/op	2150.19 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32Xorshift1024-12                                   	33225739	       352 ns/op	2910.57 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32Xorshift65536-12                                  	  537823	     22274 ns/op	2942.30 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32Xorshift65536Concurrent-12                        	 4218126	      2778 ns/op	23588.19 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32XorshiftWithReseed1-12                            	1000000000	         5.29 ns/op	 188.91 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32XorshiftWithReseed16-12                           	1000000000	         7.89 ns/op	2028.97 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32XorshiftWithReseed1024-12                         	32440348	       368 ns/op	2781.00 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32XorshiftWithReseed65536-12                        	  518901	     23133 ns/op	2833.01 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32XorshiftWithReseed65536Concurrent-12              	 3472923	      3502 ns/op	18713.59 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32PCG1-12                                           	1000000000	         6.85 ns/op	 145.96 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32PCG16-12                                          	941233701	        13.3 ns/op	1204.04 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32PCG1024-12                                        	25314841	       479 ns/op	2139.24 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32PCG65536-12                                       	  403969	     30289 ns/op	2163.68 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32PCG65536Concurrent-12                             	 1735347	      6712 ns/op	9764.48 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32PCGWithReseed1-12                                 	1000000000	         6.68 ns/op	 149.59 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32PCGWithReseed16-12                                	877823248	        13.4 ns/op	1193.18 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32PCGWithReseed1024-12                              	22678197	       551 ns/op	1860.05 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32PCGWithReseed65536-12                             	  357106	     34185 ns/op	1917.10 MB/s	       0 B/op	       0 allocs/op
BenchmarkReadUint32PCGWithReseed65536Concurrent-12                   	 1471394	      7832 ns/op	8367.92 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1-12                                            	620167975	        19.0 ns/op	  52.51 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16-12                                           	377367787	        30.9 ns/op	 517.92 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1024-12                                         	11654818	      1034 ns/op	 990.69 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead65536-12                                        	  183195	     65467 ns/op	1001.06 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16777216-12                                     	     698	  16642155 ns/op	1008.12 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine1-12                                            	461865512	        23.5 ns/op	  42.61 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine1024-12                                         	36983640	       322 ns/op	3184.26 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine65536-12                                        	  933505	     12436 ns/op	5269.75 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine16777216-12                                     	    3508	   3393452 ns/op	4944.00 MB/s	       0 B/op	       0 allocs/op
BenchmarkUint32AddRotateMultiply-12                                  	1000000000	         1.76 ns/op	2267.26 MB/s	       0 B/op	       0 allocs/op
BenchmarkUint32MultiplyAdd-12                                        	1000000000	         1.86 ns/op	2150.87 MB/s	       0 B/op	       0 allocs/op
BenchmarkUint32AddRotate-12                                          	1000000000	         1.60 ns/op	2500.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkUint32AddIfShiftXOR-12                                      	1000000000	         2.29 ns/op	1747.12 MB/s	       0 B/op	       0 allocs/op
BenchmarkUint32Xorshift-12                                           	1000000000	         2.05 ns/op	1950.09 MB/s	       0 B/op	       0 allocs/op
BenchmarkUint32PCG-12                                                	1000000000	         1.61 ns/op	2491.62 MB/s	       0 B/op	       0 allocs/op
BenchmarkUint64AddRotateMultiply-12                                  	1000000000	         1.95 ns/op	4105.49 MB/s	       0 B/op	       0 allocs/op
BenchmarkUint64AddNRotateMultiply-12                                 	1000000000	         3.70 ns/op	2163.08 MB/s	       0 B/op	       0 allocs/op
BenchmarkUint64MultiplyAdd-12                                        	1000000000	         1.87 ns/op	4267.77 MB/s	       0 B/op	       0 allocs/op
BenchmarkUint64AddRotate-12                                          	1000000000	         1.57 ns/op	5110.60 MB/s	       0 B/op	       0 allocs/op
BenchmarkUint64AddIfShiftXOR-12                                      	1000000000	         2.07 ns/op	3856.38 MB/s	       0 B/op	       0 allocs/op
BenchmarkUint64Xorshift-12                                           	1000000000	         2.11 ns/op	3783.07 MB/s	       0 B/op	       0 allocs/op
BenchmarkUint64Xoshiro256-12                                         	1000000000	         3.41 ns/op	2342.74 MB/s	       0 B/op	       0 allocs/op
BenchmarkUint64MSWS-12                                               	1000000000	         2.08 ns/op	3850.44 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardIntn-12                                             	688455276	        16.7 ns/op	 479.42 MB/s	       0 B/op	       0 allocs/op
BenchmarkValyalaUint32n-12                                           	755310522	        16.2 ns/op	 247.60 MB/s	       0 B/op	       0 allocs/op
BenchmarkValyalaUint32n_withPreparedRNG-12                           	1000000000	         2.23 ns/op	1793.55 MB/s	       0 B/op	       0 allocs/op
BenchmarkNebulousLabsIntn-12                                         	79827997	       150 ns/op	  53.42 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampineUint64n-12                                      	319181235	        36.5 ns/op	 218.92 MB/s	       0 B/op	       0 allocs/op
BenchmarkTimeNowUnixNano-12                                          	336841614	        35.6 ns/op	 224.61 MB/s	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/rand/mathrand	1801.846s
```


Other fast PRNGs implementations:

* [Valyala](https://github.com/valyala/fastrand)
* [NebulousLabs](https://gitlab.com/NebulousLabs/fastrand)
* [Lukechampine](https://lukechampine.com/frand)                	