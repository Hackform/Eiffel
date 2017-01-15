package q

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
