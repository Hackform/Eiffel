package tau

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Tau(t *testing.T) {
	assert := assert.New(t)

	assert.True(Timestamp() > 0)
}
