package eiffel

import (
	"github.com/labstack/echo"
)

type (
	Route interface {
		Register(g *echo.Group)
	}

	RouteConfig map[string]Route
)

func (e *Eiffel) InitRoute(prefix string, r RouteConfig) {
	group := e.server.Group(prefix)
	for k, v := range r {
		g := group.Group(k)
		v.Register(g)
	}
}
