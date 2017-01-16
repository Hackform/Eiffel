package main

import (
	"github.com/Hackform/Eiffel"
	userRoute "github.com/Hackform/Eiffel/route/users"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := eiffel.New()
	e.InitService(eiffel.ServiceConfig{})
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
