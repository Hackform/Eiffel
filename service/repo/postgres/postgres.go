package postgres

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/Hackform/Eiffel/service/kappa"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/Hackform/Eiffel/service/repo/action"
	"github.com/Hackform/Eiffel/service/repo/bound"
	_ "github.com/lib/pq"
	"strconv"
)

const (
	driverName = "postgres"
	arg_escape = "$"
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

func (p *Postgres) Transaction() (repo.Tx, error) {
	t, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	return newTx(t), nil
}

/////////////////
// Transaction //
/////////////////

type (
	tx struct {
		t *sql.Tx
	}
)

func newTx(t *sql.Tx) *tx {
	return &tx{
		t: t,
	}
}

func (t *tx) EscapeSequence() string {
	return arg_escape
}

func (t *tx) Statement(bound.Bound) (repo.Stmt, error) {
	return nil, nil
}

func (t *tx) Commit() error {
	return nil
}

func (t *tx) Rollback() error {
	return nil
}

func parseBound(escape_sequence string, b bound.Bound) string {
	query := bytes.Buffer{}
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
		if len(b.Cons) > 0 {
			query.WriteString(" where ")
			query.WriteString(parseConstraints(escape_sequence, b.Cons))
		}
		query.WriteString(";")
	}
	return query.String()
}

func parseConstraints(escape_sequence string, cons bound.Constraints) string {
	k := kappa.New()
	l := len(cons) - 1
	clause := bytes.Buffer{}
	for n, i := range cons {
		switch i.Condition {
		case action.EQUAL:
			clause.WriteString(i.Key + " = ")
			if i.Value == escape_sequence {
				clause.WriteString(escape_sequence + strconv.Itoa(k.Get()))
			} else {
				clause.WriteString(i.Value)
			}
		}
		if n < l {
			clause.WriteString(" and ")
		}
	}
	return clause.String()
}

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
