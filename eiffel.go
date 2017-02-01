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
		if err := e.services[i].Start(); err != nil {
			fmt.Println(err)
			return
		}
		e.servicesActive = n + 1
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

// Shutdown kills the server and kills any other running resources
func (e *Eiffel) Shutdown() {
	if e.serverRunning {
		e.server.Shutdown(5 * time.Second)
		e.serverRunning = false
	}
	for n, i := range e.serviceList {
		if n >= e.servicesActive {
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
	// Service is an injectable interface for services
	Service interface {
		Start() error
		Shutdown()
	}

	// ServiceConfig is a map to simplify service initialization
	ServiceConfig map[string]Service
)

// InitService initializes Eiffel with the provided services
func (e *Eiffel) InitService(s ServiceConfig, order []string) {
	e.serviceList = order
	e.services = s
}

////////////
// Routes //
////////////

type (
	// Route is an injectable interface for routes
	Route interface {
		Register(g *echo.Group)
	}

	// RouteConfig is a map to simplify route initialization
	RouteConfig map[string]Route
)

// InitRoute initializes Eiffel with the provided routes and middleware
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
