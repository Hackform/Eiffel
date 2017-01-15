package q

/////////////
// Actions //
/////////////

const (
	ACTION_QUERY_MULTI = iota
	ACTION_QUERY_ONE
	ACTION_INSERT
	ACTION_UPDATE
	ACTION_CREATE_TABLE
)

////////////////
// Conditions //
////////////////

const (
	EQUAL = iota
	UNEQUAL
	GREATER
	LESSER
	GREATER_EQ
	LESSER_EQ
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
		Vals   Props
		Mods   Constraints
		Limit  int
	}

	Props []string

	Constraints []*Constraint
)

func NewQOne(sector string, props Props, cons Constraints) Q {
	return Q{
		Action: ACTION_QUERY_ONE,
		Sector: sector,
		RProps: props,
		Cons:   cons,
	}
}

func NewQMulti(sector string, props Props, cons Constraints, limit int) Q {
	return Q{
		Action: ACTION_QUERY_MULTI,
		Sector: sector,
		RProps: props,
		Cons:   cons,
		Limit:  limit,
	}
}

func NewI(sector string, props Props, vals Props) Q {
	return Q{
		Action: ACTION_INSERT,
		Sector: sector,
		RProps: props,
		Vals:   vals,
	}
}

func NewU(sector string, mods Constraints, cons Constraints) Q {
	return Q{
		Action: ACTION_UPDATE,
		Sector: sector,
		Mods:   mods,
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

func NewEq(key string, value string) *Constraint {
	return &Constraint{
		Key:       key,
		Condition: EQUAL,
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
