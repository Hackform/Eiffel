package setup

type (
	SetupModel struct {
		Name    string `cql:"eiffel_name"`
		Setup   bool   `cql:"eiffel_setup_complete"`
		Version string `cql:"eiffel_version"`
	}
)

func NewModel() *SetupModel {
	return &SetupModel{}
}
