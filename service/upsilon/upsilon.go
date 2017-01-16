package upsilon

type (
	Upsilon struct {
		timebits,
		hashbits,
		randbits int
		up []byte
	}
)

func New(time, hash, rand int) *Upsilon {
	k := []byte{}
	return &Upsilon{
		timebits: time,
		hashbits: hash,
		randbits: rand,
		up:       k,
	}
}
