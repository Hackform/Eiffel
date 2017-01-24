package cassandra

import (
	"github.com/Hackform/Eiffel"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Cassandra_New(t *testing.T) {
	assert := assert.New(t)
	cass := New("keyspace", []string{"nodeIp"}, "username", "password")

	assert.Implements((*eiffel.Service)(nil), cass, "Cassandra should implement the eiffel.Service interface")
	assert.Implements((*repo.Repo)(nil), cass, "Cassandra should implement the repo.Repo interface")
}

func Test_BuilderTable(t *testing.T) {
	assert := assert.New(t)
	v, err := BuilderTable("table_1", Fields{
		"field_2": "set<text>",
		"field_1": "int",
	}, []string{"field_1"}, []string{"field_2"})
	assert.Nil(err, "table properties should be valid")
	assert.Equal("CREATE TABLE table_1 (field_1 int, field_2 set<text>, PRIMARY KEY (field_1, field_2))", v, "should have alphabetized fields and no paretheses around the partition key")

	v, err = BuilderTable("table_1", Fields{
		"field_2": "set<text>",
		"field_3": "int",
		"field_1": "int",
	}, []string{"field_1", "field_3"}, []string{"field_2"})
	assert.Nil(err, "table properties should be valid")
	assert.Equal("CREATE TABLE table_1 (field_1 int, field_2 set<text>, field_3 int, PRIMARY KEY ((field_1, field_3), field_2))", v, "should have alphabetized fields and paretheses around the composite partition key")

	v, err = BuilderTable("table_1", Fields{
		"field_2": "set<text>",
		"field_3": "int",
		"field_1": "int",
	}, []string{"field_1", "field_3"}, []string{})
	assert.Nil(err, "table properties should be valid")
	assert.Equal("CREATE TABLE table_1 (field_1 int, field_2 set<text>, field_3 int, PRIMARY KEY ((field_1, field_3)))", v, "cluster key is not mandatory")
}
