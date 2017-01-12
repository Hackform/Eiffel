package postgres

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/Hackform/Eiffel/service/repo/action"
	"github.com/Hackform/Eiffel/service/repo/bound"
	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
)

///////////////////////
// Connection String //
///////////////////////

var (
	connectionArgs = []string{
		"dbname",
		"user",
		"password",
		"host",
		"port",
		"connect_timeout",
		"fallback_application_name",
		"sslmode",
		"sslcert",
		"sslkey",
		"sslrootcert",
	}
)

type (
	ConnectionString map[string]string
)

func New(s ConnectionString) *Postgres {
	return &Postgres{
		dcn: s,
	}
}

func (s *ConnectionString) stringify() string {
	k := bytes.Buffer{}
	l := len(connectionArgs) - 1
	for n, i := range connectionArgs {
		if val, ok := (*s)[i]; ok {
			k.WriteString(i + "=" + val)
			if n < l {
				k.WriteString(" ")
			}
		}
	}
	return k.String()
}

//////////////
// Postgres //
//////////////

type (
	Postgres struct {
		dcn ConnectionString
		db  *sql.DB
	}
)

func (p *Postgres) Start() bool {
	db, err := sql.Open(driverName, p.dcn.stringify())
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	p.db = db
	return true
}

func (p *Postgres) Shutdown() {
	p.db.Close()
}

func (p *Postgres) Transaction() repo.Tx {
	t, err := p.db.Begin()
	if err != nil {
		return nil
	}
	return newTx(t)
}

/////////////////
// Transaction //
/////////////////

type (
	tx struct {
		t *sql.Tx
	}

	stmt struct {
		s *sql.Stmt
	}
)

func newTx(t *sql.Tx) *tx {
	return &tx{
		t: t,
	}
}

func newStmt(s *sql.Stmt) *stmt {
	return &stmt{
		s: s,
	}
}

func parseBound(t *sql.Tx, b bound.Bound) (*stmt, []interface{}) {
	query := bytes.Buffer{}
	args := []interface{}{}
	switch b.Action {
	case action.QUERY_ONE:
		query.WriteString("select ")
		l := len(b.Vals) - 1
		for n, i := range b.Vals {
			query.WriteString(i)
			if n < l {
				query.WriteString(", ")
			}
		}
		query.WriteString(" from " + b.Sector)
		query.WriteString(";")
	}
	stmt, err := t.Prepare(query.String())
	if err != nil {
		return nil, []interface{}{}
	}
	return newStmt(stmt), args
}

func (t *tx) Query(b bound.Bound) []*repo.Data {
	return nil
}

func (t *tx) QueryOne(b bound.Bound) *repo.Data {
	return nil
}

func (t *tx) Exec(b bound.Bound) error {
	return nil
}

func (t *tx) Commit() error {
	return nil
}

func (t *tx) Rollback() error {
	return nil
}
