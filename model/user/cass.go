package user

import (
	"errors"
	"github.com/Hackform/Eiffel/service/repo/cassandra"
	"github.com/Hackform/Eiffel/service/util/upsilon"
)

const (
	tableName = "users"
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
	return nil
}

func cassSelect(t *cassandra.Tx, u *upsilon.Upsilon) (*Model, error) {
	k := Model{}
	// if err := t.S.Query(`SELECT name, setup_complete, version FROM eiffel_setup WHERE eiffel_name = ? LIMIT 1`,
	// 	setup_name).Scan(&k.Name, &k.Setup, &k.Version); err != nil {
	// 	return nil, errors.New("Unable to get setup table")
	// }
	return &k, nil
}

func cassInsert(t *cassandra.Tx, u *Model) error {
	// if err := t.S.Query(`INSERT INTO eiffel_setup (name, setup_complete, version) VALUES (?, ?, ?)`,
	// 	setup_name, true, setup_version).Exec(); err != nil {
	// 	return errors.New("Unable to insert setup complete")
	// }
	return nil
}
