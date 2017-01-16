package postgres

import (
	"github.com/Hackform/Eiffel"
	"github.com/Hackform/Eiffel/service/kappa"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/Hackform/Eiffel/service/repo/q"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ConnectionString_stringify(t *testing.T) {
	assert := assert.New(t)
	s := ConnectionString{
		"dbname": "database_name",
		"user":   "username",
		"host":   "localhost",
		"port":   "5432",
	}
	assert.Equal("dbname=database_name user=username host=localhost port=5432", s.stringify(), "stringify should concatenate all the properties of the ConnectionString")

	s2 := ConnectionString{
		"user":   "username",
		"port":   "5432",
		"host":   "localhost",
		"dbname": "database_name",
	}
	assert.Equal(s.stringify(), s2.stringify(), "stringify should concatenate all the properties of the ConnectionString and return them in the same order no matter the original appearance")
}

func Test_Postgres_New(t *testing.T) {
	assert := assert.New(t)
	var pg interface{} = New(ConnectionString{
		"dbname": "database_name",
		"user":   "username",
		"host":   "localhost",
		"port":   "5432",
	})

	_, ok := pg.(eiffel.Service)
	assert.True(ok, "Postgres should implement the eiffel_Service interface")

	_, ok = pg.(repo.Repo)
	assert.True(ok, "Postgres should implement the repo_Repo interface")
}

func Test_Transaction(t *testing.T) {
	assert := assert.New(t)
	var k interface{} = &tx{}
	_, ok := k.(repo.Tx)
	assert.True(ok, "Transaction should implement the repo_Tx interface")
}

func Test_Statement(t *testing.T) {
	assert := assert.New(t)
	var k interface{} = &stmt{}
	_, ok := k.(repo.Stmt)
	assert.True(ok, "Statement should implement the repo_Stmt interface")
}

func Test_parseConstraints(t *testing.T) {
	assert := assert.New(t)
	a := q.NewCon("key_1", q.EQUAL, "value_1")
	qc := q.Constraints{a}
	k := kappa.New()
	assert.Equal("key_1 = value_1", parseConstraints(qc, k), "parseConstraints should properly render q_EQUAL")

	a.Condition = q.UNEQUAL
	k = kappa.New()
	assert.Equal("key_1 <> value_1", parseConstraints(qc, k), "parseConstraints should properly render q_UNEQUAL")

	b := q.NewCon("key_2", q.EQUAL, "$")
	qc = q.Constraints{b}
	k = kappa.New()
	assert.Equal("key_2 = $1", parseConstraints(qc, k), "parseConstraints should properly substitute escape_sequence")

	qc = q.Constraints{a, b}
	k = kappa.New()
	assert.Equal("key_1 <> value_1 AND key_2 = $1", parseConstraints(qc, k), "parseConstraints should by default adjoin multiple constraints with 'AND'")

	c := q.NewOp(a, q.OR, b)
	qc = q.Constraints{c}
	k = kappa.New()
	assert.Equal("(key_1 <> value_1 OR key_2 = $1)", parseConstraints(qc, k), "parseConstraints should properly render q_OR")

	d := q.NewOp(c, q.OR, q.NewCon("key_4", q.EQUAL, "$"))
	qc = q.Constraints{d}
	k = kappa.New()
	assert.Equal("((key_1 <> value_1 OR key_2 = $1) OR key_4 = $2)", parseConstraints(qc, k), "parseConstraints should properly render q_OR")
}

func Test_parseQ_One(t *testing.T) {
	assert := assert.New(t)
	query := q.NewQOne("test_sector", q.Props{"prop_1", "prop_2", "another_prop"}, nil)
	assert.Equal("SELECT prop_1, prop_2, another_prop FROM test_sector;", parseQ(query), "parseQ should properly render q_ACTION_QUERY_ONE")

	qc := q.Constraints{q.NewCon("key_1", q.EQUAL, "value_1")}
	query = q.NewQOne("test_sector", q.Props{"prop_1", "prop_2", "another_prop"}, qc)
	assert.Equal("SELECT prop_1, prop_2, another_prop FROM test_sector WHERE key_1 = value_1;", parseQ(query), "parseQ should properly render constraints")
}

