package postgres

import (
	"bytes"
	"database/sql"
	"github.com/Hackform/Eiffel/service/kappa"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/Hackform/Eiffel/service/repo/q"
	"strconv"
)

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
		k := kappa.New()
		query.WriteString(parseConstraints(qu.Cons, k))
	}
	if qu.Action == q.ACTION_QUERY_MULTI {
		if qu.Order != nil {
			query.WriteString(" ORDER BY ")
			l := len(qu.Order) - 1
			for n, i := range qu.Order {
				query.WriteString(i.Key)
				switch i.Condition {
				case q.ASC:
					query.WriteString(" ASC")
				case q.DESC:
					query.WriteString(" DESC")
				}
				if n < l {
					query.WriteString(", ")
				}
			}
		}
		query.WriteString(" LIMIT " + strconv.Itoa(qu.Limit))
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

func parseUpdate(qu q.Q) string {
	query := bytes.Buffer{}
	query.WriteString("UPDATE " + qu.Sector + " SET ")
	k := kappa.New()
	l := len(qu.Mods) - 1
	for n, i := range qu.Mods {
		query.WriteString(i.Key + " = ")
		if i.Value == escape_sequence {
			query.WriteString(escape_sequence + strconv.Itoa(k.Get()))
		} else {
			query.WriteString(i.Value)
		}
		if n < l {
			query.WriteString(", ")
		}
	}
	if qu.Cons != nil {
		query.WriteString(" WHERE ")
		query.WriteString(parseConstraints(qu.Cons, k))
	}
	query.WriteString(";")
	return query.String()
}

func parseDelete(qu q.Q) string {
	query := bytes.Buffer{}
	query.WriteString("DELETE FROM " + qu.Sector)
	if qu.Cons != nil {
		k := kappa.New()
		query.WriteString(" WHERE ")
		query.WriteString(parseConstraints(qu.Cons, k))
	}
	query.WriteString(";")
	return query.String()
}

func parseQ(qu q.Q) string {
	switch qu.Action {
	case q.ACTION_QUERY_ONE, q.ACTION_QUERY_MULTI:
		return parseQuery(qu)
	case q.ACTION_INSERT:
		return parseInsert(qu)
	case q.ACTION_UPDATE:
		return parseUpdate(qu)
	case q.ACTION_DELETE:
		return parseDelete(qu)
	}
	return ""
}

func parseConstraints(cons q.Constraints, k *kappa.Kappa) string {
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
			clause.WriteString("(" + parseConstraints(q.Constraints{i.Con1}, k) + " AND " + parseConstraints(q.Constraints{i.Con2}, k) + ")")
		case q.OR:
			clause.WriteString("(" + parseConstraints(q.Constraints{i.Con1}, k) + " OR " + parseConstraints(q.Constraints{i.Con2}, k) + ")")
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
