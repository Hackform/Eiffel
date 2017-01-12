package bound

///////////
// Bound //
///////////

type (
	Bound struct {
		Action int
		Sector string
		Vals   Values
		Cons   Constraints
	}

	Values []string

	Constraints []Constraint
)

func New(action int, sector, key string, vals Values, cons Constraints) Bound {
	return Bound{
		Action: action,
		Sector: sector,
		Vals:   vals,
		Cons:   cons,
	}
}

////////////////
// Constraint //
////////////////

type (
	Constraint struct {
		Key       string
		Condition int
		Value     interface{}
	}
)

func NewConstraint(key string, condition int, value interface{}) Constraint {
	return Constraint{
		Key:       key,
		Condition: condition,
		Value:     value,
	}
}
