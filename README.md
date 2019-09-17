[![GoDoc](https://godoc.org/github.com/xaionaro-go/fastrand?status.svg)](https://godoc.org/github.com/xaionaro-go/fastrand)

This's an inaccurate fast PRNG (which is *not* appropriate for cryptographic tasks).

This PRNG is represented here as `BenchmarkOur*`.

`*Safe` functions are supposed to be used if it's required to avoid extra
repeats in a massively-concurrent case (for example if you generate
[nonce](https://en.wikipedia.org/wiki/Cryptographic_nonce) very-very often and
very concurrently).

```
goarch: amd64
pkg: github.com/xaionaro-go/fastrand
BenchmarkOurRead65536Concurrent-8   	  115795	     14780 ns/op	4434.09 MB/s	       4 B/op	       0 allocs/op
BenchmarkOurRead1-8                 	100000000	        10.6 ns/op	  94.65 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead15-8                	100000000	        15.5 ns/op	 967.01 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16-8                	100000000	        11.9 ns/op	1344.96 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead1024-8              	 4493490	       223 ns/op	4592.65 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead65536-8             	   69462	     16492 ns/op	3973.83 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16777216-8          	     228	   4460071 ns/op	3761.65 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe1-8             	60834021	        16.9 ns/op	  59.07 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe15-8            	61023537	        19.7 ns/op	 762.17 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe16-8            	56041033	        20.2 ns/op	 790.14 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe1024-8          	 4341667	       286 ns/op	3577.52 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe65536-8         	   67015	     17000 ns/op	3855.15 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe16777216-8      	     225	   5194070 ns/op	3230.07 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1-8            	35815052	        31.7 ns/op	  31.52 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead15-8           	21532710	        65.5 ns/op	 228.90 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16-8           	19137955	        73.8 ns/op	 216.79 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1024-8         	  563085	      1851 ns/op	 553.31 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead65536-8        	    8043	    165679 ns/op	 395.56 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16777216-8     	      45	  26171578 ns/op	 641.05 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampineRead1-8           	32648130	        36.9 ns/op	  27.10 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampineRead15-8          	23006984	        59.8 ns/op	 250.95 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampineRead16-8          	22161181	        45.5 ns/op	 351.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampineRead1024-8        	 2605623	       485 ns/op	2110.40 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampineRead65536-8       	   65628	     18583 ns/op	3526.73 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampineRead16777216-8    	     234	   5262785 ns/op	3187.90 MB/s	       9 B/op	       0 allocs/op
BenchmarkOurRead65536Concurrent-8         124290          9566 ns/op    6850.75 MB/s           4 B/op          0 allocs/op
BenchmarkLukechampineRead65536Concurrent-8 198040         5921 ns/op   11069.26 MB/s           2 B/op          0 allocs/op
BenchmarkOurUint32n-8               	417679042	         3.12 ns/op	1283.88 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint64n-8               	362613686	         2.84 ns/op	2812.21 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint32nSafe-8           	210415731	         6.04 ns/op	 661.75 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint64nSafe-8           	207220892	         6.03 ns/op	1326.21 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardIntn-8             	39605470	        26.0 ns/op	 307.62 MB/s	       0 B/op	       0 allocs/op
BenchmarkValyalaUint32n-8           	41430604	        24.8 ns/op	 161.45 MB/s	       0 B/op	       0 allocs/op
BenchmarkNebulousLabsIntn-8         	 5387990	       265 ns/op	  30.22 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampineUint64n-8      	17338686	        60.1 ns/op	 133.15 MB/s	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/fastrand	47.305s
```

Links to similar projects:

* [`*Valyala*`](github.com/valyala/fastrand)
* [`*NebulousLabs*`](gitlab.com/NebulousLabs/fastrand)
* [`*Lukechampine*`](lukechampine.com/frand)goos: linux