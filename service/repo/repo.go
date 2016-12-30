package repo

type (
	Data struct {
		Value interface{}
	}

	Tx interface {
		Query() []*Data
		QueryOne() *Data
		Exec() error
		Commit() error
		Rollback() error
	}

	Repo interface {
		Start() bool
		Shutdown()
		Transaction() Tx
	}
)
