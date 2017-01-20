package repo

import ()

type (
	Repo interface {
		Start() bool
		Shutdown()
		Transaction(string, *Opts) (Tx, error)
	}

	Opts map[string]interface{}

	Tx interface {
		// Statement(q.Q) (Stmt, error)
		// Commit() error
		// Rollback() error
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
