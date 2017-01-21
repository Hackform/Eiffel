package cassandra

import (
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/gocassa/gocassa"
)

type (
	Cassandra struct {
		keySpace gocassa.KeySpace
		space    map[string]gocassa.Table
		props    connectionProps
		config   Config
	}

	connectionProps struct {
		keySpace string
		nodeIps  []string
		username,
		password string
	}

	cassOpts struct {
		model      interface{}
		kpartition []string
		kcluster   []string
	}

	Config map[string]*cassOpts
)

func Opts(model interface{}, kpartition, kcluster []string) *cassOpts {
	return &cassOpts{
		model:      model,
		kpartition: kpartition,
		kcluster:   kcluster,
	}
}

func New(keyspace string, nodeIps []string, username, password string, config Config) *Cassandra {
	return &Cassandra{
		props: connectionProps{
			keySpace: keyspace,
			nodeIps:  nodeIps,
			username: username,
			password: password,
		},
		config: config,
	}
}

func (c *Cassandra) Start() bool {
	keyspace, err := gocassa.ConnectToKeySpace(c.props.keySpace, c.props.nodeIps, c.props.username, c.props.password)
	if err != nil {
		return false
	}
	c.keySpace = keyspace

	for k, v := range c.config {
		c.space[k] = c.keySpace.Table(k, v.model, gocassa.Keys{
			PartitionKeys:     v.kpartition,
			ClusteringColumns: v.kcluster,
		})
	}
	return true
}

func (c *Cassandra) Shutdown() {
}

func (c *Cassandra) Transaction() (repo.Tx, error) {
	return newTx()
}

type (
	transaction struct {
	}
)

func newTx() (*transaction, error) {
	return &transaction{}, nil
}

func (t *transaction) Commit() error {
	return nil
}

func (t *transaction) Rollback() error {
	return nil
}
