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

// User returns the user rank
func User() uint8 {
	return user
}

// Admin returns the admin rank
func Admin() uint8 {
	return admin
}

// SuperUser returns the superuser rank
func SuperUser() uint8 {
	return superuser
}

// IsPrivileged returns whether a rank is a privileged rank
func IsPrivileged(k uint8) bool {
	return k < mod
}
