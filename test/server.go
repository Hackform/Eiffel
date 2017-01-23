package main

import (
	"github.com/Hackform/Eiffel"
	"github.com/Hackform/Eiffel/model/user"
	"github.com/Hackform/Eiffel/route/users"
	"github.com/Hackform/Eiffel/service/repo/cassandra"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := eiffel.New()

	r := cassandra.New("eiffel_keyspace", []string{"127.0.0.1"}, "eiffel", "tower", cassandra.Config{
		"users": user.CassOpts(),
	})

	e.InitService(eiffel.ServiceConfig{
		"repo": r,
	})

	e.InitRoute(
		"/api",
		eiffel.RouteConfig{
			"/users": users.New(),
		},
		middleware.Recover(),
		middleware.Logger(),
	)

	e.Start(":8080")
}
