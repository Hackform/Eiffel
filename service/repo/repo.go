package repo

type (
	Tx interface {
	}

	Repo interface {
		Connect() bool
		Disconnect()
		Transaction() Tx
	}

	RepoBase struct{}
)

func (r *RepoBase) Start() bool {
	return r.Connect()
}

func (r *RepoBase) Shutdown() {
	r.Disconnect()
}

// Override

func (r *RepoBase) Connect() bool {
	return false
}

func (r *RepoBase) Disconnect() {
}
