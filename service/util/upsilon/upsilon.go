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
	Upsilon struct {
		timebits,
		hashbits,
		randbits,
		size int
		t, h, r, u []byte
	}
)

func New(time_size, hash_size, random_size int, input ...[]byte) (return_upsilon *Upsilon, return_error error) {
	defer func() {
		if r := recover(); r != nil {
			if r, ok := r.(error); ok {
				return_error = r
			} else {
				return_error = errors.New("Create Upsilon panic")
			}
		}
	}()

	k := new(bytes.Buffer)

	t := make([]byte, time_size)
	binary.BigEndian.PutUint64(t, tau.Timestamp())
	k.Write(t)

	var h []byte
	if hash_size > 0 {
		if len(input) < 1 {
			return nil, errors.New("No input provided")
		}
		h = make([]byte, hash_size)
		l := len(input[0]) - hash_size
		for i := 0; i < len(h); i++ {
			if l+i > -1 {
				h[i] = input[0][l+i]
			}
		}
		k.Write(h)
	}

	r := make([]byte, random_size)
	_, err := rand.Read(r)
	if err != nil {
		return nil, err
	}
	k.Write(r)

	return &Upsilon{
		timebits: time_size,
		hashbits: hash_size,
		randbits: random_size,
		size:     time_size + hash_size + random_size,
		t:        t,
		h:        h,
		r:        r,
		u:        k.Bytes(),
	}, nil
}

func (u *Upsilon) Bytes() []byte {
	return u.u
}

func (u *Upsilon) Time() []byte {
	return u.t
}

func (u *Upsilon) Hash() []byte {
	return u.h
}

func (u *Upsilon) Rand() []byte {
	return u.r
}

func (u *Upsilon) Base64() string {
	return base64.StdEncoding.EncodeToString(u.u)
}

func (u *Upsilon) TimeBase64() string {
	return base64.StdEncoding.EncodeToString(u.t)
}

func (u *Upsilon) HashBase64() string {
	return base64.StdEncoding.EncodeToString(u.h)
}

func (u *Upsilon) RandBase64() string {
	return base64.StdEncoding.EncodeToString(u.r)
}
