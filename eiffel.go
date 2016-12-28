package eiffel

import (
	"fmt"
	"github.com/labstack/echo"
	"time"
)

type (
	Eiffel struct {
		services       ServiceConfig
		serviceList    []string
		servicesActive int
		server         *echo.Echo
	}
)

func New() *Eiffel {
	return &Eiffel{
		services:       ServiceConfig{},
		serviceList:    []string{},
		servicesActive: 0,
		server:         echo.New(),
	}
}

func (e *Eiffel) Start(url string) {
	defer e.Shutdown()
	for n, i := range e.serviceList {
		if !e.services[i].Start() {
			fmt.Println(i, "service failed")
			return
		}
		e.servicesActive = n
	}
	e.server.Logger.Fatal(e.server.Start(url))
}

func (e *Eiffel) Shutdown() {
	for n, i := range e.serviceList {
		if n > e.servicesActive {
			e.servicesActive = 0
			break
		}
		e.services[i].Shutdown()
	}
	e.server.Shutdown(5 * time.Second)
}
