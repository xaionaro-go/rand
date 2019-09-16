This's an inaccurate (especially not thread-safe) fast PRNG (which is *not* appropriate for cryptographic tasks). It's inaccurate intentionally to make it even faster.goos: linux

This PRNG is represented here as `BenchmarkOur*`.

```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/fastrand
BenchmarkOurRead1-8               	441302528	         9.13 ns/op	 109.49 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead15-8              	396589956	        10.8 ns/op	1389.22 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16-8              	402581781	         9.33 ns/op	1715.19 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead1024-8            	16676594	       198 ns/op	5171.03 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead65536-8           	  272344	     13909 ns/op	4711.62 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16777216-8        	    1089	   3561305 ns/op	4710.97 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1-8          	136577620	        28.3 ns/op	  35.32 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead15-8         	79265095	        44.2 ns/op	 339.72 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16-8         	72111427	        47.2 ns/op	 339.30 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1024-8       	 1949004	      1587 ns/op	 645.35 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead65536-8      	   37124	    105705 ns/op	 619.99 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16777216-8   	     126	  27769541 ns/op	 604.16 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint32nMax-8          	927874076	         3.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkOurUint32n-8             	1000000000	         3.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkOurUint64n-8             	1000000000	         3.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkStandardIntn-8           	140180552	        25.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkValyalaUint32n-8         	156074469	        26.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkNebulousLabsIntn-8       	14412055	       216 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/fastrand	82.678s
```