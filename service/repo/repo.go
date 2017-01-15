package repo

import (
	"github.com/Hackform/Eiffel/service/repo/q"
)

type (
	Repo interface {
		Start() bool
		Shutdown()
		Transaction() (Tx, error)
	}

	Tx interface {
		Statement(q.Q) (Stmt, error)
		Commit() error
		Rollback() error
	}

	Stmt interface {
		Query(args ...interface{}) ([]*Data, error)
		QueryOne(args ...interface{}) (*Data, error)
		Exec(args ...interface{}) error
	}

	Data struct {
		Value interface{}
	}
)
