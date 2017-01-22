package setup

import (
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

func New(repo repo.Repo) *setup {
	return &setup{
		repo: repo,
	}
}

func (r *setup) Register(g *echo.Group) {
	g.POST("/", func(c echo.Context) error {
		if err := r.repo.Setup(); err != nil {
			return c.JSON(http.StatusGone, resSetup{
				Message: "Setup already complete",
			})
		}
		k := &reqSetup{}
		if err := c.Bind(k); err != nil {
			return c.JSON(http.StatusBadRequest, resSetup{
				Message: "Failed to provide valid setup config",
			})
		}
		if len(k.username) < username_length {
			return c.JSON(http.StatusBadRequest, resSetup{
				Message: "No username provided",
			})
		}
		if len(k.password) < password_length {
			return c.JSON(http.StatusBadRequest, resSetup{
				Message: "Min password length 8",
			})
		}

		// create user

		return c.JSON(http.StatusOK, resSetup{
			Message: "Setup complete",
		})
	})
}
