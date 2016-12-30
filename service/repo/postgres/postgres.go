package postgres

import (
	"database/sql"
	"github.com/Hackform/Eiffel/service/repo"
	_ "github.com/lib/pq"
	"strings"
)

const (
	driverName = "postgres"
)

var (
	connectionArgs = []string{
		"dbname",
		"user",
		"password",
		"host",
		"port",
		"connect_timeout",
		"fallback_application_name",
		"sslmode",
		"sslcert",
		"sslkey",
		"sslrootcert",
	}
)

type (
	ConnectionString map[string]string

	Postgres struct {
		repo.RepoBase
		dcn ConnectionString
		db  *sql.DB
	}
)

func (s *ConnectionString) stringify() string {
	k := ""
	for _, i := range connectionArgs {
		if val, ok := (*s)[i]; ok {
			k += i + "=" + val + " "
		}
	}
	return strings.TrimSpace(k)
}

func New(s ConnectionString) *Postgres {
	return &Postgres{
		dcn: s,
	}
}

func (p *Postgres) Connect() bool {
	db, err := sql.Open(driverName, p.dcn.stringify())
	if err != nil {
		return false
	}
	p.db = db
	return true
}

func (p *Postgres) Disconnect() {
	p.db.Close()
}
