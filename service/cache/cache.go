package cache

import ()

type (
	// Cache is an interface to a cache service
	Cache interface {
		Start() error
		Shutdown()
		Transaction() (Tx, error)
	}

	// Tx is a transaction for a Cache
	Tx interface {
		Adapter() string
	}
)
