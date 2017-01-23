package user

import (
	"github.com/Hackform/Eiffel/service/repo/cassandra"
)

func CassOpts() *cassandra.CassOpts {
	return cassandra.Opts(NewModel(), []string{"id"}, nil)
}
