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
		Key        string
		Condition  int
		Value      string
		Con1, Con2 *Constraint
	}
)

func NewConstraint(key string, condition int, value string) *Constraint {
	return &Constraint{
		Key:       key,
		Condition: condition,
		Value:     value,
	}
}

func NewOperator(condition int, con1, con2 *Constraint) *Constraint {
	return &Constraint{
		Condition: condition,
		Con1:      con1,
		Con2:      con2,
	}
}
