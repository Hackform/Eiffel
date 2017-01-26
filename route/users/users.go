package userroute

import (
	"fmt"
	"github.com/Hackform/Eiffel"
	"github.com/labstack/echo"
	"net/http"
)

type (
	userroute struct{}
)

// New creates a user router
func New() eiffel.Route {
	return &userroute{}
}

func (r *userroute) Register(g *echo.Group) {
	g.GET("/:username", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("GET /users/%s", c.Param("username")))
	})
	g.GET("/:username/private", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("GET /users/%s/private", c.Param("username")))
	})
	g.POST("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("POST /users"))
	})
}
