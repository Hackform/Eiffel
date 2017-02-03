package main

import (
	"github.com/hackform/eiffel"
	"github.com/hackform/eiffel/route/setup"
	"github.com/hackform/eiffel/route/users"
	"github.com/hackform/eiffel/service/repo/cassandra"
)

func main() {
	e := eiffel.New()

	r := cassandra.New("eiffel_keyspace", []string{"127.0.0.1"}, "eiffel", "tower")

	e.InitService(eiffel.ServiceConfig{
		"repo": r,
	}, []string{"repo"})

	e.InitRoute(
		"/api",
		eiffel.RouteConfig{
			"/users": userroute.New(r),
			"/setup": setuproute.New(r),
		},
	)

	e.Start(":8080", 30)
}
