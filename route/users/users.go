package userroute

import (
	"fmt"
	"github.com/Hackform/Eiffel"
	"github.com/Hackform/Eiffel/model/user"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/Sirupsen/logrus"
	"net/http"
)

type (
	userroute struct {
		repo repo.Repo
	}

	resUsersPrivate struct {
		Data user.ModelPrivate `json:"user"`
	}

	resUsersPublic struct {
		Data user.ModelPublic `json:"user"`
	}

	resUsersErr struct {
		Message string
	}
)

const (
	minUsernameLength = 1
)

var (
	log = logrus.WithFields(logrus.Fields{
		"module": "users router",
	})
)

// New creates a user router
func New() eiffel.Route {
	return &userroute{}
}

func (r *userroute) Register(g *echo.Group) {
	g.GET("/u/:username", func(c echo.Context) error {
		username := c.Param("username")
		if len(username) < minUsernameLength {
			return c.JSON(http.StatusBadRequest, resUsersErr{
				Message: "No username provided",
			})
		}

		var tx repo.Tx
		var err error
		if tx, err = r.repo.Transaction(); err != nil {
			log.Errorf("transaction: %s", err)
			return c.JSON(http.StatusInternalServerError, resUsersErr{
				Message: "Failed transaction",
			})
		}

		usermodel, err := user.SelectByUsername(tx, username)
		if err != nil {
			log.Warnf("Select user by username: %s", err)
			return c.JSON(http.StatusNotFound, resUsersErr{
				Message: "Failed to find user",
			})
		}

		return c.JSON(http.StatusOK, resUsersPublic{
			Data: *usermodel.GetPublic(),
		})
	})

	g.GET("/u/:username/private", func(c echo.Context) error {
		username := c.Param("username")
		if len(username) < minUsernameLength {
			return c.JSON(http.StatusBadRequest, resUsersErr{
				Message: "No username provided",
			})
		}

		var tx repo.Tx
		var err error
		if tx, err = r.repo.Transaction(); err != nil {
			log.Errorf("transaction: %s", err)
			return c.JSON(http.StatusInternalServerError, resUsersErr{
				Message: "Failed transaction",
			})
		}

		usermodel, err := user.SelectByUsername(tx, username)
		if err != nil {
			log.Warnf("Select user by username: %s", err)
			return c.JSON(http.StatusNotFound, resUsersErr{
				Message: "Failed to find user",
			})
		}

		return c.JSON(http.StatusOK, resUsersPrivate{
			Data: *usermodel.GetPrivate(),
		})
	}) // TODO: get jwt middleware

	g.GET("/id/:userid", func(c echo.Context) error {
		userid := c.Param("userid")

		var tx repo.Tx
		var err error
		if tx, err = r.repo.Transaction(); err != nil {
			log.Errorf("transaction: %s", err)
			return c.JSON(http.StatusInternalServerError, resUsersErr{
				Message: "Failed transaction",
			})
		}

		usermodel, err := user.Select(tx, userid)
		if err != nil {
			log.Warnf("Select user by id: %s", err)
			return c.JSON(http.StatusNotFound, resUsersErr{
				Message: "Failed to find user",
			})
		}

		return c.JSON(http.StatusOK, resUsersPublic{
			Data: *usermodel.GetPublic(),
		})
	})

	g.GET("/id/:userid/private", func(c echo.Context) error {
		userid := c.Param("userid")

		var tx repo.Tx
		var err error
		if tx, err = r.repo.Transaction(); err != nil {
			log.Errorf("transaction: %s", err)
			return c.JSON(http.StatusInternalServerError, resUsersErr{
				Message: "Failed transaction",
			})
		}

		usermodel, err := user.Select(tx, userid)
		if err != nil {
			log.Warnf("Select user by id: %s", err)
			return c.JSON(http.StatusNotFound, resUsersErr{
				Message: "Failed to find user",
			})
		}

		return c.JSON(http.StatusOK, resUsersPrivate{
			Data: *usermodel.GetPrivate(),
		})
	}) // TODO: get jwt middleware

	g.POST("/:code", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("POST /users"))
	})
}
