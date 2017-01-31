package user

import (
	"errors"
	"github.com/Hackform/Eiffel/service/repo/cassandra"
	"github.com/Hackform/Eiffel/service/util/upsilon"
	"github.com/gocql/gocql"
)

const (
	tableName    = "users"
	tableNameMap = "users_usernames_idx"
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
	if err := t.S.Query(`SELECT id, email, username, auth_level, auth_tags, first_name, last_name, date, pass_hash, pass_salt, pass_version FROM users WHERE id = ? LIMIT 1`,
		gocqlid).Scan(&id, &k.Email, &k.Username, &k.auth.Level, &k.auth.Tags, &k.name.First, &k.name.Last, &k.Date, &k.passhash.Hash, &k.passhash.Salt, &k.passhash.Version); err != nil {
		return nil, errors.New("Unable to get setup table")
	}

	// convert gocql uuid to upsilon

	return &k, nil
}

func cassInsert(t *cassandra.Tx, u *Model) error {
	// if err := t.S.Query(`INSERT INTO eiffel_setup (name, setup_complete, version) VALUES (?, ?, ?)`,
	// 	setup_name, true, setup_version).Exec(); err != nil {
	// 	return errors.New("Unable to insert setup complete")
	// }
	return nil
}
