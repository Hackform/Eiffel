package userroute

import (
	"context"
	"fmt"
	"github.com/Hackform/Eiffel"
	"github.com/Hackform/Eiffel/model/user"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"github.com/pressly/lg"
	"net/http"
	"regexp"
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

// New creates a user router
func New(repository repo.Repo) eiffel.Route {
	return &userroute{
		repo: repository,
	}
}

var (
	regexUsername = regexp.MustCompile(`[a-zA-Z0-9_]{1,32}`)
)

func isValidUsername(username string) bool {
	return regexUsername.MatchString(username)
}

type (
	contextKey string
)

const (
	ctxKeyUsermodel = contextKey("usermodel")
)

func (rr *userroute) Register(r chi.Router) {
	r.Route("/u/:username", func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				username := chi.URLParam(r, "username")

				if !isValidUsername(username) {
					render.Status(r, http.StatusBadRequest)
					render.JSON(w, r, resUsersErr{
						Message: "Username must be between 1 and 32 characters",
					})
					return
				}

				tx, err := rr.repo.Transaction()
				if err != nil {
					lg.RequestLog(r).WithError(err).Error("Failed to acquire transaction")
					render.Status(r, http.StatusInternalServerError)
					render.JSON(w, r, resUsersErr{
						Message: "Failed transaction",
					})
					return
				}

				usermodel, err := user.SelectByUsername(tx, username)
				if err != nil {
					lg.RequestLog(r).WithError(err).Warn("User not found")
					render.Status(r, http.StatusNotFound)
					render.JSON(w, r, resUsersErr{
						Message: "User not found",
					})
					return
				}

				ctx := context.WithValue(r.Context(), ctxKeyUsermodel, usermodel)
				next.ServeHTTP(w, r.WithContext(ctx))
			})
		})

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			// TODO: status logs
			usermodel := r.Context().Value(ctxKeyUsermodel).(*user.Model)
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, resUsersPublic{
				Data: *usermodel.GetPublic(),
			})
			return
		})

		r.Get("/private", func(w http.ResponseWriter, r *http.Request) {
			// TODO: protect jwt
			usermodel := r.Context().Value(ctxKeyUsermodel).(*user.Model)
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, resUsersPrivate{
				Data: *usermodel.GetPrivate(),
			})
			return
		})
	})

	r.Get("/id/:userid", func(c echo.Context) error {
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
