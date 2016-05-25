# ABool
[![GoDoc](https://godoc.org/github.com/tevino/abool?status.svg)](https://godoc.org/github.com/tevino/abool)

:bulb: Atomic boolean library for Golang, optimized for performance yet simple to use.

Use this for cleaner code.

## Usage

```go
import "github.com/tevino/abool"

cond := abool.New()  // default to false

cond.Set()           // set to true
cond.IsSet()         // returns true
cond.UnSet()         // set to false
cond.SetTo(true)     // set to whatever you want


// embedding
type Foo struct {
    cond *abool.AtomicBool  // always use pointer to avoid copy
}
```

## Benchmark:

- Golang 1.6.2
- OS X 10.11.4

```shell
# Read
BenchmarkMutexRead-4       	100000000	        21.1 ns/op
BenchmarkAtomicValueRead-4 	200000000	         6.33 ns/op
BenchmarkAtomicBoolRead-4  	300000000	         4.28 ns/op  # <--- This package

# Write
BenchmarkMutexWrite-4      	100000000	        21.7 ns/op
BenchmarkAtomicValueWrite-4	 30000000	        47.8 ns/op
BenchmarkAtomicBoolWrite-4 	200000000	         9.83 ns/op  # <--- This package
```

