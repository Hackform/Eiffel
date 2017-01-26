package user

import (
	"errors"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/Hackform/Eiffel/service/repo/cassandra"
	"github.com/Hackform/Eiffel/service/util/eta"
	"github.com/Hackform/Eiffel/service/util/rho"
	"github.com/Hackform/Eiffel/service/util/upsilon"
	"github.com/gocql/gocql"
	"time"
)

type (
	// Model defines a user
	Model struct {
		userID
		userInfo
		passhash
	}

	userID struct {
		ID gocql.UUID `json:"id" cql:"id"`
	}

	userInfo struct {
		em
		userInfoPub
	}

	em struct {
		Email string `json:"email" cql:"email"`
	}

	userInfoPub struct {
		uname
		auth
		name
		props
	}

	uname struct {
		Username string `json:"username" cql:"username"`
	}

	auth struct {
		Level uint8  `json:"auth_level" cql:"auth_level"`
		Tags  []byte `json:"auth_tags" cql:"auth_tags"`
	}

	name struct {
		First string `json:"first_name" cql:"first_name"`
		Last  string `json:"last_name" cql:"last_name"`
	}

	props struct {
		Date time.Time `json:"date" cql:"date"`
	}

	passhash struct {
		Hash    []byte `cql:"pass_hash"`
		Salt    []byte `cql:"pass_salt"`
		Version int    `cql:"pass_version"`
	}
)

// NewModel creates a new Model
func NewModel() *Model {
	return &Model{}
}

// New creates a new Model from arguments
func New(username, password, email, firstName, lastName string, authLevel uint8) (*Model, error) {
	id, err := upsilon.New(8, 0, 8)
	if err != nil {
		return nil, err
	}
	uid, err := gocql.UUIDFromBytes(id.Bytes())
	if err != nil {
		return nil, err
	}
	t := time.Now()
	h, s, v, err := eta.Hash(password, eta.Latest)
	if err != nil {
		return nil, err
	}
	return &Model{
		userID: userID{
			ID: uid,
		},
		userInfo: userInfo{
			em: em{
				Email: email,
			},
			userInfoPub: userInfoPub{
				uname: uname{
					Username: username,
				},
				auth: auth{
					Level: authLevel,
					Tags:  []byte{},
				},
				name: name{
					First: firstName,
					Last:  lastName,
				},
				props: props{
					Date: t,
				},
			},
		},
		passhash: passhash{
			Hash:    h,
			Salt:    s,
			Version: v,
		},
	}, nil
}

// NewUser creates a new User with default access level
func NewUser(username, password, email, firstName, lastName string) (*Model, error) {
	return New(username, password, email, firstName, lastName, rho.User())
}

// NewAdmin creates a new Administrator
func NewAdmin(username, password, email, firstName, lastName string) (*Model, error) {
	return New(username, password, email, firstName, lastName, rho.Admin())
}

// NewSuperUser creates a new superuser
func NewSuperUser(username, password string) (*Model, error) {
	return New(username, password, "", "", "", rho.SuperUser())
}

// Create creates a new User Table on the cassandra cluster
func Create(t repo.Tx) error {
	switch t.Adapter() {
	case cassandra.AdapterID:
		return cassCreate(t.(*cassandra.Tx))
	default:
		return errors.New("Repo adapter not found")
	}
}

// Select finds a given User based on ID
func Select(t repo.Tx, u *upsilon.Upsilon) (*Model, error) {
	switch t.Adapter() {
	case cassandra.AdapterID:
		return cassSelect(t.(*cassandra.Tx), u)
	default:
		return nil, errors.New("Repo adapter not found")
	}
}

// Insert creates a given User
func Insert(t repo.Tx, u *Model) error {
	switch t.Adapter() {
	case cassandra.AdapterID:
		return cassInsert(t.(*cassandra.Tx), u)
	default:
		return errors.New("Repo adapter not found")
	}
}
