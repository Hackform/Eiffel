package main

import (
	"github.com/Hackform/Eiffel"
	"github.com/Hackform/Eiffel/route/setup"
	"github.com/Hackform/Eiffel/route/users"
	"github.com/Hackform/Eiffel/service/repo/cassandra"
	"github.com/labstack/echo/middleware"
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
			"/users": userroute.New(),
			"/setup": setuproute.New(r),
		},
		middleware.Recover(),
		middleware.Logger(),
	)

	e.Start(":8080")
}
