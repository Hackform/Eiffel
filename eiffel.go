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
		serverRunning  bool
	}
)

func New() *Eiffel {
	return &Eiffel{
		services:       ServiceConfig{},
		serviceList:    []string{},
		servicesActive: 0,
		server:         echo.New(),
		serverRunning:  false,
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
	fmt.Println(`

                    |
                    |
                    A
                  _/X\_
                  \/X\/
                   |V|
                   |A|
                   |V|
                  /XXX\
                  |\/\|
                  |/\/|
                  |\/\|
                  |/\/|
                  |\/\|
                  |/\/|
                 IIIIIII
                 |\/_\/|
                /\// \\/\
                |/|   |\|
               /\X/___\X/\
              IIIIIIIIIIIII
             /'-\/XXXXX\/-'\
           /'.-'/\|/I\|/\'-.'\
          /'\-/_.-"' '"-._ \-/\
         /.-'.'           '.'-.\
        /'\-/               \-/'\
     _/'-'/'_               _'\'-'\_
    '"""""""'               '"""""""'
      ____ __  ____  ____  ____ __
     ||    || ||    ||    ||    ||
     ||==  || ||==  ||==  ||==  ||
     ||___ || ||    ||    ||___ ||__|

    `)
	e.serverRunning = true
	e.server.Logger.Fatal(e.server.Start(url))
}

func (e *Eiffel) Shutdown() {
	if e.serverRunning {
		e.server.Shutdown(5 * time.Second)
		e.serverRunning = false
	}
	for n, i := range e.serviceList {
		if n > e.servicesActive {
			e.servicesActive = 0
			break
		}
		e.services[i].Shutdown()
	}
}

///////////////
// Services  //
///////////////

type (
	Service interface {
		Start() bool
		Shutdown()
	}

	ServiceConfig map[string]Service
)

func (e *Eiffel) InitService(s ServiceConfig) {
	for k, v := range s {
		e.serviceList = append(e.serviceList, k)
		e.services[k] = v
	}
}

////////////
// Routes //
////////////

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
