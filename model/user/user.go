package user

import (
	"github.com/Hackform/Eiffel/service/util/eta"
	"github.com/Hackform/Eiffel/service/util/upsilon"
	"github.com/gocql/gocql"
	"time"
)

type (
	UserModel struct {
		userId
		userInfo
		passhash
		auth
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
		name
		props
	}

	uname struct {
		Username string `json:"username" cql:"username"`
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

	auth struct {
		Level int    `json:"auth_level" cql:"auth_level"`
		Tags  []byte `json:"auth_tags" cql:"auth_tags"`
	}
)

func New(username, password, email, first_name, last_name string) (*UserModel, error) {
	id, err := upsilon.New(32, 0, 32)
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
		auth: auth{
			Level: 0,
			Tags:  []byte{},
		},
	}, nil
}
