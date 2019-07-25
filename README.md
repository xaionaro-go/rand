This's an inaccurate (especially not thread-safe) fast PRNG (which is *not* appropriate for cryptographic tasks). It's inaccurate intentionally to make it even faster.goos: linux

This PRNG is represented here as `BenchmarkOurIntn`.

```
goarch: amd64
pkg: github.com/xaionaro-go/fastrand
BenchmarkOurIntn-8             	1000000000	         3.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkStandardIntn-8        	50000000	        27.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkValyalaUintn-8        	50000000	        24.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkNebulousLabsUintn-8   	10000000	       224 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/fastrand	10.201s
```