package user

import (
	// "github.com/Hackform/Eiffel/service/util/upsilon"
	"github.com/gocql/gocql"
	"time"
)

type (

	///////////////
	// UserModel //
	///////////////

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
		email
		userInfoPub
	}

	email struct {
		Email string `json:"email" cql:"email"`
	}

	userInfoPub struct {
		username
		name
		props
	}

	username struct {
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
		Hash []byte `cql:"password"`
		Salt []byte `cql:"salt"`
	}

	auth struct {
		Level int    `json:"auth_level" cql:"auth_level"`
		Tags  []byte `json:"auth_tags" cql:"auth_tags"`
	}

	//////////////
	// Requests //
	//////////////

	password struct {
		Password string `json:"password"`
	}

	reqNewUser struct {
		userInfo
		password
	}
)
