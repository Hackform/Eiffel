package kappa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Kappa_Get(t *testing.T) {
	assert := assert.New(t)
	k := New()
	assert.Equal(k.Get(), 1, "property value should start at 1")
	assert.Equal(k.Get(), 2, "property value should increment by 1")
	assert.Equal(k.Get(), 3, "property value should increment by 1")
}
