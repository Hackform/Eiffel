package repo

import ()

type (
	Repo interface {
		Start() bool
		Shutdown()
		Transaction() (Tx, error)
	}

	Tx interface {
		Commit() error
		Rollback() error
	}

	// Stmt interface {
	// 	Query(args ...interface{}) ([]*Data, error)
	// 	QueryOne(args ...interface{}) (*Data, error)
	// 	Exec(args ...interface{}) error
	// }

	Data struct {
		Value interface{}
	}
)
