# ABool :bulb:
[![Go Report Card](https://goreportcard.com/badge/github.com/tevino/abool)](https://goreportcard.com/report/github.com/tevino/abool)
[![GoDoc](https://godoc.org/github.com/tevino/abool?status.svg)](https://godoc.org/github.com/tevino/abool)

Atomic Boolean library for Go, optimized for performance yet simple to use.

Use this for cleaner code.

## Usage

```go
import "github.com/tevino/abool"

cond := abool.New()  // default to false

cond.Set()                 // Set to true
cond.IsSet()               // Returns true
cond.UnSet()               // Set to false
cond.SetTo(true)           // Set to whatever you want
cond.SetToIf(false, true)  // Set to true if it is false, returns false(not set)
cond.Toggle() *AtomicBool  // Negates boolean atomically and returns a new AtomicBool object which holds previous boolean value.


// embedding
type Foo struct {
    cond *abool.AtomicBool  // always use pointer to avoid copy
}
```

## Benchmark

- Go 1.11.5
- OS X 10.14.5

```shell
# Read
BenchmarkMutexRead-4            100000000               14.7 ns/op
BenchmarkAtomicValueRead-4      2000000000               0.45 ns/op
BenchmarkAtomicBoolRead-4       2000000000               0.35 ns/op  # <--- This package

# Write
BenchmarkMutexWrite-4           100000000               14.5 ns/op
BenchmarkAtomicValueWrite-4     100000000               10.5 ns/op
BenchmarkAtomicBoolWrite-4      300000000                5.21 ns/op  # <--- This package

# CAS
BenchmarkMutexCAS-4             50000000                31.3 ns/op
BenchmarkAtomicBoolCAS-4        200000000                7.18 ns/op  # <--- This package

# Toggle
BenchmarkMutexToggle-4          50000000                32.6 ns/op
BenchmarkAtomicBoolToggle-4     300000000                5.21 ns/op  # <--- This package
```
