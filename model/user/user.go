package user

import (
	"errors"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/Hackform/Eiffel/service/repo/cassandra"
	"github.com/Hackform/Eiffel/service/util/eta"
	"github.com/Hackform/Eiffel/service/util/rho"
	"github.com/Hackform/Eiffel/service/util/upsilon"
	"time"
)

const (
	modelIDTimeBits = 8
	modelIDHashBits = 0
	modelIDRandBits = 8
)

type (
	// Model defines a user
	Model struct {
		userID
		userInfo
		passhash
	}

	// UnameMap maps a username to a user ID
	UnameMap struct {
		uname
		userID
	}

	// ModelPublic defines the public properties of a user
	ModelPublic struct {
		ID string `json:"id"`
		userInfoPub
	}

	// ModelPrivate defines the public and private properties of a user
	ModelPrivate struct {
		ID string `json:"id"`
		userInfo
	}

	userID struct {
		ID *upsilon.Upsilon `json:"id" cql:"id"`
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

// NewUnameMap creates a new Username map to user ID
func NewUnameMap(username string, id *upsilon.Upsilon) (*UnameMap, error) {
	return &UnameMap{
		userID: userID{
			ID: id,
		},
		uname: uname{
			Username: username,
		},
	}, nil
}

// New creates a new Model from arguments
func New(username, password, email, firstName, lastName string, authLevel uint8) (*Model, error) {
	id, err := upsilon.New(modelIDTimeBits, modelIDHashBits, modelIDRandBits)
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
			ID: id,
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

// GetPublic returns the portion of the Model that is public
func (m *Model) GetPublic() *ModelPublic {
	return &ModelPublic{
		m.ID.Base64(),
		m.userInfoPub,
	}
}

// GetPrivate returns the portion of the Model that is public and private
func (m *Model) GetPrivate() *ModelPrivate {
	return &ModelPrivate{
		m.ID.Base64(),
		m.userInfo,
	}
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
func Select(t repo.Tx, userid string) (*Model, error) {
	u, err := upsilon.FromBase64(modelIDTimeBits, modelIDHashBits, modelIDRandBits, userid)
	if err != nil {
		return nil, err
	}

	switch t.Adapter() {
	case cassandra.AdapterID:
		return cassSelect(t.(*cassandra.Tx), u)
	default:
		return nil, errors.New("Repo adapter not found")
	}
}

// SelectByUsername finds a given User based on username
func SelectByUsername(t repo.Tx, username string) (*Model, error) {
	switch t.Adapter() {
	case cassandra.AdapterID:
		return cassSelectByUsername(t.(*cassandra.Tx), username)
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
