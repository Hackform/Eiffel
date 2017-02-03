package repo

type (
	// Repo is a repository service interface
	Repo interface {
		Start() error
		Shutdown()
		Transaction() (Tx, error)
	}

	// Tx is a transaction interface
	Tx interface {
		Adapter() string
	}
)
