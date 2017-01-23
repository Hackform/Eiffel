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
		Commit() error
		Rollback() error
		Insert(string, *Data) error
	}

	Data struct {
		Value interface{}
	}
)
