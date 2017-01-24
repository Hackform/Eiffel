package setup

import (
	setupModel "github.com/Hackform/Eiffel/model/setup"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/labstack/echo"
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
		if setup_complete {
			return c.JSON(http.StatusGone, resSetup{
				Message: "Setup already complete",
			})
		}

		var t repo.Tx
		var err error

		if t, err = r.repo.Transaction(); err != nil {
			return c.JSON(http.StatusInternalServerError, resSetup{
				Message: "Failed setup process: transaction error",
			})
		}

		var k *setupModel.SetupModel
		if k, err = setupModel.Select(t); err == nil && k.Setup {
			setup_complete = true
			return c.JSON(http.StatusGone, resSetup{
				Message: "Setup already complete",
			})
		}

		if err = setupModel.Create(t); err != nil {
			return c.JSON(http.StatusInternalServerError, resSetup{
				Message: "Failed setup process: create setup error",
			})
		}

		// create users table

		r := &reqSetup{}
		if err := c.Bind(k); err != nil {
			return c.JSON(http.StatusBadRequest, resSetup{
				Message: "Failed to provide valid setup config",
			})
		}
		if len(r.username) < username_length {
			return c.JSON(http.StatusBadRequest, resSetup{
				Message: "No username provided",
			})
		}
		if len(r.password) < password_length {
			return c.JSON(http.StatusBadRequest, resSetup{
				Message: "Min password length 8",
			})
		}

		// create user

		// insert into setup
		// toggle setup complete

		return c.JSON(http.StatusOK, resSetup{
			Message: "Setup complete",
		})
	})
}
