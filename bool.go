// Package abool provides atomic Boolean type for cleaner code and
// better performance.
package abool

import (
	"encoding/json"
	"sync/atomic"
)

// New creates an AtomicBool with default set to false.
func New() *AtomicBool {
	return new(AtomicBool)
}

// NewBool creates an AtomicBool with given default value.
func NewBool(ok bool) *AtomicBool {
	ab := New()
	if ok {
		ab.Set()
	}
	return ab
}

// AtomicBool is an atomic Boolean.
// Its methods are all atomic, thus safe to be called by multiple goroutines simultaneously.
// Note: When embedding into a struct one should always use *AtomicBool to avoid copy.
type AtomicBool struct {
	boolean int32
}

// Set sets the Boolean to true.
func (ab *AtomicBool) Set() {
	atomic.StoreInt32(&ab.boolean, 1)
}

// UnSet sets the Boolean to false.
func (ab *AtomicBool) UnSet() {
	atomic.StoreInt32(&ab.boolean, 0)
}

// IsSet returns whether the Boolean is true.
func (ab *AtomicBool) IsSet() bool {
	return atomic.LoadInt32(&ab.boolean)&1 == 1
}

// IsNotSet returns whether the Boolean is false.
func (ab *AtomicBool) IsNotSet() bool {
	return !ab.IsSet()
}

// SetTo sets the boolean with given Boolean.
func (ab *AtomicBool) SetTo(yes bool) {
	if yes {
		atomic.StoreInt32(&ab.boolean, 1)
	} else {
		atomic.StoreInt32(&ab.boolean, 0)
	}
}

// Toggle inverts the Boolean then returns the value before inverting.
// Based on: https://github.com/uber-go/atomic/blob/3504dfaa1fa414923b1c8693f45d2f6931daf229/bool_ext.go#L40
func (ab *AtomicBool) Toggle() bool {
	var old bool
	for {
		old = ab.IsSet()
		if ab.SetToIf(old, !old) {
			return old
		}
	}
}

// SetToIf sets the Boolean to new only if the Boolean matches the old.
// Returns whether the set was done.
func (ab *AtomicBool) SetToIf(old, new bool) (set bool) {
	var o, n int32
	if old {
		o = 1
	}
	if new {
		n = 1
	}
	return atomic.CompareAndSwapInt32(&ab.boolean, o, n)
}

// MarshalJSON behaves the same as if the AtomicBool is a builtin.bool.
// NOTE: There's no lock during the process, usually it shouldn't be called with other methods in parallel.
func (ab *AtomicBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(ab.IsSet())
}

// UnmarshalJSON behaves the same as if the AtomicBool is a builtin.bool.
// NOTE: There's no lock during the process, usually it shouldn't be called with other methods in parallel.
func (ab *AtomicBool) UnmarshalJSON(b []byte) error {
	var v bool
	err := json.Unmarshal(b, &v)

	if err == nil {
		ab.SetTo(v)
	}
	return err
}
