package main

import (
	"github.com/Hackform/Eiffel"
	"github.com/Hackform/Eiffel/route/users"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := eiffel.New()
	e.InitService(eiffel.ServiceConfig{
		"repo": repository(),
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
