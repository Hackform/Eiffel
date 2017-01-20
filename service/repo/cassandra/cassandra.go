package cassandra

import (
	"errors"
	"github.com/Hackform/Eiffel/service/repo"
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

	cassOpts struct {
		model      interface{}
		kpartition []string
		kcluster   []string
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

func (c *Cassandra) Transaction(table string, opts *repo.Opts) (repo.Tx, error) {
	o, err := parseOpts(opts)
	if err != nil {
		return nil, err
	}
	return c.keySpace.Table(table, o.model, gocassa.Keys{
		PartitionKeys:     o.kpartition,
		ClusteringColumns: o.kcluster,
	}), nil
}

func parseOpts(opts *repo.Opts) (*cassOpts, error) {
	var model interface{}
	var kpartition, kcluster []string
	var ok bool

	if model, ok = (*opts)["model"]; !ok {
		return nil, errors.New("model is not provided")
	}

	if _, ok = (*opts)["partition"]; !ok {
		return nil, errors.New("partition keys not provided")
	}

	if kpartition, ok = ((*opts)["partition"]).([]string); !ok {
		return nil, errors.New("partition keys are not a []string")
	}

	if _, ok = (*opts)["cluster"]; !ok {
		return nil, errors.New("cluster keys are not a []string")
	}

	if kcluster, ok = ((*opts)["cluster"]).([]string); ok {
		return nil, errors.New("cluster keys not provided")
	}

	return &cassOpts{
		model:      model,
		kpartition: kpartition,
		kcluster:   kcluster,
	}, nil
}
