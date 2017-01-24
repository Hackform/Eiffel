package rho

//////////
// Rank //
//////////

const (
	root       uint8 = 1
	superuser  uint8 = 8
	admin      uint8 = 16
	maintainer uint8 = 32
	mod        uint8 = 128
	auser      uint8 = 176
	user       uint8 = 192
	apublic    uint8 = 240
	public     uint8 = 255
)

func User() uint8 {
	return user
}

func Admin() uint8 {
	return admin
}

func SuperUser() uint8 {
	return superuser
}

func IsPrivileged(k uint8) bool {
	return k < mod
}
