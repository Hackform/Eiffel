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
	UserModel struct {
		userId
		userInfo
		passhash
	}

	userId struct {
		Id gocql.UUID `json:"id" cql:"id"`
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

func NewModel() *UserModel {
	return &UserModel{}
}

func New(username, password, email, first_name, last_name string, auth_level uint8) (*UserModel, error) {
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
	return &UserModel{
		userId: userId{
			Id: uid,
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
					Level: auth_level,
					Tags:  []byte{},
				},
				name: name{
					First: first_name,
					Last:  last_name,
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

func NewUser(username, password, email, first_name, last_name string) (*UserModel, error) {
	return New(username, password, email, first_name, last_name, rho.User())
}

func NewAdmin(username, password, email, first_name, last_name string) (*UserModel, error) {
	return New(username, password, email, first_name, last_name, rho.Admin())
}

func NewSuperUser(username, password string) (*UserModel, error) {
	return New(username, password, "", "", "", rho.SuperUser())
}

func Create(t repo.Tx) error {
	switch t.Adapter() {
	case cassandra.AdapterId:
		return cassCreate(t.(*cassandra.Tx))
	default:
		return errors.New("Repo adapter not found")
	}
}

func Select(t repo.Tx) (*UserModel, error) {
	switch t.Adapter() {
	case cassandra.AdapterId:
		return cassSelect(t.(*cassandra.Tx))
	default:
		return nil, errors.New("Repo adapter not found")
	}
}

func Insert(t repo.Tx, u *UserModel) error {
	switch t.Adapter() {
	case cassandra.AdapterId:
		return cassInsert(t.(*cassandra.Tx), u)
	default:
		return errors.New("Repo adapter not found")
	}
}
