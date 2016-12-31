package main

import (
	"github.com/Hackform/Eiffel"
	userRoute "github.com/Hackform/Eiffel/route/users"
	"github.com/Hackform/Eiffel/service/repo/postgres"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := eiffel.New()
	e.InitService(eiffel.ServiceConfig{
		"repo": postgres.New(postgres.ConnectionString{
			"dbname":                    "hackform",
			"user":                      "eiffel",
			"password":                  "tower",
			"host":                      "localhost",
			"port":                      "5432",
			"connect_timeout":           "16",
			"fallback_application_name": "eiffel",
			"sslmode":                   "disable",
		}),
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
