package q

///////
// Q //
///////

type (
	Q struct {
		Action        int
		Sector        string
		RProps        Props
		Cons          Constraints
		Vals          Props
		Mods          Constraints
		Limit, Offset int
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
		Limit:  1,
		Offset: 0,
	}
}

func NewQMulti(sector string, props Props, cons Constraints, limit int, offset int) Q {
	return Q{
		Action: ACTION_QUERY_MULTI,
		Sector: sector,
		RProps: props,
		Cons:   cons,
		Limit:  limit,
		Offset: offset,
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

func NewDB(sector string, cons Constraints) Q {
	return Q{
		Action: ACTION_CREATE_TABLE,
		Sector: sector,
		Cons:   cons,
	}
}
