package upsilon

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"github.com/Hackform/Eiffel/service/util/tau"
)

///////////////////////
// Unique Identifier //
///////////////////////

type (
	// Upsilon is an identifier that can be initialized with a custom length composed of a user specified time, hash, and random bits
	Upsilon struct {
		timebits,
		hashbits,
		randbits,
		size int
		u []byte
	}
)

// New creates a new Upsilon
// arg input is only required if arg hashSize is greater than zero
func New(timeSize, hashSize, randomSize int, input ...[]byte) (returnUpsilon *Upsilon, returnError error) {
	defer func() {
		if r := recover(); r != nil {
			if r, ok := r.(error); ok {
				returnError = r
			} else {
				returnError = errors.New("Create Upsilon panic")
			}
		}
	}()

	k := new(bytes.Buffer)

	t := make([]byte, timeSize)
	binary.BigEndian.PutUint64(t, tau.Timestamp())
	k.Write(t)

	var h []byte
	if hashSize > 0 {
		if len(input) < 1 {
			return nil, errors.New("No input provided")
		}
		h = make([]byte, hashSize)
		l := len(input[0]) - hashSize
		for i := 0; i < len(h); i++ {
			if l+i > -1 {
				h[i] = input[0][l+i]
			}
		}
		k.Write(h)
	}

	r := make([]byte, randomSize)
	_, err := rand.Read(r)
	if err != nil {
		return nil, err
	}
	k.Write(r)

	return &Upsilon{
		timebits: timeSize,
		hashbits: hashSize,
		randbits: randomSize,
		size:     timeSize + hashSize + randomSize,
		u:        k.Bytes(),
	}, nil
}

// FromBytes creates a new Upsilon from an existing byte slice
func FromBytes(timeSize, hashSize, randomSize int, b []byte) *Upsilon {
	return &Upsilon{
		timebits: timeSize,
		hashbits: hashSize,
		randbits: randomSize,
		size:     timeSize + hashSize + randomSize,
		u:        b,
	}
}

// Bytes returns the full raw bytes of an Upsilon
func (u *Upsilon) Bytes() []byte {
	return u.u
}

// Time returns only the time bytes of an Upsilon
func (u *Upsilon) Time() []byte {
	return u.u[:u.timebits]
}

// Hash returns only the hash initialization bytes of an Upsilon
func (u *Upsilon) Hash() []byte {
	return u.u[u.timebits : u.timebits+u.hashbits]
}

// Rand returns only the random bytes of an Upsilon
func (u *Upsilon) Rand() []byte {
	return u.u[u.timebits+u.hashbits:]
}

// Base64 returns the full raw bytes of an Upsilon encoded in standard padded base64
func (u *Upsilon) Base64() string {
	return base64.StdEncoding.EncodeToString(u.u)
}

// TimeBase64 returns only the time bytes of an Upsilon encoded in standard padded base64
func (u *Upsilon) TimeBase64() string {
	return base64.StdEncoding.EncodeToString(u.Time())
}

// HashBase64 returns only the hash initialization bytes of an Upsilon encoded in standard padded base64
func (u *Upsilon) HashBase64() string {
	return base64.StdEncoding.EncodeToString(u.Hash())
}

// RandBase64 returns only the random bytes of an Upsilon encoded in standard padded base64
func (u *Upsilon) RandBase64() string {
	return base64.StdEncoding.EncodeToString(u.Rand())
}
