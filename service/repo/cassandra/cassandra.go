package cassandra

import (
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/gocql/gocql"
)

const (
	AdapterId = "cassandra"
)

///////////////
// Cassandra //
///////////////

type (
	Cassandra struct {
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

func New(keyspace string, nodeIps []string, username, password string) *Cassandra {
	return &Cassandra{
		props: connectionProps{
			keySpace: keyspace,
			nodeIps:  nodeIps,
			username: username,
			password: password,
		},
	}
}

func (c *Cassandra) Start() error {
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

func (c *Cassandra) Shutdown() {
	c.session.Close()
}

func (c *Cassandra) Transaction() (repo.Tx, error) {
	return newTx(c)
}

func (c *Cassandra) Setup() error {

	// create tables

	return nil
}

/////////////////
// Transaction //
/////////////////

type (
	Tx struct {
		C *Cassandra
	}
)

func newTx(c *Cassandra) (*Tx, error) {
	return &Tx{
		C: c,
	}, nil
}

func (t *Tx) Adapter() string {
	return AdapterId
}
