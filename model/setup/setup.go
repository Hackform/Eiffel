package setup

import (
	"errors"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/Hackform/Eiffel/service/repo/cassandra"
)

type (
	SetupModel struct {
		Name    string `cql:"name"`
		Setup   bool   `cql:"setup_complete"`
		Version string `cql:"version"`
	}
)

func NewModel() *SetupModel {
	return &SetupModel{}
}

func Create(t repo.Tx) error {
	switch t.Adapter() {
	case cassandra.AdapterId:
		return cassCreate(t.(*cassandra.Tx))
	default:
		return errors.New("Repo adapter not found")
	}
}

func Select(t repo.Tx) (*SetupModel, error) {
	switch t.Adapter() {
	case cassandra.AdapterId:
		return cassSelect(t.(*cassandra.Tx))
	default:
		return nil, errors.New("Repo adapter not found")
	}
}

func Insert(t repo.Tx) error {
	switch t.Adapter() {
	case cassandra.AdapterId:
		return cassInsert(t.(*cassandra.Tx))
	default:
		return errors.New("Repo adapter not found")
	}
}
