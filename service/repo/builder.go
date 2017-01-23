package repo

const (
	action_create_table = iota
	action_insert
	action_update
	action_query
)

type (
	Table struct {
		Fields []string
		Types  []string
	}

	Action struct {
		A      int
		Fields []string
		Values []string
	}
)
