package repo

import ()

type (
	Repo interface {
		Start() error
		Shutdown()
		Transaction() (Tx, error)
	}

	Tx interface {
		Adapter() string
	}

	Data struct {
		Value interface{}
	}
)
