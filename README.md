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
BenchmarkOurRead1-8               	345725725	         9.48 ns/op	 105.53 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead15-8              	356927011	         9.75 ns/op	1538.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16-8              	358622312	        10.2 ns/op	1562.52 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead1024-8            	15735012	       242 ns/op	4235.85 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead65536-8           	  246480	     15999 ns/op	4096.21 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16777216-8        	     822	   3963771 ns/op	4232.64 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe1-8           	186120896	        20.4 ns/op	  49.10 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe15-8          	184437913	        23.0 ns/op	 651.92 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe16-8          	203232429	        18.5 ns/op	 864.38 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe1024-8        	15831406	       254 ns/op	4031.47 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe65536-8       	  230686	     14295 ns/op	4584.51 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe16777216-8    	     931	   3947902 ns/op	4249.65 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1-8          	129588265	        26.5 ns/op	  37.67 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead15-8         	53796502	        63.3 ns/op	 236.95 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16-8         	61517721	        63.5 ns/op	 251.90 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1024-8       	 1955476	      1724 ns/op	 594.07 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead65536-8      	   34784	    118683 ns/op	 552.19 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16777216-8   	     129	  27449696 ns/op	 611.20 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint32n-8             	1000000000	         2.89 ns/op	1382.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint64n-8             	1000000000	         3.11 ns/op	2574.18 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint32nSafe-8         	563245276	         6.16 ns/op	 649.40 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint64nSafe-8         	611604748	         6.41 ns/op	1248.33 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardIntn-8           	133002574	        27.8 ns/op	 287.49 MB/s	       0 B/op	       0 allocs/op
BenchmarkValyalaUint32n-8         	100000000	        31.4 ns/op	 127.47 MB/s	       0 B/op	       0 allocs/op
BenchmarkNebulousLabsIntn-8       	16765630	       279 ns/op	  28.68 MB/s	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/fastrand	117.146s
```