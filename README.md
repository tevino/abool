# ABool :bulb:

[![Go Report Card](https://goreportcard.com/badge/github.com/tevino/abool)](https://goreportcard.com/report/github.com/tevino/abool)
[![GoDoc](https://godoc.org/github.com/tevino/abool?status.svg)](https://godoc.org/github.com/tevino/abool)

Atomic Boolean package for Go, optimized for performance yet simple to use.

Designed for cleaner code.

## Usage

```go
import "github.com/tevino/abool/v2"

cond := abool.New()     // default to false

cond.Set()              // Sets to true
cond.IsSet()            // Returns true
cond.UnSet()            // Sets to false
cond.IsNotSet()         // Returns true
cond.SetTo(any)         // Sets to whatever you want
cond.SetToIf(old, new)  // Sets to `new` only if the Boolean matches the `old`, returns whether succeeded
cond.Toggle()           // Flip the value of `cond`, returns the value before flipping


// embedding
type Foo struct {
    cond *abool.AtomicBool  // always use pointer to avoid copy
}
```

## Benchmark

```
go: v1.18.2
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz

# Read
BenchmarkMutexRead-12           	100000000	          10.24   ns/op
BenchmarkAtomicValueRead-12     	1000000000	         0.4690 ns/op
BenchmarkAtomicBoolRead-12      	1000000000	         0.2345 ns/op  # <--- This package

# Write
BenchmarkMutexWrite-12          	100000000	        10.19  ns/op
BenchmarkAtomicValueWrite-12    	164918696	         7.235 ns/op
BenchmarkAtomicBoolWrite-12     	278729533	         4.341 ns/op  # <--- This package

# CAS
BenchmarkMutexCAS-12            	57333123	        20.26  ns/op
BenchmarkAtomicBoolCAS-12       	203575494	         5.755 ns/op  # <--- This package

# Toggle
BenchmarkAtomicBoolToggle-12    	145249862	         8.196 ns/op  # <--- This package
```

## Special thanks to contributors

- [barryz](https://github.com/barryz)
  - Added the `Toggle` method
- [Lucas Rouckhout](https://github.com/LucasRouckhout)
  - Implemented JSON Unmarshal and Marshal interface
- [Sebastian Schicho](https://github.com/schicho)
  - Reported a regression with test case
  - Reintroduced the `Toggle` method