func Test_parseQ_Multi(t *testing.T) {
	assert := assert.New(t)
	qc := q.Constraints{q.NewCon("key_1", q.EQUAL, "value_1")}

	query := q.NewQMulti("test_sector", q.Props{"prop_1", "prop_2", "another_prop"}, nil, 5, nil)
	assert.Equal("SELECT prop_1, prop_2, another_prop FROM test_sector LIMIT 5;", parseQ(query), "parseQ should properly render multi select")

	qorder := q.Constraints{q.NewOrd("prop_1", q.ASC), q.NewOrd("prop_2", q.DESC)}
	query = q.NewQMulti("test_sector", q.Props{"prop_1", "prop_2", "another_prop"}, qc, 5, qorder)
	assert.Equal("SELECT prop_1, prop_2, another_prop FROM test_sector WHERE key_1 = value_1 ORDER BY prop_1 ASC, prop_2 DESC LIMIT 5;", parseQ(query), "parseQ should properly render multi select and constraints")
}

func Test_parseQ_Insert(t *testing.T) {
	assert := assert.New(t)

	query := q.NewI("test_sector", q.Props{"prop_1", "prop_2", "another_prop"}, q.Props{"val_first", "val_2", "val_3"})
	assert.Equal("INSERT INTO test_sector (prop_1, prop_2, another_prop) VALUES (val_first, val_2, val_3);", parseQ(query), "parseQ should properly render insert")

	query = q.NewI("test_sector", q.Props{"prop_1", "prop_2", "another_prop"}, q.Props{"$", "$", "$"})
	assert.Equal("INSERT INTO test_sector (prop_1, prop_2, another_prop) VALUES ($1, $2, $3);", parseQ(query), "parseQ should properly render insert and arg values")
}

func Test_parseQ_Update(t *testing.T) {
	assert := assert.New(t)
	qv := q.Constraints{q.NewEq("key_1", "value_1")}
	qc := q.Constraints{q.NewCon("key_2", q.EQUAL, "value_2")}

	query := q.NewU("test_sector", qv, nil)
	assert.Equal("UPDATE test_sector SET key_1 = value_1;", parseQ(query), "parseQ should properly render update")

	query = q.NewU("test_sector", qv, qc)
	assert.Equal("UPDATE test_sector SET key_1 = value_1 WHERE key_2 = value_2;", parseQ(query), "parseQ should properly render update with constraints")

	qv = q.Constraints{q.NewEq("key_1", "$"), q.NewEq("key_2", "$")}
	qc = q.Constraints{q.NewCon("key_3", q.EQUAL, "$")}
	query = q.NewU("test_sector", qv, qc)
	assert.Equal("UPDATE test_sector SET key_1 = $1, key_2 = $2 WHERE key_3 = $3;", parseQ(query), "parseQ should properly render update with constraints")
}

func Test_parseQ_Delete(t *testing.T) {
	assert := assert.New(t)
	qc := q.Constraints{q.NewCon("key_2", q.EQUAL, "value_2")}

	query := q.NewD("test_sector", nil)
	assert.Equal("DELETE FROM test_sector;", parseQ(query), "parseQ should properly render delete")

	query = q.NewD("test_sector", qc)
	assert.Equal("DELETE FROM test_sector WHERE key_2 = value_2;", parseQ(query), "parseQ should properly render delete with a condition")

	qc = q.Constraints{q.NewEq("key_1", "$"), q.NewEq("key_2", "$")}
	query = q.NewD("test_sector", qc)
	assert.Equal("DELETE FROM test_sector WHERE key_1 = $1 AND key_2 = $2;", parseQ(query), "parseQ should properly render delete with arg conditions")
}

func Test_parseQ_Table(t *testing.T) {
	assert := assert.New(t)

	query := q.NewT("test_sector", q.Constraints{q.NewType("key_1", q.UUID, q.PRIMARY, 0), q.NewType("key_2", q.VARCHAR, q.NOT_NULL_UNIQUE, 32)})
	assert.Equal("CREATE TABLE test_sector (key_1 UUID PRIMARY KEY, key_2 VARCHAR(32) NOT NULL UNIQUE);", parseQ(query), "parseQ should properly render create table")
}
