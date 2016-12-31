package users

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type (
	users struct{}
)

func New() *users {
	return &users{}
}

func (r *users) Register(g *echo.Group) {
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
