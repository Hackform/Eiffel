package user

import (
	"errors"
	"github.com/Hackform/Eiffel/service/repo/cassandra"
	"github.com/Hackform/Eiffel/service/util/upsilon"
	"github.com/gocql/gocql"
	"strings"
)

const (
	tableName    = "users"
	tableNameMap = "users_usernames_idx"
)

var (
	cassModelFields            = []string{"id", "email", "username", "auth_level", "auth_tags", "first_name", "last_name", "date", "pass_hash", "pass_salt", "pass_version"}
	cassModelFieldsString      = strings.Join(cassModelFields, ", ")
	cassModelPlaceholderString = strings.Trim(strings.Repeat("?, ", len(cassModelFields)), ", ")

	cassQueryByIDString   = "SELECT " + cassModelFieldsString + " FROM " + tableName + " WHERE id = ? LIMIT 1"
	cassInsertTableString = "INSERT INTO " + tableName + " (" + cassModelFieldsString + ") VALUES (" + cassModelPlaceholderString + ")"

	cassIdxFields            = []string{"id", "username"}
	cassIdxFieldsString      = strings.Join(cassIdxFields, ", ")
	cassIdxPlaceholderString = strings.Trim(strings.Repeat("?, ", len(cassIdxFields)), ", ")

	cassQueryByUsernameString = "SELECT " + cassIdxFieldsString + " FROM " + tableName + " WHERE username = ? LIMIT 1"
	cassInsertIdxString       = "INSERT INTO " + tableName + " (" + cassIdxFieldsString + ") VALUES (" + cassIdxPlaceholderString + ")"
)

func cassCreate(t *cassandra.Tx) error {
	if err := t.S.Query(cassandra.BuilderTable(tableName, cassandra.Fields{
		"id":           "uuid",
		"email":        "varchar",
		"username":     "varchar",
		"auth_level":   "smallint",
		"auth_tags":    "blob",
		"first_name":   "varchar",
		"last_name":    "varchar",
		"date":         "timestamp",
		"pass_hash":    "blob",
		"pass_salt":    "blob",
		"pass_version": "int",
	}, []string{"id"}, nil)).RetryPolicy(nil).Exec(); err != nil {
		return errors.New("Unable to create table users")
	}

	if err := t.S.Query(cassandra.BuilderTable(tableNameMap, cassandra.Fields{
		"id":       "uuid",
		"username": "varchar",
	}, []string{"username"}, nil)).RetryPolicy(nil).Exec(); err != nil {
		return errors.New("Unable to create table users_usernames_idx")
	}

	return nil
}

func cassSelect(t *cassandra.Tx, u *upsilon.Upsilon) (*Model, error) {
	gocqlid, err := gocql.UUIDFromBytes(u.Bytes())
	if err != nil {
		return nil, err
	}

	k := Model{}
	var id gocql.UUID
	if err = t.S.Query(cassQueryByIDString, gocqlid).Scan(&id, &k.Email, &k.Username, &k.auth.Level, &k.auth.Tags, &k.name.First, &k.name.Last, &k.Date, &k.passhash.Hash, &k.passhash.Salt, &k.passhash.Version); err != nil {
		return nil, err
	}
	k.ID, err = upsilon.FromBytes(modelIDTimeBits, modelIDHashBits, modelIDRandBits, id.Bytes())
	if err != nil {
		return nil, err
	}

	return &k, nil
}

func cassSelectByUsername(t *cassandra.Tx, username string) (*Model, error) {
	var id gocql.UUID
	var uname string

	if err := t.S.Query(cassQueryByUsernameString, username).Scan(&id, &uname); err != nil {
		return nil, err
	}

	k, err := upsilon.FromBytes(modelIDTimeBits, modelIDHashBits, modelIDRandBits, id.Bytes())
	if err != nil {
		return nil, err
	}

	return cassSelect(t, k)
}

func cassInsert(t *cassandra.Tx, u *Model) error {
	gocqlid, err := gocql.UUIDFromBytes(u.ID.Bytes())
	if err != nil {
		return err
	}

	if err := t.S.Query(cassInsertTableString, gocqlid, u.Email, u.Username, u.auth.Level, u.auth.Tags, u.name.First, u.name.Last, u.Date, u.passhash.Hash, u.passhash.Salt, u.passhash.Version).Exec(); err != nil {
		return err
	}

	if err := t.S.Query(cassInsertIdxString, gocqlid, u.Username).Exec(); err != nil {
		return err
	}

	return nil
}
