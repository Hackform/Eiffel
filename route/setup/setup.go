package setuproute

import (
	"fmt"
	"github.com/Hackform/Eiffel"
	"github.com/Hackform/Eiffel/model/setup"
	"github.com/Hackform/Eiffel/model/user"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
)

type (
	setuproute struct {
		repo repo.Repo
	}

	reqSetup struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	resSetup struct {
		Message string `json:"message"`
	}
)

const (
	minUsernameLength = 1
	minPasswordLength = 8
)

var (
	setupComplete = false

	log = logrus.WithFields(logrus.Fields{
		"module": "setup router",
	})
)

// New creates a new setup router
func New(repo repo.Repo) eiffel.Route {
	return &setuproute{
		repo: repo,
	}
}

func (r *setuproute) Register(g *echo.Group) {
	g.POST("/", func(c echo.Context) error {
		// check if setup already cached and completed
		if setupComplete {
			return c.JSON(http.StatusGone, resSetup{
				Message: "Setup already complete",
			})
		}

		var err error

		// acquire transaction
		var tx repo.Tx
		if tx, err = r.repo.Transaction(); err != nil {
			log.Errorf("Transaction: %s", err)
			return c.JSON(http.StatusInternalServerError, resSetup{
				Message: "Failed setup process: transaction",
			})
		}

		// attempt to select setup
		var k *setup.Model
		if k, err = setup.Select(tx); err == nil && k.Setup {
			setupComplete = true
			return c.JSON(http.StatusGone, resSetup{
				Message: "Setup already complete",
			})
		}

		log.Info("Begin setup process")

		// check request validity
		req := &reqSetup{}
		if err = c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, resSetup{
				Message: "Failed to provide valid setup config",
			})
		}
		if len(req.Username) < minUsernameLength {
			return c.JSON(http.StatusBadRequest, resSetup{
				Message: "No username provided",
			})
		}
		if len(req.Password) < minPasswordLength {
			return c.JSON(http.StatusBadRequest, resSetup{
				Message: fmt.Sprintf("Min password length %d", minPasswordLength),
			})
		}

		// create user
		var newUser *user.Model
		if newUser, err = user.NewSuperUser(req.Username, req.Password); err != nil {
			log.Errorf("Create new user: %s", err)
			return c.JSON(http.StatusInternalServerError, resSetup{
				Message: "Failed to create new user",
			})
		}
		log.Info("Begin database initialization")

		// create setup table
		if err = setup.Create(tx); err != nil {
			log.Errorf("Create setup table: %s", err)
			return c.JSON(http.StatusInternalServerError, resSetup{
				Message: "Failed setup process: create setup table",
			})
		}
		log.Info("Setup table created")

		// create user table
		if err = user.Create(tx); err != nil {
			log.Errorf("Create user table: %s", err)
			return c.JSON(http.StatusInternalServerError, resSetup{
				Message: "Failed setup process: create user table",
			})
		}
		log.Info("User table created")

		// insert new user
		if err = user.Insert(tx, newUser); err != nil {
			log.Errorf("Insert new user: %s", err)
			return c.JSON(http.StatusInternalServerError, resSetup{
				Message: "Failed setup process: insert new user",
			})
		}
		log.Infof("New superuser %s inserted", newUser.Username)

		// insert setup table complete
		if err = setup.Insert(tx); err != nil {
			log.Errorf("Insert setup complete: %s", err)
			return c.JSON(http.StatusInternalServerError, resSetup{
				Message: "Failed setup process: insert setup complete",
			})
		}
		log.Info("Setup complete")

		return c.JSON(http.StatusOK, resSetup{
			Message: "Setup complete",
		})
	})
}
