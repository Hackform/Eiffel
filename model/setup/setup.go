package setup

import (
	"errors"
	"github.com/hackform/eiffel/service/repo"
	"github.com/hackform/eiffel/service/repo/cassandra"
)

type (
	// Model defines the setup configuration
	Model struct {
		Name    string `cql:"name"`
		Setup   bool   `cql:"setup_complete"`
		Version string `cql:"version"`
	}
)

// New creates a new Model
func New() *Model {
	return &Model{}
}

// Create creates a new Setup Table on the cassandra cluster
func Create(t repo.Tx) error {
	switch t.Adapter() {
	case cassandra.AdapterID:
		return cassCreate(t.(*cassandra.Tx))
	default:
		return errors.New("Repo adapter not found")
	}
}

// Select finds the setup configuration on the cassandra cluster
func Select(t repo.Tx) (*Model, error) {
	switch t.Adapter() {
	case cassandra.AdapterID:
		return cassSelect(t.(*cassandra.Tx))
	default:
		return nil, errors.New("Repo adapter not found")
	}
}

// Insert creates a new setup configuration on the cassandra cluster
func Insert(t repo.Tx) error {
	switch t.Adapter() {
	case cassandra.AdapterID:
		return cassInsert(t.(*cassandra.Tx))
	default:
		return errors.New("Repo adapter not found")
	}
}
