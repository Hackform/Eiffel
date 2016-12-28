package eiffel

import (
	"github.com/labstack/echo"
)

type (
	Eiffel struct {
		services ServiceConfig
		server   *echo.Echo
	}
)

func New() *Eiffel {
	return &Eiffel{
		services: ServiceConfig{},
		server:   echo.New(),
	}
}

func (e *Eiffel) Start(url string) {
	e.server.Logger.Fatal(e.server.Start(url))
}
