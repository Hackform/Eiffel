package eta

import (
	"bytes"
	"crypto/rand"
	"golang.org/x/crypto/scrypt"
)

//////////
// Hash //
//////////

type (
	config struct {
		version,
		hash_length,
		salt_length,
		work_factor,
		mem_blocksize,
		parallel_factor int
	}
)

const (
	Latest = 1
)

var (
	v001 = &config{
		version:         1,
		hash_length:     64,
		salt_length:     32,
		work_factor:     16384,
		mem_blocksize:   8,
		parallel_factor: 1,
	}

	latest_config = v001
)

func newConfig(version int) *config {
	switch version {
	case v001.version:
		return v001
	default:
		return latest_config
	}
}

func (c *config) Version() int {
	return c.version
}

func Hash(password string, version int) (h, s []byte, v int, e error) {
	c := newConfig(version)
	salt := make([]byte, c.salt_length)
	_, err := rand.Read(salt)
	if err != nil {
		return []byte{}, salt, 0, err
	}
	hash, err := scrypt.Key([]byte(password), salt, c.work_factor, c.mem_blocksize, c.parallel_factor, c.hash_length)
	return hash, salt, c.version, err
}

func Verify(password string, salt, passhash []byte, version int) bool {
	c := newConfig(version)
	dk, err := scrypt.Key([]byte(password), salt, c.work_factor, c.mem_blocksize, c.parallel_factor, c.hash_length)
	if err != nil {
		return false
	}
	return bytes.Equal(dk, passhash)
}
