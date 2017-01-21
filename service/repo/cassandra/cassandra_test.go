package cassandra

import (
	"github.com/Hackform/Eiffel"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Cassandra_New(t *testing.T) {
	assert := assert.New(t)
	cass := New("keyspace", []string{"nodeIp"}, "username", "password", Config{
		"table-1": Opts(struct{}{}, []string{"partition-1"}, []string{"cluster-1"}),
	})

	assert.Implements((*eiffel.Service)(nil), cass, "Cassandra should implement the eiffel.Service interface")
	assert.Implements((*repo.Repo)(nil), cass, "Cassandra should implement the repo.Repo interface")
}
