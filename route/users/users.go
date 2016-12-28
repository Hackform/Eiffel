package users

import (
	"github.com/labstack/echo"
)

type (
	users struct{}
)

func New() *users {
	return &users{}
}

func (r *users) Register(g *echo.Group) {

}
