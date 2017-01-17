package main

import (
	"github.com/Hackform/Eiffel"
	userRoute "github.com/Hackform/Eiffel/route/users"
	"github.com/Hackform/Eiffel/service/repo/cassandra"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := eiffel.New()
	e.InitService(eiffel.ServiceConfig{
		"repo": cassandra.New("eiffel_keyspace", []string{"127.0.0.1"}, "eiffel", "tower"),
	})
	e.InitRoute(
		"/api",
		eiffel.RouteConfig{
			"/users": userRoute.New(),
		},
		middleware.Recover(),
		middleware.Logger(),
	)
	e.Start(":8080")
}
