package setup

import (
	setupModel "github.com/Hackform/Eiffel/model/setup"
	"github.com/Hackform/Eiffel/model/user"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type (
	setup struct {
		repo repo.Repo
	}

	reqSetup struct {
		username,
		password string
	}

	resSetup struct {
		Message string `json:"message"`
	}
)

const (
	username_length = 1
	password_length = 8
)

var (
	setup_complete = false
)

func New(repo repo.Repo) *setup {
	return &setup{
		repo: repo,
	}
}

func (r *setup) Register(g *echo.Group) {
	g.POST("/", func(c echo.Context) error {
		log := log.WithFields(log.Fields{
			"module": "setup router",
		})

		// check if setup already cached and completed
		if setup_complete {
			return c.JSON(http.StatusGone, resSetup{
				Message: "Setup already complete",
			})
		}

		// acquire transaction
		var tx repo.Tx
		if t, err := r.repo.Transaction(); err != nil {
			log.Errorf("transaction: %s", err)
			return c.JSON(http.StatusInternalServerError, resSetup{
				Message: "Failed setup process: transaction",
			})
		} else {
			tx = t
		}

		// attempt to select setup
		if k, err := setupModel.Select(tx); err == nil && k.Setup {
			setup_complete = true
			return c.JSON(http.StatusGone, resSetup{
				Message: "Setup already complete",
			})
		} else {
			log.Info("Begin setup process")
		}

		// check request validity
		req := &reqSetup{}
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, resSetup{
				Message: "Failed to provide valid setup config",
			})
		}
		if len(req.username) < username_length {
			return c.JSON(http.StatusBadRequest, resSetup{
				Message: "No username provided",
			})
		}
		if len(req.password) < password_length {
			return c.JSON(http.StatusBadRequest, resSetup{
				Message: "Min password length 8",
			})
		}

		// create user
		var newUser *user.UserModel
		if u, err := user.NewSuperUser(req.username, req.password); err != nil {
			log.Errorf("Create new user: %s", err)
			return c.JSON(http.StatusInternalServerError, resSetup{
				Message: "Failed to create new user",
			})
		} else {
			newUser = u
		}

		log.Info("Begin database initialization")

		// create setup table
		if err := setupModel.Create(tx); err != nil {
			log.Errorf("Create setup table: %s", err)
			return c.JSON(http.StatusInternalServerError, resSetup{
				Message: "Failed setup process: create setup table",
			})
		} else {
			log.Info("Setup table created")
		}

		// create user table
		if err := user.Create(tx); err != nil {
			log.Errorf("Create user table: %s", err)
			return c.JSON(http.StatusInternalServerError, resSetup{
				Message: "Failed setup process: create user table",
			})
		} else {
			log.Info("User table created")
		}

		// insert new user
		if err := user.Insert(tx, newUser); err != nil {
			log.Errorf("Insert new user: %s", err)
			return c.JSON(http.StatusInternalServerError, resSetup{
				Message: "Failed setup process: insert new user",
			})
		} else {
			log.Infof("New superuser %s inserted", newUser.Username)
		}

		// insert setup table complete
		if err := setupModel.Insert(tx); err != nil {
			log.Errorf("Insert setup complete: %s", err)
			return c.JSON(http.StatusInternalServerError, resSetup{
				Message: "Failed setup process: insert setup complete",
			})
		} else {
			log.Info("Setup complete")
		}

		return c.JSON(http.StatusOK, resSetup{
			Message: "Setup complete",
		})
	})
}
