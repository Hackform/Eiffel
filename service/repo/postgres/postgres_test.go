package postgres

import (
	"github.com/Hackform/Eiffel"
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
	assert.Equal("key_1 = value_1", parseConstraints(qc), "parseConstraints should properly render q_EQUAL")

	a.Condition = q.UNEQUAL
	assert.Equal("key_1 <> value_1", parseConstraints(qc), "parseConstraints should properly render q_UNEQUAL")

	b := q.NewCon("key_2", q.EQUAL, "$")
	qc = q.Constraints{b}
	assert.Equal("key_2 = $1", parseConstraints(qc), "parseConstraints should properly substitute escape_sequence")

	qc = q.Constraints{a, b}
	assert.Equal("key_1 <> value_1 AND key_2 = $1", parseConstraints(qc), "parseConstraints should by default adjoin multiple constraints with 'AND'")

	c := q.NewOp(a, q.OR, b)
	qc = q.Constraints{c}
	assert.Equal("(key_1 <> value_1 OR key_2 = $1)", parseConstraints(qc), "parseConstraints should properly render q_OR")
}

func Test_parseQ(t *testing.T) {
	assert := assert.New(t)
	query := q.NewQOne("test_sector", q.Props{"prop_1", "prop_2", "another_prop"}, nil)
	assert.Equal("SELECT prop_1, prop_2, another_prop FROM test_sector LIMIT 1;", parseQ(query), "parseQ should properly render q_ACTION_QUERY_ONE")

	qc := q.Constraints{q.NewCon("key_1", q.EQUAL, "value_1")}
	query = q.NewQOne("test_sector", q.Props{"prop_1", "prop_2", "another_prop"}, qc)
	assert.Equal("SELECT prop_1, prop_2, another_prop FROM test_sector LIMIT 1 WHERE key_1 = value_1;", parseQ(query), "parseQ should properly render constraints")
}
