package eiffel

import (
	"context"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/pressly/lg"
	"net/http"
	"time"
)

const (
	eiffelVersion = "v0.1.0"
)

type (
	// Eiffel is a service Gateway for Hackform
	Eiffel struct {
		services       ServiceConfig
		serviceList    []string
		servicesActive int
		server         *chi.Mux
		serverRunning  bool
		logger         *logrus.Logger
	}
)

// New creates a new Eiffel
func New() *Eiffel {
	logger := logrus.New()
	// TODO: allow output and type to be configured
	// logger.Formatter = &logrus.JSONFormatter{}

	lg.RedirectStdlogOutput(logger)
	lg.DefaultLogger = logger

	logger.WithFields(logrus.Fields{
		"module": "eiffel",
	}).Info("Initializing gateway")

	return &Eiffel{
		services:       ServiceConfig{},
		serviceList:    []string{},
		servicesActive: 0,
		server:         chi.NewRouter(),
		serverRunning:  false,
		logger:         logger,
	}
}

// Start starts all the services and the server
func (e *Eiffel) Start(url string, shutdownDelay int) {
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

	lg.Log(lg.WithLoggerContext(context.Background(), e.logger)).WithFields(logrus.Fields{
		"module":  "eiffel",
		"version": eiffelVersion,
	}).Infof("Server started on %s", url)

	e.serverRunning = true
	// TODO: gracefully shutdown
	http.ListenAndServe(url, e.server)
}

// Shutdown kills the server and kills any other running resources
func (e *Eiffel) Shutdown() {
	e.logger.WithFields(logrus.Fields{
		"module": "eiffel",
	}).Info("Server shutting down")

	if e.serverRunning {
		// TODO: graceful shutdown
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
		SetLogger(*logrus.Logger)
		Register(g chi.Router)
	}

	// RouteConfig is a map to simplify route initialization
	RouteConfig map[string]Route
)

// InitRoute initializes Eiffel with the provided routes and middleware
func (e *Eiffel) InitRoute(prefix string, routeConfig RouteConfig, m ...func(http.Handler) http.Handler) {
	r := chi.NewRouter()
	r.Use(middleware.StripSlashes)
	r.Use(middleware.RequestID, middleware.RealIP, lg.RequestLogger(e.logger), middleware.CloseNotify, middleware.Timeout(30*time.Second))
	r.Use(middleware.DefaultCompress)
	r.Use(m...)
	for k, v := range routeConfig {
		v.SetLogger(e.logger)
		r.Route(k, v.Register)
	}
	e.server.Mount(prefix, r)
}
