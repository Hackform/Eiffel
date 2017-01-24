package repo

import ()

type (
	Repo interface {
		Start() error
		Shutdown()
		Transaction() (Tx, error)
		Setup() error
	}

	Tx interface {
		Adapter() string
	}

	Data struct {
		Value interface{}
	}
)
