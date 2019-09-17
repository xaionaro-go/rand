[![GoDoc](https://godoc.org/github.com/xaionaro-go/fastrand?status.svg)](https://godoc.org/github.com/xaionaro-go/fastrand)

This's an inaccurate fast PRNG (which is *not* appropriate for cryptographic tasks).

This PRNG is represented here as `BenchmarkOur*`.

`*Safe` functions are supposed to be used if it's required to avoid extra
repeats in a massively-concurrent case (for example if you generate
[nonce](https://en.wikipedia.org/wiki/Cryptographic_nonce) very-very often and
very concurrently).

```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/fastrand
BenchmarkOurRead1-8                          	133874481	         9.31 ns/op	 107.40 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead15-8                         	100000000	        11.5 ns/op	1305.13 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16-8                         	92312860	        10.9 ns/op	1465.32 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead1024-8                       	 6751867	       165 ns/op	6199.97 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead65536-8                      	  104180	     12392 ns/op	5288.77 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16777216-8                   	     351	   2987566 ns/op	5615.68 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe1-8                      	76636980	        19.9 ns/op	  50.26 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe15-8                     	63476184	        21.6 ns/op	 694.36 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe16-8                     	66076305	        17.9 ns/op	 893.05 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe1024-8                   	 5667475	       195 ns/op	5243.92 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe65536-8                  	  116233	     10377 ns/op	6315.26 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe16777216-8               	     409	   2907537 ns/op	5770.25 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1-8                     	34547760	        29.3 ns/op	  34.19 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead15-8                    	21637260	        62.5 ns/op	 240.01 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16-8                    	19601065	        62.3 ns/op	 256.82 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1024-8                  	  602856	      1834 ns/op	 558.28 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead65536-8                 	   11546	    115369 ns/op	 568.06 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16777216-8              	      45	  26065307 ns/op	 643.66 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine1-8                     	29040774	        37.8 ns/op	  26.46 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine15-8                    	28462902	        45.7 ns/op	 328.38 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine16-8                    	27898988	        45.0 ns/op	 355.85 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine1024-8                  	 2465785	       474 ns/op	2158.30 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine65536-8                 	   60693	     18640 ns/op	3515.91 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine16777216-8              	     236	   5117649 ns/op	3278.31 MB/s	       9 B/op	       0 allocs/op
BenchmarkOurRead65536Concurrent-8            	  397102	      3011 ns/op	21762.81 MB/s	       1 B/op	       0 allocs/op
BenchmarkLukechampineRead65536Concurrent-8   	  202239	      6165 ns/op	10630.87 MB/s	       2 B/op	       0 allocs/op
BenchmarkOurUint32n-8                        	267784942	         3.92 ns/op	1019.74 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint64n-8                        	426234124	         2.97 ns/op	2695.64 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint32nSafe-8                    	183023590	         6.26 ns/op	 639.23 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint64nSafe-8                    	196154478	         6.40 ns/op	1250.36 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardIntn-8                      	35115006	        29.3 ns/op	 272.93 MB/s	       0 B/op	       0 allocs/op
BenchmarkValyalaUint32n-8                    	37192550	        28.2 ns/op	 141.61 MB/s	       0 B/op	       0 allocs/op
BenchmarkNebulousLabsIntn-8                  	 4744070	       222 ns/op	  36.08 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampineUint64n-8               	18547360	        54.0 ns/op	 148.02 MB/s	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/fastrand	51.408s
```

Links to similar projects:

* [`*Valyala*`](github.com/valyala/fastrand)
* [`*NebulousLabs*`](gitlab.com/NebulousLabs/fastrand)
* [`*Lukechampine*`](lukechampine.com/frand)
