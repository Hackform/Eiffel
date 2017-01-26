package setup

import (
	"errors"
	"github.com/Hackform/Eiffel/service/repo/cassandra"
)

const (
	tableName    = "eiffel_setup"
	setupName    = "hackform.eiffel"
	setupVersion = "v0.1.0"
)

func cassCreate(t *cassandra.Tx) error {
	if err := t.S.Query(cassandra.BuilderTable(tableName, cassandra.Fields{
		"eiffel_name":           "varchar",
		"eiffel_setup_complete": "boolean",
		"eiffel_version":        "varchar",
	}, []string{"eiffel_name"}, nil)).RetryPolicy(nil).Exec(); err != nil {
		return errors.New("Unable to create table eiffel_setup")
	}
	return nil
}

func cassSelect(t *cassandra.Tx) (*Model, error) {
	k := Model{}
	if err := t.S.Query(`SELECT name, setup_complete, version FROM eiffel_setup WHERE eiffel_name = ? LIMIT 1`,
		setupName).Scan(&k.Name, &k.Setup, &k.Version); err != nil {
		return nil, errors.New("Unable to get setup table")
	}
	return &k, nil
}

func cassInsert(t *cassandra.Tx) error {
	if err := t.S.Query(`INSERT INTO eiffel_setup (name, setup_complete, version) VALUES (?, ?, ?)`,
		setupName, true, setupVersion).Exec(); err != nil {
		return errors.New("Unable to insert setup complete")
	}
	return nil
}
