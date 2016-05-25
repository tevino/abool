# ABool
:bulb: Atomic boolean library for Golang, optimized for performance yet simple to use.

## Benchmark:

- Golang 1.6.2
- OS X 10.11.4

```
# Read
BenchmarkMutexRead-4       	100000000	        21.1 ns/op
BenchmarkAtomicValueRead-4 	200000000	         6.33 ns/op
BenchmarkAtomicBoolRead-4  	300000000	         4.28 ns/op

# Write
BenchmarkMutexWrite-4      	100000000	        21.7 ns/op
BenchmarkAtomicValueWrite-4	30000000	        47.8 ns/op
BenchmarkAtomicBoolWrite-4 	200000000	         9.83 ns/op
```

## Usage

```
var v AtomicBool
v.Set() // set to true
v.IsSet()  // returns true
v.UnSet()  // set to false
v.SetTo(true)  // set with gieven boolean
```
