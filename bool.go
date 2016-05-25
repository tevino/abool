package abool

import "sync/atomic"

// New creates a pointer to an AtomicBool
func New() *AtomicBool {
	return new(AtomicBool)
}

// AtomicBool is a atomic boolean
// Note: When embedding into a struct, one should always use
// *AtomicBool to avoid copy
type AtomicBool int32

// SetTo sets the boolean with given bool
func (ab *AtomicBool) SetTo(yes bool) {
	if yes {
		atomic.StoreInt32((*int32)(ab), 1)
	} else {
		atomic.StoreInt32((*int32)(ab), 0)
	}
}

// Set sets the bool to true
func (ab *AtomicBool) Set() {
	atomic.StoreInt32((*int32)(ab), 1)
}

// UnSet sets the bool to false
func (ab *AtomicBool) UnSet() {
	atomic.StoreInt32((*int32)(ab), 0)
}

// IsSet returns whether the bool is true
func (ab *AtomicBool) IsSet() bool {
	return atomic.LoadInt32((*int32)(ab)) == 1
}
