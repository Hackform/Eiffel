package kappa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Kappa_Get(t *testing.T) {
	assert := assert.New(t)
	k := New()
	assert.Equal(k.Get(), 1, "kappa should start at 1")
	assert.Equal(k.Get(), 2, "kappa should increment by 1")
	assert.Equal(k.Get(), 3, "kappa should increment by 1")
}
