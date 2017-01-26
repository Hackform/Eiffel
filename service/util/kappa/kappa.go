package kappa

import (
	"sync/atomic"
)

//////////////////////////
// Incremented Constant //
//////////////////////////

type (
	// Kappa is a thread-safe universal counter
	Kappa struct {
		value int32
	}
)

// New creates a new Kappa
func New() *Kappa {
	return &Kappa{
		value: 0,
	}
}

// Get returns the next incremented value
func (k *Kappa) Get() int {
	return int(atomic.AddInt32(&k.value, 1))
}
