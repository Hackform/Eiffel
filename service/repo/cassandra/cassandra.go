package cassandra

import (
	"github.com/gocassa/gocassa"
)

type (
	Cassandra struct {
		keySpace gocassa.KeySpace
		props    connectionProps
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

func (c *Cassandra) Start() bool {
	keyspace, err := gocassa.ConnectToKeySpace(c.props.keySpace, c.props.nodeIps, c.props.username, c.props.password)
	if err != nil {
		return false
	}
	c.keySpace = keyspace
	return true
}

func (c *Cassandra) Shutdown() {
}
