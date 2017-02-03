package userroute

import (
	"fmt"
	"github.com/hackform/eiffel"
	"github.com/hackform/eiffel/model/user"
	"github.com/hackform/eiffel/service/repo"
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
	regexUsername = regexp.MustCompile(`^[a-zA-Z0-9_\-]{1,32}$`)
	regexUserid   = regexp.MustCompile(`^[a-zA-Z0-9_\-=]{1,32}$`)
)

func isValidUsername(username string) bool {
	return regexUsername.MatchString(username)
}
func isValidUserid(userid string) bool {
	return regexUserid.MatchString(userid)
}

func (rr *userroute) Register(r chi.Router) {
	r.Route("/u/:username", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			username := chi.URLParam(r, "username")
			if !isValidUsername(username) {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, resUsersErr{
					Message: "Username must be between 1 and 32 characters and only be composed of a-z, A-Z, 0-9, _, -",
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
			// TODO: status logs
			render.Status(r, http.StatusOK)
			render.JSON(w, r, resUsersPublic{
				Data: *usermodel.GetPublic(),
			})
			return
		})

		r.Get("/private", func(w http.ResponseWriter, r *http.Request) {
			// TODO: protect jwt
			username := chi.URLParam(r, "username")
			if !isValidUsername(username) {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, resUsersErr{
					Message: "Username must be between 1 and 32 characters and only be composed of a-z, A-Z, 0-9, _, -",
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
			render.Status(r, http.StatusOK)
			render.JSON(w, r, resUsersPrivate{
				Data: *usermodel.GetPrivate(),
			})
			return
		})
	})

	r.Route("/id/:userid", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			userid := chi.URLParam(r, "userid")
			if !isValidUserid(userid) {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, resUsersErr{
					Message: "Invalid userid",
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
			usermodel, err := user.Select(tx, userid)
			if err != nil {
				lg.RequestLog(r).WithError(err).Warn("User not found")
				render.Status(r, http.StatusNotFound)
				render.JSON(w, r, resUsersErr{
					Message: "User not found",
				})
				return
			}
			render.Status(r, http.StatusOK)
			render.JSON(w, r, resUsersPublic{
				Data: *usermodel.GetPublic(),
			})
			return
		})

		r.Get("/private", func(w http.ResponseWriter, r *http.Request) {
			// TODO: get jwt middleware
			userid := chi.URLParam(r, "userid")
			if !isValidUserid(userid) {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, resUsersErr{
					Message: "Invalid userid",
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
			usermodel, err := user.Select(tx, userid)
			if err != nil {
				lg.RequestLog(r).WithError(err).Warn("User not found")
				render.Status(r, http.StatusNotFound)
				render.JSON(w, r, resUsersErr{
					Message: "User not found",
				})
				return
			}
			render.Status(r, http.StatusOK)
			render.JSON(w, r, resUsersPrivate{
				Data: *usermodel.GetPrivate(),
			})
			return
		})
	})

	r.Post("/:nonce", func(w http.ResponseWriter, r *http.Request) {
		nonce := chi.URLParam(r, "nonce")
		render.PlainText(w, r, fmt.Sprintf("create new user %s", nonce))
		return
	})
}
