package main

import (
	"github.com/Hackform/Eiffel"
	userRoute "github.com/Hackform/Eiffel/route/users"
)

func main() {
	e := eiffel.New()
	e.InitRoute("/api", eiffel.RouteConfig{
		"/users": userRoute.New(),
	})
	e.Start(":8080")
}
