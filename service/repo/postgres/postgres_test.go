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
	assert.Equal("key_1 = value_1", parseConstraints("$", qc), "parseConstraints should properly render q_EQUAL")

	b := q.NewCon("key_2", q.EQUAL, "$")
	qc = q.Constraints{b}
	assert.Equal("key_2 = $1", parseConstraints("$", qc), "parseConstraints should properly substitute escape_sequence")

	qc = q.Constraints{a, b}
	assert.Equal("key_1 = value_1 and key_2 = $1", parseConstraints("$", qc), "parseConstraints should by default adjoin multiple constraints with 'and'")
}
