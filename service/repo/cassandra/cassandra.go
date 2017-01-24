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
	cassandra struct {
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

func New(keyspace string, nodeIps []string, username, password string) *cassandra {
	return &cassandra{
		props: connectionProps{
			keySpace: keyspace,
			nodeIps:  nodeIps,
			username: username,
			password: password,
		},
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

	return nil
}

func (c *cassandra) Shutdown() {
	c.session.Close()
}

func (c *cassandra) Transaction() (repo.Tx, error) {
	return newTx(c)
}

/////////////////
// Transaction //
/////////////////

type (
	Tx struct {
		S *gocql.Session
	}
)

func newTx(c *cassandra) (*Tx, error) {
	return &Tx{
		S: c.session,
	}, nil
}

func (t *Tx) Adapter() string {
	return AdapterId
}
