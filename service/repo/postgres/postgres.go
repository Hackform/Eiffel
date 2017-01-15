package postgres

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/Hackform/Eiffel/service/repo"
	_ "github.com/lib/pq"
	"strings"
)

const (
	driverName      = "postgres"
	escape_sequence = "$"
)

///////////////////////
// Connection String //
///////////////////////

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
)

func (s *ConnectionString) stringify() string {
	k := bytes.Buffer{}
	for _, i := range connectionArgs {
		if val, ok := (*s)[i]; ok {
			k.WriteString(i + "=" + val + " ")
		}
	}
	return strings.TrimSpace(k.String())
}

//////////////
// Postgres //
//////////////

type (
	Postgres struct {
		dcn ConnectionString
		db  *sql.DB
	}
)

func New(s ConnectionString) *Postgres {
	return &Postgres{
		dcn: s,
	}
}

func (p *Postgres) Start() bool {
	db, err := sql.Open(driverName, p.dcn.stringify())
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	p.db = db
	return true
}

func (p *Postgres) Shutdown() {
	p.db.Close()
}

func (p *Postgres) Transaction() (repo.Tx, error) {
	t, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	return newTx(t), nil
}
