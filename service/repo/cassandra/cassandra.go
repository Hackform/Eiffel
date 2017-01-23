package cassandra

import (
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/gocql/gocql"
)

const (
	setup_table_name = "eiffel_setup"
	setup_table_pk   = "eiffel_name"

	setup_name    = "hackform.eiffel"
	setup_version = "v0.1.0"
)

//////////
// Opts //
//////////

type (
	CassOpts struct {
		model      interface{}
		kpartition []string
		kcluster   []string
	}

	Config map[string]*CassOpts
)

func Opts(model interface{}, kpartition, kcluster []string) *CassOpts {
	return &CassOpts{
		model:      model,
		kpartition: kpartition,
		kcluster:   kcluster,
	}
}

///////////
// Setup //
///////////

type (
	setupModel struct {
		name    string `cql:"eiffel_name"`
		setup   bool   `cql:"eiffel_setup_complete"`
		version string `cql:"eiffel_version"`
	}
)

func sampleSetupModel() *setupModel {
	return &setupModel{}
}

///////////////
// Cassandra //
///////////////

type (
	cassandra struct {
		session *gocql.Session
		props   connectionProps
		config  Config
	}

	connectionProps struct {
		keySpace string
		nodeIps  []string
		username,
		password string
	}
)

func New(keyspace string, nodeIps []string, username, password string, config Config) *cassandra {
	return &cassandra{
		props: connectionProps{
			keySpace: keyspace,
			nodeIps:  nodeIps,
			username: username,
			password: password,
		},
		config: config,
	}
}

func (c *cassandra) Start() error {
	cluster := gocql.NewCluster(c.props.nodeIps...)
	cluster.Keyspace = c.props.keySpace
	cluster.Consistency = gocql.Quorum
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: c.props.username,
		Password: c.props.password,
	}

	s, err := cluster.CreateSession()
	if err != nil {
		return err
	}
	c.session = s

	c.config[setup_table_name] = Opts(sampleSetupModel(), []string{setup_table_pk}, nil)

	return nil
}

func (c *cassandra) Shutdown() {
	c.session.Close()
}

func (c *cassandra) Transaction() (repo.Tx, error) {
	return newTx(c)
}

func (c *cassandra) Setup() error {

	// create tables

	return nil
}

/////////////////
// Transaction //
/////////////////

type (
	transaction struct {
		c *cassandra
	}
)

func newTx(c *cassandra) (*transaction, error) {
	return &transaction{
		c: c,
	}, nil
}

func (t *transaction) Commit() error {
	return nil
}

func (t *transaction) Rollback() error {
	return nil
}

func (t *transaction) Insert(sector string, d *repo.Data) error {
	return nil
}
