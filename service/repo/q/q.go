package q

/////////////
// Actions //
/////////////

const (
	ACTION_QUERY_ALL = iota
	ACTION_QUERY_ONE
)

////////////////
// Conditions //
////////////////

const (
	EQUAL = iota
	UNEQUAL
	AND
	OR
)

///////
// Q //
///////

type (
	Q struct {
		Action int
		Sector string
		RProps Props
		Cons   Constraints
	}

	Props []string

	Constraints []*Constraint
)

func New(action int, sector string, props Props, cons Constraints) Q {
	return Q{
		Action: action,
		Sector: sector,
		RProps: props,
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

func NewCon(key string, condition int, value string) *Constraint {
	return &Constraint{
		Key:       key,
		Condition: condition,
		Value:     value,
	}
}

func NewOp(con1 *Constraint, condition int, con2 *Constraint) *Constraint {
	return &Constraint{
		Condition: condition,
		Con1:      con1,
		Con2:      con2,
	}
}
