This's an inaccurate (especially not thread-safe) fast PRNG (which is *not* appropriate for cryptographic tasks). It's inaccurate intentionally to make it even faster.goos: linux

This PRNG is represented here as `BenchmarkOur*`.

```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/fastrand
BenchmarkOurRead1-8           	228187446	         5.26 ns/op	 190.15 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead15-8          	153651139	         7.32 ns/op	2048.30 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead16-8          	155503356	         7.60 ns/op	2104.39 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurRead1024-8        	 3523779	       341 ns/op	3005.92 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1-8      	41341497	        25.5 ns/op	  39.28 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead15-8     	27557625	        43.8 ns/op	 342.57 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead16-8     	24690036	        43.3 ns/op	 369.48 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardRead1024-8   	  760071	      1588 ns/op	 645.03 MB/s	       0 B/op	       0 allocs/op
BenchmarkOurUint32nMax-8      	434313181	         2.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkOurUint32n-8         	439039567	         2.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkOurUint64n-8         	449298564	         2.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkStandardIntn-8       	49614784	        28.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkValyalaUint32n-8     	47832376	        27.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkNebulousLabsIntn-8   	 6054924	       233 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/fastrand	24.489s
```