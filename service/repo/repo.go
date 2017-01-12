package repo

import (
	"github.com/Hackform/Eiffel/service/repo/bound"
)

type (
	Data struct {
		Value interface{}
	}

	Tx interface {
		Query(bound.Bound) []*Data
		QueryOne(bound.Bound) *Data
		Exec(bound.Bound) error
		Commit() error
		Rollback() error
	}

	Repo interface {
		Start() bool
		Shutdown()
		Transaction() Tx
	}
)
