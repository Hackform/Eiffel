package cassandra

import (
	"errors"
	"github.com/Hackform/Eiffel/service/repo"
	"github.com/gocassa/gocassa"
)

const (
	setup_table_name = "eiffel_setup"
	setup_table_pk   = "eiffel_name"

	setup_name    = "hackform.eiffel"
	setup_version = "v0.1.0"

	user_table_name = "users"
)

//////////
// Opts //
//////////

type (
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
)

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

	c.config[setup_table_name] = Opts(sampleSetupModel(), []string{setup_table_pk}, nil)

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
	return newTx(c)
}

func (c *Cassandra) Setup() error {
	setupObj := &setupModel{}
	if err := c.space[setup_table_name].Where(gocassa.Eq(setup_table_pk, setup_name)).ReadOne(setupObj); err != nil {
		return errors.New("database already configured")
	} else {
		for _, v := range c.space {
			if err := v.Create(); err != nil {
				return err
			}
		}
		c.space[user_table_name].Set(nil)
		return nil
	}
}

/////////////////
// Transaction //
/////////////////

type (
	transaction struct {
		c       *Cassandra
		actions []gocassa.Op
	}
)

func newTx(c *Cassandra) (*transaction, error) {
	return &transaction{
		c:       c,
		actions: []gocassa.Op{},
	}, nil
}

func (t *transaction) Commit() error {
	return nil
}

func (t *transaction) Rollback() error {
	return nil
}
