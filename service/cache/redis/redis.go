package redis

import (
	"github.com/hackform/eiffel/service/cache"
	goredis "gopkg.in/redis.v5"
)

const (
	// AdapterID is the unique string identifying the type of transaction
	AdapterID = "redis"
)

type (
	red struct {
		client *goredis.ClusterClient
		props  connectionProps
	}

	connectionProps struct {
		nodeIps  []string
		password string
	}
)

// New creates a new redis client instance
func New(nodeIps []string, password string) cache.Cache {
	return &red{
		props: connectionProps{
			nodeIps:  nodeIps,
			password: password,
		},
	}
}

func (r *red) Start() error {
	client := goredis.NewClusterClient(&goredis.ClusterOptions{
		Addrs:    r.props.nodeIps,
		Password: r.props.password,
	})
	if err := client.Ping().Err(); err != nil {
		return err
	}
	r.client = client
	return nil
}

func (r *red) Shutdown() {
	r.client.Close()
}

func (r *red) Transaction() (cache.Tx, error) {
	return nil, nil
}
