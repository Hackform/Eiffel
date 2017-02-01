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
	g.GET("/u/:username", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("GET /users/u/%s", c.Param("username")))
	})
	g.GET("/u/:username/private", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("GET /users/u/%s/private", c.Param("username")))
	})
	g.GET("/id/:userid", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("GET /users/id/%s", c.Param("userid")))
	})
	g.GET("/id/:userid/private", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("GET /users/id/%s/private", c.Param("userid")))
	})
	g.POST("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("POST /users"))
	})
}
