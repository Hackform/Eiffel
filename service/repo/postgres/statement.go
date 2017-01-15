package postgres

import (
	"database/sql"
	"github.com/Hackform/Eiffel/service/repo"
)

///////////////
// Statement //
///////////////

type (
	stmt struct {
		s *sql.Stmt
	}
)

func newStmt(s *sql.Stmt) *stmt {
	return &stmt{
		s: s,
	}
}

func (s *stmt) Query(args ...interface{}) ([]*repo.Data, error) {
	return nil, nil
}

func (s *stmt) QueryOne(args ...interface{}) (*repo.Data, error) {
	return nil, nil
}

func (s *stmt) Exec(args ...interface{}) error {
	return nil
}
