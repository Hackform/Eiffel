package cassandra

import (
	"errors"
	"github.com/Hackform/Eiffel/service/repo"
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

func (c *cassandra) Start() bool {
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

func (c *cassandra) Shutdown() {
}

func (c *cassandra) Transaction() (repo.Tx, error) {
	return newTx(c)
}

func (c *cassandra) Setup() error {
	setupObj := &setupModel{}
	if err := c.space[setup_table_name].Where(gocassa.Eq(setup_table_pk, setup_name)).ReadOne(setupObj); err != nil {
		return errors.New("database already configured")
	} else {
		for _, v := range c.space {
			if err := v.CreateIfNotExist(); err != nil {
				return err
			}
		}
		return nil
	}
}

/////////////////
// Transaction //
/////////////////

type (
	transaction struct {
		c       *cassandra
		actions gocassa.Op
	}
)

func newTx(c *cassandra) (*transaction, error) {
	return &transaction{
		c:       c,
		actions: nil,
	}, nil
}

func (t *transaction) Commit() error {
	return nil
}

func (t *transaction) Rollback() error {
	return nil
}

func (t *transaction) Insert(sector string, d *repo.Data) error {
	if t.actions == nil {
		t.actions = t.c.space[sector].Set(d.Value)
	} else {
		t.actions = t.actions.Add(t.c.space[sector].Set(d.Value))
	}
	return nil
}
