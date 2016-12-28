package postgres

import (
	"database/sql"
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
		"fallback_application_name",
		"connect_timeout",
		"sslmode",
		"sslcert",
		"sslkey",
		"sslrootcert",
	}
)

type (
	connectionString map[string]string

	Postgres struct {
		dcn connectionString
		db  *sql.DB
	}
)

func (s *connectionString) stringify() string {
	k := ""
	for _, i := range connectionArgs {
		if val, ok := (*s)[i]; ok {
			k += i + "=" + val + " "
		}
	}
	return strings.TrimSpace(k)
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
