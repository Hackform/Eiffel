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
		ColCon     int
		Size       int
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

func NewOrd(key string, order int) *Constraint {
	return &Constraint{
		Key:       key,
		Condition: order,
	}
}

func NewType(key string, dataType int, colcon int, size int) *Constraint {
	return &Constraint{
		Key:       key,
		Condition: dataType,
		ColCon:    colcon,
		Size:      size,
	}
}
