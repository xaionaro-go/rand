This's an inaccurate (especially not thread-safe) fast PRNG (which is *not* appropriate for cryptographic tasks). It's inaccurate intentionally to make it even faster.goos: linux

This PRNG is represented here as `BenchmarkOur*`.

```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/fastrand
BenchmarkOurRead1-8               	126166237	         8.46 ns/op	 118.18 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead15-8              	129063278	         9.14 ns/op	1641.90 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16-8              	129793394	         9.00 ns/op	1778.00 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16Safe-8          	80423994	        15.2 ns/op	1054.87 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead1024-8            	 6198231	       216 ns/op	4751.56 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead65536-8           	   84787	     14426 ns/op	4542.90 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16777216-8        	     306	   3653427 ns/op	4592.19 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1-8          	41288500	        27.2 ns/op	  36.79 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead15-8         	27717960	        48.3 ns/op	 310.77 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16-8         	25788889	        45.3 ns/op	 352.92 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1024-8       	  707176	      1578 ns/op	 648.89 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead65536-8      	   12357	     96517 ns/op	 679.01 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16777216-8   	      40	  28693837 ns/op	 584.70 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint32nMax-8          	405218553	         2.68 ns/op	1491.55 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint32n-8             	410692437	         3.10 ns/op	1289.68 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint64n-8             	447721760	         2.54 ns/op	3153.34 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint64nSafe-8         	205574767	         5.54 ns/op	1445.20 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardIntn-8           	53668357	        21.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkValyalaUint32n-8         	50937702	        23.0 ns/op	 174.23 MB/s	       0 B/op	       0 allocs/op
BenchmarkNebulousLabsIntn-8       	 5895186	       195 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/fastrand	32.904s
```