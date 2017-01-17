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
