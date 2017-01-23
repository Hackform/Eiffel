package main

import (
	"github.com/Hackform/Eiffel/model/user"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/Hackform/Eiffel/service/repo/cassandra"
)

func repository() repo.Repo {
	return cassandra.New("eiffel_keyspace", []string{"127.0.0.1"}, "eiffel", "tower", cassandra.Config{
		"users": cassandra.Opts(user.Sample(), user.PartitionKeys(), user.ClusterKeys()),
	})
}
