package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/hackform/eiffel/service/repo"
)

const (
	// AdapterID is the unique string identifying the type of each transaction
	AdapterID = "cassandra"
)

///////////////
// cass //
///////////////

type (
	cass struct {
		session *gocql.Session
		props   connectionProps
	}

	connectionProps struct {
		keySpace string
		nodeIps  []string
		username,
		password string
	}
)

// New creates a new cassandra client instance
func New(keyspace string, nodeIps []string, username, password string) repo.Repo {
	return &cass{
		props: connectionProps{
			keySpace: keyspace,
			nodeIps:  nodeIps,
			username: username,
			password: password,
		},
	}
}

func (c *cass) Start() error {
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

	return nil
}

func (c *cass) Shutdown() {
	c.session.Close()
}

func (c *cass) Transaction() (repo.Tx, error) {
	return newTx(c)
}

/////////////////
// Transaction //
/////////////////

type (
	// Tx is a transaction containing a cassandra session
	Tx struct {
		S *gocql.Session
	}
)

func newTx(c *cass) (*Tx, error) {
	return &Tx{
		S: c.session,
	}, nil
}

// Adapter returns the unique string identifier of the transaction type
func (t *Tx) Adapter() string {
	return AdapterID
}
