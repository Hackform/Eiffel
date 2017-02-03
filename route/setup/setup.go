package setuproute

import (
	"fmt"
	"github.com/hackform/eiffel"
	"github.com/hackform/eiffel/model/setup"
	"github.com/hackform/eiffel/model/user"
	"github.com/hackform/eiffel/service/repo"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"github.com/pressly/lg"
	"net/http"
	"regexp"
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
	minPasswordLength = 8
	maxPasswordLength = 128
)

var (
	setupComplete = false
)

var (
	regexUsername = regexp.MustCompile(`^[a-zA-Z0-9_\-]{1,32}$`)
)

func isValidUsername(username string) bool {
	return regexUsername.MatchString(username)
}
func isValidPassword(password string) bool {
	l := len(password)
	return l >= minPasswordLength && l <= maxPasswordLength
}

// New creates a new setup router
func New(repo repo.Repo) eiffel.Route {
	return &setuproute{
		repo: repo,
	}
}

func (rr *setuproute) Register(r chi.Router) {
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		// check if setup already cached and completed
		if setupComplete {
			render.Status(r, http.StatusGone)
			render.JSON(w, r, resSetup{
				Message: "Setup already complete",
			})
			return
		}
		// acquire transaction
		tx, err := rr.repo.Transaction()
		if err != nil {
			lg.RequestLog(r).WithError(err).Error("Failed to acquire transaction")
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, resSetup{
				Message: "Failed setup process: transaction",
			})
			return
		}
		// attempt to select setup
		k, err := setup.Select(tx)
		if err == nil && k.Setup {
			setupComplete = true
			render.Status(r, http.StatusGone)
			render.JSON(w, r, resSetup{
				Message: "Setup already complete",
			})
			return
		}
		lg.WithField("module", "setuproute").Info("Begin setup process")
		// check request validity
		req := &reqSetup{}
		if err = render.Bind(r.Body, req); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resSetup{
				Message: "Failed to provide valid setup config",
			})
			return
		}
		if !isValidUsername(req.Username) {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resSetup{
				Message: "Username must be between 1 and 32 characters and only be composed of a-z, A-Z, 0-9, _, -",
			})
			return
		}
		if !isValidPassword(req.Password) {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, resSetup{
				Message: fmt.Sprintf("Password must be between %d and %d characters", minPasswordLength, maxPasswordLength),
			})
			return
		}
		// create user
		newUser, err := user.NewSuperUser(req.Username, req.Password)
		if err != nil {
			lg.RequestLog(r).WithError(err).Error("Failed to initialize a new user")
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, resSetup{
				Message: "Failed to initialize a new user",
			})
			return
		}
		lg.Info("Begin database initialization")
		// create setup table
		if err = setup.Create(tx); err != nil {
			lg.RequestLog(r).WithError(err).Error("Failed to create setup table")
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, resSetup{
				Message: "Failed to create setup table",
			})
			return
		}
		lg.Info("Setup table created")
		// create user table
		if err = user.Create(tx); err != nil {
			lg.RequestLog(r).WithError(err).Error("Failed to create user table")
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, resSetup{
				Message: "Failed to create user table",
			})
			return
		}
		lg.Info("User table created")
		// insert new user
		if err = user.Insert(tx, newUser); err != nil {
			lg.RequestLog(r).WithError(err).Error("Failed to insert superuser")
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, resSetup{
				Message: "Failed to insert superuser",
			})
			return
		}
		lg.Infof("New superuser %s inserted", newUser.Username)
		// insert setup table complete
		if err = setup.Insert(tx); err != nil {
			lg.RequestLog(r).WithError(err).Error("Failed to insert setup config")
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, resSetup{
				Message: "Failed to insert setup config",
			})
			return
		}
		lg.Info("Setup complete")
		render.Status(r, http.StatusOK)
		render.JSON(w, r, resSetup{
			Message: "Setup complete",
		})
		return
	})
}
