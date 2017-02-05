package cache

import ()

type (
	Cache interface {
		Start() error
		Shutdown()
		Transaction() (Tx, error)
	}

	Tx interface {
		Adapter() string
	}
)
