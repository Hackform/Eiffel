package kappa

//////////////////////////
// Incremented Constant //
//////////////////////////

type (
	Kappa struct {
		value int
	}
)

func New() *Kappa {
	return &Kappa{
		value: 0,
	}
}

func (k *Kappa) Get() int {
	k.value++
	return k.value
}
