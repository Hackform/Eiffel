package main

import (
	"github.com/Hackform/Eiffel"
	userRoute "github.com/Hackform/Eiffel/route/users"
	"github.com/Hackform/Eiffel/service/repo/postgres"
)

func main() {
	e := eiffel.New()
	e.InitService(eiffel.ServiceConfig{
		"repo": postgres.New(postgres.ConnectionString{
			"dbname":                    "hackform",
			"user":                      "",
			"password":                  "",
			"host":                      "localhost",
			"port":                      "5432",
			"connect_timeout":           "8",
			"fallback_application_name": "eiffel",
			"sslmode":                   "disable",
		}),
	})
	e.InitRoute("/api", eiffel.RouteConfig{
		"/users": userRoute.New(),
	})
	e.Start(":8080")
}
