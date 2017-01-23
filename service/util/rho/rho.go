package rho

//////////
// Rank //
//////////

const (
	root       = 1
	superuser  = 8
	admin      = 16
	maintainer = 32
	mod        = 128
	auser      = 176
	user       = 192
	apublic    = 240
	public     = 255
)

func User() int {
	return user
}

func Admin() int {
	return admin
}

func IsPrivileged(k int) bool {
	return k < mod
}
