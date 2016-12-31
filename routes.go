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

func (e *Eiffel) InitRoute(prefix string, r RouteConfig, m ...echo.MiddlewareFunc) {
	group := e.server.Group(prefix)
	for _, m := range m {
		group.Use(m)
	}
	for k, v := range r {
		g := group.Group(k)
		v.Register(g)
	}
}
