package cassandra

import (
	"github.com/Hackform/Eiffel"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Cassandra_New(t *testing.T) {
	assert := assert.New(t)
	var cass interface{} = New("keyspace", []string{"nodeIp"}, "username", "password")

	_, ok := cass.(eiffel.Service)
	assert.True(ok, "Cassandra should implement the eiffel.Service interface")

	_, ok = cass.(repo.Repo)
	assert.True(ok, "Cassandra should implement the repo.Repo interface")
}
