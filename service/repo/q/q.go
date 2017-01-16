package q

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
		Order  Constraints
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

func NewQMulti(sector string, props Props, cons Constraints, limit int, order Constraints) Q {
	return Q{
		Action: ACTION_QUERY_MULTI,
		Sector: sector,
		RProps: props,
		Cons:   cons,
		Limit:  limit,
		Order:  order,
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

func NewD(sector string, cons Constraints) Q {
	return Q{
		Action: ACTION_DELETE,
		Sector: sector,
		Cons:   cons,
	}
}

func NewT(sector string, dataTypes Constraints) Q {
	return Q{
		Action: ACTION_CREATE_TABLE,
		Sector: sector,
		Cons:   dataTypes,
	}
}
