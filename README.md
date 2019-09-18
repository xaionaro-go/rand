[![GoDoc](https://godoc.org/github.com/xaionaro-go/fastrand?status.svg)](https://godoc.org/github.com/xaionaro-go/fastrand)

This's an inaccurate fast PRNG (and does not have any cryptographic strength, of corse).

This PRNG is represented here as `BenchmarkOur*`.

`*Safe` functions are supposed to be used if it's required to avoid extra
repeats in a massively-concurrent case (for example if you generate
[nonce](https://en.wikipedia.org/wiki/Cryptographic_nonce) very-very often and
very concurrently). Also you may avoid repeats by using seperate instances of
`PRNG` on each concurrent goroutine.

```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/fastrand
BenchmarkOurRead1-8                          	615473011	        10.7 ns/op	  93.72 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16-8                         	481814466	        11.8 ns/op	1356.64 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead1024-8                       	21212814	       284 ns/op	3610.52 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead65536-8                      	  343382	     17520 ns/op	3740.71 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16777216-8                   	    1261	   4849586 ns/op	3459.52 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe1-8                      	314425584	        18.5 ns/op	  54.10 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe16-8                     	270144727	        22.3 ns/op	 716.98 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe1024-8                   	19699413	       293 ns/op	3497.85 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe65536-8                  	  342253	     17317 ns/op	3784.59 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadSafe16777216-8               	    1232	   4926582 ns/op	3405.45 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadFast1-8                      	621846406	         9.50 ns/op	 105.29 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadFast16-8                     	529537584	        11.4 ns/op	1405.69 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadFast1024-8                   	49048400	       118 ns/op	8652.62 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadFast65536-8                  	 1000000	      5971 ns/op	10975.25 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadFast16777216-8               	    1699	   3628539 ns/op	4623.68 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadFastSafe1-8                  	311999089	        19.4 ns/op	  51.55 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadFastSafe16-8                 	289733929	        20.6 ns/op	 776.93 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadFastSafe1024-8               	49307298	       125 ns/op	8161.79 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadFastSafe65536-8              	  978957	      5934 ns/op	11044.04 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadFastSafe16777216-8           	    1678	   3690796 ns/op	4545.69 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurXORRead16-8                      	500404272	        11.9 ns/op	1342.01 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurXORRead65536-8                   	  325017	     18000 ns/op	3640.94 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurXORReadSafe16-8                  	269237541	        22.1 ns/op	 725.52 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurXORReadSafe65536-8               	  324918	     18606 ns/op	3522.37 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1-8                     	200496430	        30.1 ns/op	  33.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16-8                    	100000000	        53.8 ns/op	 297.21 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1024-8                  	 2965345	      1926 ns/op	 531.54 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead65536-8                 	   42576	    120732 ns/op	 542.82 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16777216-8              	     199	  29064634 ns/op	 577.24 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine1-8                     	149717118	        40.3 ns/op	  24.84 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine1024-8                  	12100960	       508 ns/op	2016.94 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine65536-8                 	  300075	     19695 ns/op	3327.53 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampine16777216-8              	     991	   5567938 ns/op	3013.18 MB/s	       2 B/op	       0 allocs/op
BenchmarkOurReadSafe65536Concurrent-8        	 1606766	      3728 ns/op	17581.10 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurReadFastSafe65536Concurrent-8    	 3210757	      1872 ns/op	35004.05 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurXORReadSafe65536Concurrent-8     	 1283068	      4633 ns/op	14145.28 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampineRead65536Concurrent-8   	  959941	      6044 ns/op	10842.47 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint32n-8                        	1000000000	         3.28 ns/op	1218.44 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint32nSafe-8                    	919258447	         6.18 ns/op	 647.52 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint32nFast-8                    	1000000000	         2.39 ns/op	1671.16 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint64n-8                        	1000000000	         3.42 ns/op	2341.35 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint64nSafe-8                    	931240846	         6.47 ns/op	1236.80 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint64nFast-8                    	1000000000	         2.38 ns/op	3359.24 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardIntn-8                      	241809102	        23.9 ns/op	 334.24 MB/s	       0 B/op	       0 allocs/op
BenchmarkValyalaUint32n-8                    	223066089	        27.1 ns/op	 147.53 MB/s	       0 B/op	       0 allocs/op
BenchmarkValyalaUint32n_withPreparedRNG-8    	1000000000	         3.99 ns/op	1003.52 MB/s	       0 B/op	       0 allocs/op
BenchmarkNebulousLabsIntn-8                  	28250271	       222 ns/op	  35.96 MB/s	       0 B/op	       0 allocs/op
BenchmarkLukechampineUint64n-8               	100000000	        60.1 ns/op	 133.01 MB/s	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/fastrand	330.464s
?   	github.com/xaionaro-go/fastrand/internal/plot	[no test files]
```

Links to similar projects:

* [`*Valyala*`](https://github.com/valyala/fastrand)
* [`*NebulousLabs*`](https://gitlab.com/NebulousLabs/fastrand)
* [`*Lukechampine*`](https://lukechampine.com/frand)