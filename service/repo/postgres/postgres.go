package postgres

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/Hackform/Eiffel/service/kappa"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/Hackform/Eiffel/service/repo/q"
	_ "github.com/lib/pq"
	"strconv"
	"strings"
)

const (
	driverName      = "postgres"
	escape_sequence = "$"
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

func (s *ConnectionString) stringify() string {
	k := bytes.Buffer{}
	for _, i := range connectionArgs {
		if val, ok := (*s)[i]; ok {
			k.WriteString(i + "=" + val + " ")
		}
	}
	return strings.TrimSpace(k.String())
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

func New(s ConnectionString) *Postgres {
	return &Postgres{
		dcn: s,
	}
}

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

func (t *tx) Statement(q.Q) (repo.Stmt, error) {
	return nil, nil
}

func (t *tx) Commit() error {
	return nil
}

func (t *tx) Rollback() error {
	return nil
}

/*
UPDATE T
SET C1 = 1
WHERE C2 = 'a'
INSERT INTO products (product_no, name, price) VALUES
    (1, 'Cheese', 9.99),
    (2, 'Bread', 1.99),
    (3, 'Milk', 2.99);
*/

func parseQuery(qu q.Q) string {
	query := bytes.Buffer{}
	query.WriteString("SELECT ")
	l := len(qu.RProps) - 1
	for n, i := range qu.RProps {
		query.WriteString(i)
		if n < l {
			query.WriteString(", ")
		}
	}
	query.WriteString(" FROM " + qu.Sector)
	if qu.Cons != nil {
		query.WriteString(" WHERE ")
		query.WriteString(parseConstraints(qu.Cons))
	}
	if qu.Action == q.ACTION_QUERY_MULTI {
		query.WriteString(" LIMIT " + strconv.Itoa(qu.Limit) + " OFFSET " + strconv.Itoa(qu.Offset))
	}
	query.WriteString(";")
	return query.String()
}

func parseInsert(qu q.Q) string {
	query := bytes.Buffer{}
	query.WriteString("INSERT INTO " + qu.Sector + " (")
	l := len(qu.RProps) - 1
	for n, i := range qu.RProps {
		query.WriteString(i)
		if n < l {
			query.WriteString(", ")
		}
	}
	query.WriteString(") VALUES (")
	k := kappa.New()
	for n, i := range qu.Vals {
		if i == escape_sequence {
			query.WriteString(escape_sequence + strconv.Itoa(k.Get()))
		} else {
			query.WriteString(i)
		}
		if n < l {
			query.WriteString(", ")
		}
	}
	query.WriteString(");")
	return query.String()
}

func parseQ(qu q.Q) string {
	switch qu.Action {
	case q.ACTION_QUERY_ONE, q.ACTION_QUERY_MULTI:
		return parseQuery(qu)
	case q.ACTION_INSERT:
		return parseInsert(qu)
	}
	return ""
}

func parseConstraints(cons q.Constraints) string {
	k := kappa.New()
	l := len(cons) - 1
	clause := bytes.Buffer{}
	for n, i := range cons {
		c := i.Condition
		if c != q.AND && c != q.OR {
			clause.WriteString(i.Key)
		}

		switch c {
		case q.EQUAL:
			clause.WriteString(" = ")
		case q.UNEQUAL:
			clause.WriteString(" <> ")
		case q.GREATER:
			clause.WriteString(" > ")
		case q.LESSER:
			clause.WriteString(" < ")
		case q.GREATER_EQ:
			clause.WriteString(" >= ")
		case q.LESSER_EQ:
			clause.WriteString(" <= ")

		case q.AND:
			clause.WriteString("(" + parseConstraints(q.Constraints{i.Con1}) + " AND " + parseConstraints(q.Constraints{i.Con2}) + ")")
		case q.OR:
			clause.WriteString("(" + parseConstraints(q.Constraints{i.Con1}) + " OR " + parseConstraints(q.Constraints{i.Con2}) + ")")
		}

		if c != q.AND && c != q.OR {
			if i.Value == escape_sequence {
				clause.WriteString(escape_sequence + strconv.Itoa(k.Get()))
			} else {
				clause.WriteString(i.Value)
			}
		}

		if n < l {
			clause.WriteString(" AND ")
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
