package repo

import (
	"github.com/Hackform/Eiffel/service/repo/bound"
)

type (
	Repo interface {
		Start() bool
		Shutdown()
		Transaction() (Tx, error)
	}

	Tx interface {
		EscapeSequence() string
		Statement(bound.Bound) (Stmt, error)
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
