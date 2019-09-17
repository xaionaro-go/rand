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
BenchmarkOurRead1-8                          	689462806	         9.14 ns/op	 109.38 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead15-8                         	634215814	        10.5 ns/op	1430.99 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16-8                         	554489727	        11.2 ns/op	1423.94 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead1024-8                       	46232006	       142 ns/op	7201.26 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead65536-8                      	  774494	      8442 ns/op	7763.21 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16777216-8                   	    2071	   2977328 ns/op	5634.99 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe1-8                      	311763897	        18.0 ns/op	  55.47 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe15-8                     	280863398	        19.0 ns/op	 791.38 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe16-8                     	249089362	        21.0 ns/op	 761.87 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe1024-8                   	48204207	       142 ns/op	7215.68 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe65536-8                  	  734068	      7141 ns/op	9177.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe16777216-8               	    2078	   2900122 ns/op	5785.00 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1-8                     	215769259	        28.0 ns/op	  35.65 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead15-8                    	100000000	        51.7 ns/op	 290.16 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16-8                    	100000000	        50.2 ns/op	 318.74 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1024-8                  	 3458767	      1899 ns/op	 539.16 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead65536-8                 	   46071	    114612 ns/op	 571.81 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16777216-8              	     222	  27502402 ns/op	 610.03 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine1-8                     	152826290	        37.7 ns/op	  26.49 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine15-8                    	129546694	        46.9 ns/op	 319.90 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine16-8                    	126510373	        46.2 ns/op	 346.02 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine1024-8                  	12077936	       515 ns/op	1989.16 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine65536-8                 	  316239	     18653 ns/op	3513.36 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine16777216-8              	    1141	   5100014 ns/op	3289.64 MB/s	       1 B/op	       0 allocs/op
BenchmarkOurRead65536Concurrent-8            	 2291564	      2618 ns/op	25036.98 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampineRead65536Concurrent-8   	  968508	      5955 ns/op	11005.82 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint32n-8                        	1000000000	         3.31 ns/op	1209.84 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint64n-8                        	1000000000	         2.35 ns/op	3407.82 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint32nSafe-8                    	918330576	         6.18 ns/op	 647.25 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint64nSafe-8                    	934976372	         6.27 ns/op	1275.67 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardIntn-8                      	205415931	        26.5 ns/op	 302.38 MB/s	       0 B/op	       0 allocs/op
BenchmarkValyalaUint32n-8                    	249376179	        24.1 ns/op	 165.77 MB/s	       0 B/op	       0 allocs/op
BenchmarkNebulousLabsIntn-8                  	27060601	       220 ns/op	  36.31 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampineUint64n-8               	97116636	        59.7 ns/op	 133.94 MB/s	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/fastrand	239.277s
```

Links to similar projects:

* [`*Valyala*`](https://github.com/valyala/fastrand)
* [`*NebulousLabs*`](https://gitlab.com/NebulousLabs/fastrand)
* [`*Lukechampine*`](https://lukechampine.com/frand)
