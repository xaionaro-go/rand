[![GoDoc](https://godoc.org/github.com/xaionaro-go/fastrand?status.svg)](https://godoc.org/github.com/xaionaro-go/fastrand)

This's an inaccurate fast PRNG (and does not have any cryptographic strength, of corse).

This PRNG is represented here as `BenchmarkOur*`.

`*Safe` functions are supposed to be used if it's required to avoid extra
repeats in a massively-concurrent case (for example if you generate
[nonce](https://en.wikipedia.org/wiki/Cryptographic_nonce) very-very often and
very concurrently).

```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/fastrand
BenchmarkOurRead1-8                          	684153694	         7.67 ns/op	 130.37 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead15-8                         	556264449	        10.9 ns/op	1376.97 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16-8                         	579314013	        11.1 ns/op	1443.36 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead1024-8                       	46705255	       123 ns/op	8335.06 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead65536-8                      	  765792	      7396 ns/op	8860.56 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16777216-8                   	    2086	   3112056 ns/op	5391.04 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe1-8                      	371129146	        16.1 ns/op	  62.00 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe15-8                     	357416206	        18.3 ns/op	 819.08 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe16-8                     	334207412	        19.0 ns/op	 842.07 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe1024-8                   	50777352	       128 ns/op	8013.77 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe65536-8                  	  753756	      7471 ns/op	8771.84 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe16777216-8               	    2056	   3189283 ns/op	5260.50 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1-8                     	199069086	        31.3 ns/op	  31.92 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead15-8                    	100000000	        54.3 ns/op	 276.48 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16-8                    	114304114	        51.5 ns/op	 310.80 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1024-8                  	 3341830	      1719 ns/op	 595.70 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead65536-8                 	   56276	    103849 ns/op	 631.07 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16777216-8              	     169	  33226410 ns/op	 504.94 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine1-8                     	143607873	        38.3 ns/op	  26.08 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine15-8                    	125727135	        51.1 ns/op	 293.63 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine16-8                    	133744252	        47.0 ns/op	 340.24 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine1024-8                  	11906641	       495 ns/op	2067.51 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine65536-8                 	  274404	     20262 ns/op	3234.37 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine16777216-8              	    1166	   5421360 ns/op	3094.65 MB/s	       1 B/op	       0 allocs/op
BenchmarkOurRead65536Concurrent-8            	 2171890	      2605 ns/op	25153.43 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampineRead65536Concurrent-8   	 1000000	      6551 ns/op	10004.66 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint32n-8                        	1000000000	         2.40 ns/op	1666.04 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint64n-8                        	1000000000	         2.32 ns/op	3449.73 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint32nSafe-8                    	956617116	         6.32 ns/op	 632.82 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint64nSafe-8                    	931174922	         6.57 ns/op	1218.15 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardIntn-8                      	201956032	        32.9 ns/op	 243.44 MB/s	       0 B/op	       0 allocs/op
BenchmarkValyalaUint32n-8                    	215115235	        27.5 ns/op	 145.52 MB/s	       0 B/op	       0 allocs/op
BenchmarkNebulousLabsIntn-8                  	30350242	       217 ns/op	  36.92 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampineUint64n-8               	93761186	        57.8 ns/op	 138.40 MB/s	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/fastrand	251.998s
```

Links to similar projects:

* [`*Valyala*`](https://github.com/valyala/fastrand)
* [`*NebulousLabs*`](https://gitlab.com/NebulousLabs/fastrand)
* [`*Lukechampine*`](https://lukechampine.com/frand)